package seismicwave

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LoadKNET(filename string) ([]*Wave, error) {
	var dt float64
	var scaleFactor float64
	wave := newWave()
	data := make([]float64, 0)

	f, err := os.Open(filename)
	if err != nil {
		return []*Wave{wave}, err
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
			data = append(data, d*scaleFactor)
		}
	}

	wave.Name = ""
	wave.Dt = dt
	wave.Data = data
	return []*Wave{wave}, nil
}

func LoadKNETSet(basename string) ([]*Wave, error) {
	var waves []*Wave
	var dirs = []string{"NS", "EW", "UD"}

	for _, dir := range dirs {
		ws, err := LoadKNET(basename + "." + dir)
		if err != nil {
			return waves, err
		}
		waves = append(waves, ws[0])
	}

	return waves, nil
}
