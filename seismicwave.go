package seismicwave

import (
	"bufio"
	"encoding/csv"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

type Wave struct {
	Name string
	Dt   float64
	Data []float64
}

func newWave() *Wave {
	p := new(Wave)
	return p
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

type InputWave struct {
	Waves []InputWaveInfo `toml:"wave"`
}

type InputWaveInfo struct {
	Name   string  `toml:"name"`
	File   string  `toml:"file"`
	Format string  `toml:"format"`
	Dt     float64 `toml:"dt"`
	NData  int     `toml:"ndata"`
	Skip   int     `toml:"skip"`
}

func LoadJMA(filename string) ([]*Wave, error) {
	var waves []*Wave
	var ns, ew, ud *Wave
	var dt float64
	var dataNS, dataEW, dataUD []float64
	var flg bool

	ns = newWave()
	ns.Name = "NS"
	ew = newWave()
	ew.Name = "EW"
	ud = newWave()
	ud.Name = "UD"

	f, err := os.Open(filename)
	if err != nil {
		return waves, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			ns.Dt = dt
			ns.Data = dataNS
			waves = append(waves, ns)
			ew.Dt = dt
			ew.Data = dataEW
			waves = append(waves, ew)
			ud.Dt = dt
			ud.Data = dataUD
			waves = append(waves, ud)
			break
		}
		if flg {
			d0, _ := strconv.ParseFloat(row[0], 64)
			dataNS = append(dataNS, d0)
			d1, _ := strconv.ParseFloat(row[1], 64)
			dataEW = append(dataEW, d1)
			d2, _ := strconv.ParseFloat(row[2], 64)
			dataUD = append(dataUD, d2)
		}
		if strings.Index(row[0], " NS") == 0 {
			flg = true
		}
		if strings.Index(row[0], " SAMPLING RATE") == 0 {
			srb := regexp.MustCompile(`\d+`).Find([]byte(row[0]))
			sr, _ := strconv.ParseFloat(string(srb), 64)
			dt = 1.0 / sr
		}
	}

	return waves, nil
}

func LoadFixedFormat(filename, wavename, format string, dt float64, ndata, skip int) ([]*Wave, error) {
	wave := newWave()
	wave.Name = wavename
	wave.Dt = dt

	fstrings := regexp.MustCompile("[Ff.]").Split(format, 3)
	fn, _ := strconv.Atoi(fstrings[0])
	fl, _ := strconv.Atoi(fstrings[1])
	lineCount := ndata / fn
	if ndata%fn > 0 {
		lineCount += 1
	}
	var data []float64

	f, err := os.Open(filename)
	if err != nil {
		return []*Wave{wave}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for i := 0; i < skip; i++ {
		scanner.Scan()
	}
	for i := 0; i < lineCount; i++ {
		scanner.Scan()
		line := scanner.Text()
		datas := splitN(line, fl)
		for j, s := range datas {
			if j < fn {
				d, _ := strconv.ParseFloat(strings.Trim(s, " "), 64)
				data = append(data, d)
			}
		}
	}
	wave.Data = data

	return []*Wave{wave}, nil
}

func splitN(s string, l int) []string {
	var r []string

	for i := 0; i < len(s); i += l {
		if i+l < len(s) {
			r = append(r, s[i:(i+l)])
		} else {
			r = append(r, s[i:])
		}
	}

	return r
}

func LoadFixedFormatWithTOML(inputfile string) ([]*Wave, error) {
	var waves []*Wave

	var input InputWave
	_, err := toml.DecodeFile(inputfile, &input)
	if err != nil {
		return waves, err
	}

	for _, w := range input.Waves {
		ws, err := LoadFixedFormat(w.File, w.Name, w.Format, w.Dt, w.NData, w.Skip)
		if err != nil {
			return waves, err
		}
		waves = append(waves, ws[0])
	}

	return waves, nil
}
