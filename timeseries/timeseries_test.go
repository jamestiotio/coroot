package timeseries

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func slice2str(d []float32) string {
	values := make([]string, 0)
	for _, v := range d {
		values = append(values, Value(v).String())
	}
	return "[" + strings.Join(values, " ") + "]"
}

func TestLastN(t *testing.T) {
	ts := NewWithData(0, 1, []float32{0, 1, 2, NaN})
	assert.Equal(t, "TimeSeries(0, 4, 1, [0 1 2 .])", ts.String())
	assert.Equal(t, "[1 2 .]", slice2str(ts.LastN(3)))
	assert.Equal(t, "[0 1 2 .]", slice2str(ts.LastN(4)))
	assert.Equal(t, "[. . . 0 1 2 .]", slice2str(ts.LastN(7)))
}

func TestIncrease(t *testing.T) {
	x := NewWithData(0, 1, []float32{NaN, 1, 1, 1, 2, 2, 2, NaN, NaN, 10, NaN, 11, 12})
	status := NewWithData(0, 1, []float32{1, 1, 1, 1, 1, 1, 1, NaN, 1, 1, 0, 1, 1})
	assert.Equal(t, "TimeSeries(0, 13, 1, [. 1 0 0 1 0 0 . . 10 . . 1])", Increase(x, status).String())
}
