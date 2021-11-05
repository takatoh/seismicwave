package seismicwave

import (
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
