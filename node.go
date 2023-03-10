package selva

import (
	"fmt"
	"github.com/joshuarubin/go-sway"
)

type Node struct {
	*sway.Node
}

func (n Node) Resize(resizer Resizer) string {
	return fmt.Sprintf("[con_id=%d] %s", n.ID, resizer.Resize())
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
