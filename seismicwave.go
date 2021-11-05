package seismicwave

import (
	"math"
)

type Wave struct {
	Name string
	Dt   float64
	Data []float64
}

func New() *Wave {
	return newWave()
}

func newWave() *Wave {
	p := new(Wave)
	return p
}

func (w *Wave) DT() float64 {
	return w.Dt
}

func (w *Wave) NData() int {
	return len(w.Data)
}

func (w *Wave) Length() float64 {
	return float64(len(w.Data)) * w.Dt
}

func (w *Wave) Max() float64 {
	max := w.Data[0]
	n := len(w.Data)
	for i := 0; i < n; i++ {
		if w.Data[i] > max {
			max = w.Data[i]
		}
	}
	return max
}

func (w *Wave) AbsMax() float64 {
	absMax := w.Data[0]
	n := len(w.Data)
	for i := 0; i < n; i++ {
		if math.Abs(w.Data[i]) > absMax {
			absMax = math.Abs(w.Data[i])
		}
	}
	return absMax
}

func (w *Wave) Min() float64 {
	min := w.Data[0]
	n := len(w.Data)
	for i := 0; i < n; i++ {
		if w.Data[i] < min {
			min = w.Data[i]
		}
	}
	return min
}
