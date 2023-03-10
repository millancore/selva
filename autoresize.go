package selva

import (
	"context"
	"time"
)

func AutoResize(ctx context.Context, client Client, rootNode *Node) {
	time.Sleep(100 * time.Millisecond)

	visibleNodes := rootNode.VisibleNodes()
	config := client.Config

	// If only exist a visible Container
	if len(visibleNodes) == 1 {
		return
	}

	focused := rootNode.FocusedNode()
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
			client.Sway.RunCommand(ctx, node.Resize(Shrink(Right, reduce)))
		}

		// If current node is to the RIGTH of focused
		if node.Rect.X > focused.Rect.X {
			client.Sway.RunCommand(ctx, node.Resize(Shrink(Left, reduce)))
		}
	}
}
