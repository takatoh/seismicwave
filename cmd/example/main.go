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

	ns := waves[0]
	ew := waves[1]
	ud := waves[2]
	dt := ns.Dt
	n := len(ns.Data)

	t := 0.0
	fmt.Printf("%s,%s,%s,%s\n", "Time", ns.Name, ew.Name, ud.Name)
	for i := 0; i < n; i++ {
		fmt.Printf("%f,%f,%f,%f\n", t, ns.Data[i], ew.Data[i], ud.Data[i])
		t += dt
	}
}
