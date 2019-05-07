package svgshapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointString(t *testing.T) {
	points := []Point{
		{    0.00,    0.00 },
		{    1.00,    1.00 },
		{ 1230.00, 1000.02 },
		{   12.34,   12.34 },
		{    0.01,    1.50 },
	}
	strs := []string{
		"0,0",
		"1,1",
		"1230,1000.02",
		"12.34,12.34",
		"0.01,1.5",
	}

	for i, point := range points {
		assert.Equal(t, strs[i], point.String())
	}
}

func TestMarshalEmptyPointList(t *testing.T) {
	list := PointList{}
	text, err := list.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(""), text)
}

func TestMarshalSinglePointList(t *testing.T) {
	list := PointList{
		Point{ 1, 2 },
	}
	text, err := list.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte("1,2"), text)
}

func TestMarshalMultiPointList(t *testing.T) {
	list := PointList{
		Point{ 1, 2 },
		Point{ 3, 4 },
		Point{ 5, 6 },
	}
	text, err := list.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte("1,2 3,4 5,6"), text)
}

