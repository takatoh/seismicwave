package seismicwave

import (
	"encoding/csv"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

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
		wave.Name = strings.TrimSpace(row[i])
		waves = append(waves, wave)
	}
	for {
		row, err = reader.Read()
		if err == io.EOF {
			dt := round(t2-t1, 2)
			for i := 0; i < n; i++ {
				waves[i].Dt = dt
			}
			return waves, nil
		}
		t1 = t2
		t2, _ = strconv.ParseFloat(strings.TrimSpace(row[0]), 64)
		for i := 1; i <= n; i++ {
			d, e := strconv.ParseFloat(strings.TrimSpace(row[i]), 64)
			if e == nil {
				waves[i-1].Data = append(waves[i-1].Data, d)
			}
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
