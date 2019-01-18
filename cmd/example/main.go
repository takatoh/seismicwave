package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/takatoh/seismicwave"
)

func main() {
	opt_jma := flag.Bool("jma", false, "Load JMA waves.")
	opt_knet := flag.Bool("knet", false, "Load KNET waves.")
	flag.Parse()

	filename := flag.Args()[0]
	var waves []*seismicwave.Wave
	var err error
	if *opt_jma {
		waves, err = seismicwave.LoadJMA(filename)
	} else if *opt_knet {
		waves, err = seismicwave.LoadKNETSet(filename)
	} else {
		waves, err = seismicwave.LoadCSV(filename)
	}
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
