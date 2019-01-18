package seismicwave

import (
	"encoding/csv"
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"regexp"
	"math"
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

func LoadCSV(filename string) ([]*Wave, error) {
	var waves []*Wave
	var reader *csv.Reader
	var row []string
	var err error
	var t1, t2 float64

	t1 = 0.0
	t2 = 0.0

	read_file, err := os.Open(filename)
	if err != nil {
		return waves, err
	}
	defer read_file.Close()

	reader = csv.NewReader(read_file)
	row, err = reader.Read()
	n := len(row) - 1
	for i := 1; i <= n; i++ {
		wave := newWave()
		wave.Name = row[i]
		waves = append(waves, wave)
	}
	for {
		row, err = reader.Read()
		if err == io.EOF {
			dt := round(t2 - t1, 2)
			for i := 0; i < n; i++ {
				waves[0].Dt = dt
			}
			return waves, nil
		}
		t1 = t2
		t2, _ = strconv.ParseFloat(row[0], 64)
		for i := 1; i <= n; i++ {
			d, _ := strconv.ParseFloat(row[i], 64)
			waves[i - 1].Data = append(waves[i - 1].Data, d)
		}
	}
}

func round(val float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}

func LoadKNETSet(basename string) ([]*Wave, error) {
	var waves []*Wave
	var dirs = []string{ "NS", "EW", "UD" }

	for _, dir := range dirs {
		ws, err := LoadKNET(basename + "." + dir)
		if err != nil {
			return waves, err
		}
		waves = append(waves, ws[0])
	}

	return waves, nil
}

func LoadKNET(filename string) ([]*Wave, error) {
	var dt float64
	var scaleFactor float64
	wave := newWave()
	data := make([]float64, 0)

	f, err := os.Open(filename)
	if err != nil {
		return []*Wave{ wave }, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "Memo") == 0 {
			break
		}
		if strings.Index(line, "Sampling Freq(Hz)") == 0 {
			s := regexp.MustCompile(" +").Split(line, 3)
			s2 := strings.Trim(s[2], "Hz")
			f, _ := strconv.ParseFloat(s2, 64)
			dt = 1.0 / f
		}
		if strings.Index(line, "Scale Factor") == 0 {
			s := regexp.MustCompile(" +").Split(line, 3)
			s2 := regexp.MustCompile(`\(gal\)/`).Split(s[2], 2)
			f1, _ := strconv.ParseFloat(s2[0], 64)
			f2, _ := strconv.ParseFloat(s2[1], 64)
			scaleFactor = f1 / f2
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		dataStrings := regexp.MustCompile(" +").Split(line, 8)
		for _, s := range dataStrings {
			d, _ := strconv.ParseFloat(s, 64)
			data = append(data, d * scaleFactor)
		}
	}

	wave.Name = ""
	wave.Dt = dt
	wave.Data = data
	return []*Wave{ wave }, nil
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
