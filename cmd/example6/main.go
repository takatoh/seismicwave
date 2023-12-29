package main

import (
	"fmt"
	"os"

	"github.com/takatoh/seismicwave"
)

func main() {
	filename := os.Args[1]

	waves, err := seismicwave.LoadCSV(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wave := waves[0]
	max, maxTime := wave.MaxWithTime()
	min, minTime := wave.MinWithTime()
	absmax, absmaxTime := wave.AbsMaxWithTime()
	ndata := wave.NData()
	length := wave.Length()

	fmt.Printf("Max.     = %f   Time = %f\n", max, maxTime)
	fmt.Printf("Min.     = %f   Time = %f\n", min, minTime)
	fmt.Printf("Abs.Max. = %f   Time = %f\n", absmax, absmaxTime)
	fmt.Printf("NData    = %d\n", ndata)
	fmt.Printf("Length   = %f sec\n", length)
}
