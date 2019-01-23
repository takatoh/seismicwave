package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/takatoh/seismicwave"
)

func main() {
	opt_name := flag.String("wavename", "", "Specify wave name.")
	opt_format := flag.String("format", "", "Wave format.")
	opt_dt := flag.Float64("dt", 0.0, "dt.")
	opt_ndata := flag.Int("ndata", 0, "Number of data.")
	opt_skip := flag.Int("skip", 0, "Skip lines.")
	flag.Parse()

	filename := flag.Args()[0]

	waves, err := seismicwave.LoadFixedFormat(
		filename,
		*opt_name,
		*opt_format,
		*opt_dt,
		*opt_ndata,
		*opt_skip,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wave := waves[0]
	fmt.Printf("Name     = %s\n", wave.Name)
	fmt.Printf("Dt       = %f\n", wave.Dt)
	fmt.Printf("Max.     = %f\n", wave.Max())
	fmt.Printf("Min.     = %f\n", wave.Min())
	fmt.Printf("Abs.Max. = %f\n", wave.AbsMax())
	fmt.Printf("NData    = %d\n", wave.NData())
	fmt.Printf("Length   = %f sec\n", wave.Length())
}
