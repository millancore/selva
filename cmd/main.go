package main

import (
	"context"
    "fmt"
    "github.com/joshuarubin/go-sway"
    "github.com/millancore/selva"
	"log"
)

type eventHandler struct {
	sway.EventHandler
	Client selva.Client
}

func (handler eventHandler) Window(ctx context.Context, ev sway.WindowEvent) {
    client := handler.Client

    rootNode, err := client.Sway.GetTree(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if ev.Change == "title" {
        return
    }

    if ev.Change == "urgent" {
        client.Sway.RunCommand(ctx, fmt.Sprintf("[con_id=%d] focus", ev.Container.ID))
    }


//    fmt.Println(ev.Change)
    selva.AutoResize(ctx, client, &selva.Node{Node: rootNode})
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
           OuputWidth: 2536,
            FocusWidth: 1500,
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
