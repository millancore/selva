package selva

import (
	"context"
	"time"

	"github.com/joshuarubin/go-sway"
)

func AutoResize(ctx context.Context, client Client, ws *Workspace, focused *sway.Node) {
	time.Sleep(30 * time.Millisecond)

	visibleNodes := ws.Nodes

	config := client.Config

	// If only exist a visible Container
	if len(visibleNodes) == 1 {
		return
	}

	if focused == nil {
		return
	}

	// If focused Container is Floating
	if focused.Type == "floating_con" {
		return
	}

	// Calculate the width to be shared by the other containers
	leftWidth := (config.OuputWidth - config.FocusWidth) / (len(visibleNodes) - 1)

	for _, node := range visibleNodes {

		node := Node{node}

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

			// smooth reduce
			for reduce > 0 {
				client.Sway.RunCommand(ctx, node.Resize(Shrink(Right, 1)))
				reduce--
			}

			//client.Sway.RunCommand(ctx, node.Resize(Shrink(Right, reduce)))
		}

		// If current node is to the RIGTH of focused
		if node.Rect.X > focused.Rect.X {

			for reduce > 0 {
				client.Sway.RunCommand(ctx, node.Resize(Shrink(Left, 1)))
				reduce--
			}

			//client.Sway.RunCommand(ctx, node.Resize(Shrink(Left, reduce)))
		}
	}
}
