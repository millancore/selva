package selva

import (
	"context"
	"fmt"

	"github.com/joshuarubin/go-sway"
)

type Workspace struct {
	Id      int64
	Num     int
	Name    string
	Focused bool
	Nodes   []*sway.Node
}

type WorkspaceList []*Workspace

func (wl *WorkspaceList) getByNum(num int) *Workspace {
	for _, ws := range *wl {
		if ws.Num == num {
			return ws
		}
	}

	return nil
}

func getFocused(sws []sway.Workspace) *sway.Workspace {
	for _, ws := range sws {
		if ws.Focused {
			return &ws
		}
	}

	return nil
}

func GetFocusedWorkspace(ctx context.Context, client Client) *Workspace {
	swayWorpaces, err := client.Sway.GetWorkspaces(ctx)

	if err != nil {
		fmt.Println(err)
	}

	tree, err := client.Sway.GetTree(ctx)
	if err != nil {
		fmt.Println(err)
	}

	swayFocused := getFocused(swayWorpaces)
	workspaces := GetWorkspaces(tree)

	focused := workspaces.getByNum(int(swayFocused.Num))
	focused.Focused = true

	return focused
}

func GetWorkspaces(root *sway.Node) WorkspaceList {

	workspaces := make(WorkspaceList, 0)

	for _, node := range root.Nodes {
		if node.Type == "workspace" {
			workspaces = append(workspaces, &Workspace{
				Id: node.ID,
				// cast string firts letter to int
				Num:     int(node.Name[0]) - 48,
				Name:    node.Name,
				Focused: node.Focused,
				Nodes:   node.Nodes,
			})
		}

		if len(node.Nodes) > 0 {
			workspaces = append(workspaces, GetWorkspaces(node)...)
		}
	}

	return workspaces
}
