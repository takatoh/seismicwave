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

func Make(name string, dt float64, data []float64) *Wave {
	p := new(Wave)
	p.Name, p.Dt, p.Data = name, dt, data
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
	max, _ := w.MaxWithTime()
	return max
}

func (w *Wave) MaxWithTime() (float64, float64) {
	max := w.Data[0]
	n := len(w.Data)
	t := 0.0
	for i := 0; i < n; i++ {
		if w.Data[i] > max {
			max = w.Data[i]
			t = w.Dt * float64(i)
		}
	}
	return max, t
}

func (w *Wave) AbsMax() float64 {
	absMax, _ := w.AbsMaxWithTime()
	return absMax
}

func (w *Wave) AbsMaxWithTime() (float64, float64) {
	absMax := w.Data[0]
	n := len(w.Data)
	t := 0.0
	for i := 0; i < n; i++ {
		if math.Abs(w.Data[i]) > absMax {
			absMax = math.Abs(w.Data[i])
			t = w.Dt * float64(i)
		}
	}
	return absMax, t
}

func (w *Wave) Min() float64 {
	min, _ := w.MinWithTime()
	return min
}

func (w *Wave) MinWithTime() (float64, float64) {
	min := w.Data[0]
	n := len(w.Data)
	t := 0.0
	for i := 0; i < n; i++ {
		if w.Data[i] < min {
			min = w.Data[i]
			t = w.Dt * float64(i)
		}
	}
	return min, t
}
