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
	max := wave.Max()
	min := wave.Min()
	absmax := wave.AbsMax()

	fmt.Printf("Max.     = %f\n", max)
	fmt.Printf("Min.     = %f\n", min)
	fmt.Printf("Abs.Max. = %f\n", absmax)
}
