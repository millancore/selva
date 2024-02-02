package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joshuarubin/go-sway"
	"github.com/millancore/selva"
)

type eventHandler struct {
	sway.EventHandler
	Client selva.Client
}

func (handler eventHandler) Window(ctx context.Context, ev sway.WindowEvent) {
	client := handler.Client

	rootNode, err := client.Sway.GetTree(ctx)

	ws := selva.GetFocusedWorkspace(ctx, client)

	if client.Config.SkipWorkspaces != nil {
		for _, wsNum := range client.Config.SkipWorkspaces {
			if ws.Num == wsNum {
				return
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	if ev.Change == "title" {
		return
	}

	if ev.Change == "urgent" {
		client.Sway.RunCommand(ctx, fmt.Sprintf("[con_id=%d] focus", ev.Container.ID))
	}

	// fmt.Println(ev.Change)
	selva.AutoResize(ctx, client, ws, rootNode.FocusedNode())
}

func main() {

	ctx := context.Background()

	swayClient, err := sway.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := selva.Client{
		Sway: swayClient,
		Config: selva.Config{
			OuputWidth:     2536,
			FocusWidth:     1500,
			SkipWorkspaces: []int{3, 6},
		},
	}

	handler := eventHandler{
		EventHandler: sway.NoOpEventHandler(),
		Client:       client,
	}

	err = sway.Subscribe(ctx, handler, sway.EventTypeWindow)
	if err != nil {
		log.Fatal(err)
	}
}
