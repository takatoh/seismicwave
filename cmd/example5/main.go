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

	for _, wave := range waves {
		fmt.Printf("Name     = %s\n", wave.Name)
		fmt.Printf("Dt       = %f\n", wave.Dt)
		fmt.Printf("Max.     = %f\n", wave.Max())
		fmt.Printf("Min.     = %f\n", wave.Min())
		fmt.Printf("Abs.Max. = %f\n", wave.AbsMax())
		fmt.Printf("NData    = %d\n", wave.NData())
		fmt.Printf("Length   = %f sec\n\n", wave.Length())
	}
}
