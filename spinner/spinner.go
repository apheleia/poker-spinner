// This file is part of the Poker Spinner package.
// Copyright (c) 2015 Martin Schenck
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package spinner

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

// Spin runs the simulation a number of times with a given configuratin
func Spin(config Configuration) {
	fmt.Println("Threads:            ", runtime.GOMAXPROCS(0))
	fmt.Println("Win rate:           ", config.WinRate, "%")
	fmt.Println("Runs:               ", config.Runs)
	fmt.Println("Tournaments per run:", config.Tournaments)
	fmt.Println("Rounding:           ", config.Rounding)
	fmt.Println()

	c := make(chan float64)
	defer close(c)

	startRuns(config, c)
	rois := gatherRunResults(config, c)

	print(rois, config.Runs, config.Rounding)
}

// startRuns starts runs a number of threads times
func startRuns(config Configuration, c chan<- float64) {
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		go run(config, c)
	}
}

// gatherRunResults reads all results from the given channel and accumulates
// them
func gatherRunResults(config Configuration, c <-chan float64) map[float64]int {
	rois := make(map[float64]int)

	for i := 0; i < runtime.GOMAXPROCS(0)*(config.Runs/runtime.GOMAXPROCS(0)); i++ {
		roi := <-c
		rois[roi] = rois[roi] + 1
	}

	return rois
}

// run starts a number fo runs
//
// runs starts x runs, where x is the number of total runs deivided by the
// number of threads, because run is called `runtime.GOMAXPROCS(0)` times
// that totales the total number of runs
//
// each run simulates `config.Tournaments` tournaments
//
// the results are written to the channel
func run(config Configuration, c chan<- float64) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := 0; i < config.Runs/runtime.GOMAXPROCS(0); i++ {
		multiplier := 0
		for j := 0; j < config.Tournaments; j++ {
			multiplier += tournament(&config, rng)
		}

		c <- round(float64(multiplier)/float64(config.Tournaments), config.Rounding)
	}
}

// tournament simulates a single tournament
func tournament(config *Configuration, rng *rand.Rand) int {
	payout, err := selectPayout(&config.Payouts, config.Denominator, rng)

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	if isWin(config.WinRate, rng) {
		return payout.Win
	}

	return payout.Lose
}

// selectPayout selects one of the configured payouts randomly
//
// the probability of a payout being selected equals that configured in the
// config.json file
func selectPayout(payouts *[]Payout, denominator float64, rng *rand.Rand) (Payout, error) {
	random := rng.Float64()

	lower := 0.0
	for _, payout := range *payouts {
		upper := lower + payout.Probability/denominator
		if lower <= random && random < upper {
			return payout, nil
		}

		lower = upper
	}

	return Payout{}, fmt.Errorf("No suitable payout found for probability %.10f", random)
}

// isWin returns whether or not the player wins this tournament
func isWin(winRate float64, rng *rand.Rand) bool {
	return 100*rng.Float64() < winRate
}

// round rounds a float to the given precision
func round(f float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Floor(f*shift+.5) / shift
}

// print prints all results to stdout
func print(rois map[float64]int, runs, precision int) {
	format := fmt.Sprintf("%%.%df: %%.7f", precision)

	fmt.Println("ROI: probability")
	for _, k := range sortedKeys(rois) {
		fmt.Printf(format, k, float64(rois[k])/float64(runs))
		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("Total ROI: %.7f", getTotalRoi(rois, runs))
	fmt.Println()
}

// sortedKeys returns all keys of the given map sorted
func sortedKeys(m map[float64]int) []float64 {
	var keys []float64
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	return keys
}

// getTotalRoi calculates the average ROI over all runs
func getTotalRoi(rois map[float64]int, runs int) float64 {
	total := 0.0
	for key, value := range rois {
		total += float64(key) * (float64(value) / float64(runs))
	}

	return total
}
