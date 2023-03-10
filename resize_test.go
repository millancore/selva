package selva_test

import (
	"fmt"
	"github.com/millancore/selva"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var directions = []selva.Direction{
	selva.Left,
	selva.Right,
	selva.Up,
	selva.Down,
}

func TestGrow(t *testing.T) {
	assert := assert.New(t)

	var grow string

	for _, direction := range directions {
		pixels := rand.Intn(100)
		grow = selva.Grow(direction, pixels).Resize()

		assert.Equal(fmt.Sprintf("resize grow %s %dpx", direction, pixels), grow)
	}
}

func TestShrink(t *testing.T) {
	assert := assert.New(t)

	var shrink string

	for _, direction := range directions {
		pixels := rand.Intn(100)
		shrink = selva.Shrink(direction, pixels).Resize()

		assert.Equal(fmt.Sprintf("resize shrink %s %dpx", direction, pixels), shrink)
	}

}

func TestSet(t *testing.T) {
	assert := assert.New(t)
	resize := selva.Set(300, 400).Resize()

	assert.Equal("resize set width 300px height 400px", resize)
}

func TestSetWidth(t *testing.T) {
	assert := assert.New(t)
	resize := selva.SetWidth(500).Resize()

	assert.Equal("resize set width 500px", resize)
}

func TestSetHeight(t *testing.T) {
	assert := assert.New(t)
	resize := selva.SetHeight(10).Resize()

	assert.Equal("resize set height 10px", resize)
}
