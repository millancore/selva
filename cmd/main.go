package main

import (
	"context"
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
