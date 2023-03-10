package selva

import "fmt"

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type Resizer interface {
	Resize() string
}

type ResizeDirectionMode struct {
	Mode      string
	Direction Direction
	Pixels    int
}

func (rd ResizeDirectionMode) Resize() string {
	return fmt.Sprintf("resize %s %s %dpx", rd.Mode, rd.Direction, rd.Pixels)
}

type ResizeDimensionMode struct {
	Width  int
	Height int
}

func (rm ResizeDimensionMode) Resize() string {

	if rm.Width != 0 && rm.Height == 0 {
		return fmt.Sprintf("resize set width %dpx", rm.Width)
	}

	if rm.Height != 0 && rm.Width == 0 {
		return fmt.Sprintf("resize set height %dpx", rm.Height)
	}

	return fmt.Sprintf("resize set width %dpx height %dpx", rm.Width, rm.Height)
}

func Grow(direction Direction, pixels int) Resizer {
	return ResizeDirectionMode{"grow", direction, pixels}
}

func Shrink(direction Direction, pixels int) Resizer {
	return ResizeDirectionMode{"shrink", direction, pixels}
}

func Set(widthPixels, heightPixels int) Resizer {
	return ResizeDimensionMode{widthPixels, heightPixels}
}

func SetWidth(widthPixels int) Resizer {
	return ResizeDimensionMode{widthPixels, 0}
}

func SetHeight(heightPixels int) Resizer {
	return ResizeDimensionMode{0, heightPixels}
}
