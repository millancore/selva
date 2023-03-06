package main

import (
	"context"
	"fmt"
	"github.com/joshuarubin/go-sway"
	"log"
)

const (
	width      = 2536
	focusWidth = 1500
)

type Node struct {
	*sway.Node
}

func (n Node) Resize(client sway.Client, resizer Resizer) {

	resizeCommand := fmt.Sprintf("[con_id=%d] %s", n.ID, resizer.Resize())
	client.RunCommand(context.Background(), resizeCommand)
}

func (n Node) VisibleNodes() []*Node {
	var visibleNodes []*Node

	for _, child := range n.Nodes {

		if child.Visible != nil && *child.Visible {
			visibleNodes = append(visibleNodes, &Node{child})
		}

		if len(child.Nodes) > 0 {
			visibleNodes = append(visibleNodes, Node{child}.VisibleNodes()...)
		}
	}

	return visibleNodes
}

type eventHandler struct {
	sway.EventHandler
	client sway.Client
}

func (eh eventHandler) Window(ctx context.Context, ev sway.WindowEvent) {
	client := eh.client
	tree, err := client.GetTree(ctx)

	if err != nil {
		log.Fatal(err)

	}

	rootNode := &Node{tree}
	visibleNodes := rootNode.VisibleNodes()

	focused := tree.FocusedNode()
    
	// If only exist a visible Container
	if len(visibleNodes) == 1 {
		return
	}

	// Calculate the width to be shared by the other containers
	leftWidth := (width - focusWidth) / (len(visibleNodes) - 1)

	for _, node := range visibleNodes {

		// Dont change size for focused
		if node.Focused {
			continue
		}

		// If the container is smaller than expected size
		if int(node.Rect.Width) <= leftWidth {
			continue
		}

		reduce := int(node.Rect.Width) - leftWidth

		// If current node is to the LEFT of focused
		if node.Rect.X < focused.Rect.X {
			node.Resize(client, Shrink(Right, reduce))
		}

		// If current node is to the RIGTH of focused
		if node.Rect.X > focused.Rect.X {
			node.Resize(client, Shrink(Left, reduce))
		}
	}
}

func main() {

	ctx := context.Background()

	client, err := sway.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	handler := eventHandler{
		EventHandler: sway.NoOpEventHandler(),
		client:       client,
	}

	err = sway.Subscribe(ctx, handler, sway.EventTypeWindow)
	if err != nil {
		log.Fatal(err)
	}
}
