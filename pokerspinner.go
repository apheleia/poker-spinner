// This file is part of the Poker Spinner package.
// Copyright (c) 2015 Martin Schenck
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"flag"
	"fmt"
	"github.com/apheleia/poker-spinner/spinner"
	"runtime"
	"time"
)

// Reads the config, parses the flags, and runs the simulation
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	startTime := time.Now()

	config, err := spinner.ReadConfig("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	parseFlags(&config)
	spinner.Spin(config)

	fmt.Printf("Runtime: %.3fs", time.Now().Sub(startTime).Seconds())
	fmt.Println()
}

// parseFlags parses the flags from the commnd line and sets the config
// accordingly
func parseFlags(config *spinner.Configuration) {
	winRate := flag.Float64("winrate", 0.0, "the win rate to use for the simulation")
	runs := flag.Int("runs", 10, "how often the simulation should be run")
	flag.Parse()

	config.WinRate = *winRate
	config.Runs = *runs
}
