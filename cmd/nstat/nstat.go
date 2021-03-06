package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/sparrc/nstat"
)

var (
	// Version can be auto-set at build time using an ldflag
	//   go build -ldflags "-X main.Version `git describe --tags --always`"
	Version string
)

func main() {
	c := &nstat.Counters{}

	// parse command-line args
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-z", "--zero":
			c.DumpZeros = true
		case "-v", "-V", "--version":
			fmt.Println("nstat - " + Version)
			return
		case "-h", "--help":
			usage()
			return
		}
	}

	counters := c.Get()
	sortedKeys := []string{}
	for key, _ := range counters {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		fmt.Printf("%-32s %d\n", key, counters[key])
	}
}

func usage() {
	fmt.Print(`NAME
       nstat - network statistics tool.

SYNOPSIS
       Usage: nstat [ options ]

DESCRIPTION
       nstat is a simple tool to monitor kernel snmp counters and network interface statistics.

OPTIONS
       -h, --help
	          Print help

       -V, --version
              Print version

       -z, --zero
              Dump zero counters too. By default they are not shown.
`)
}
