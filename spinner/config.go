// This file is part of the Poker Spinner package.
// Copyright (c) 2015 Martin Schenck
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package spinner

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration holds the overall configuration settings
type Configuration struct {
	WinRate     float64
	Runs        int
	Rounding    int
	Tournaments int
	Payouts     []Payout
	Denominator float64
}

// Payout holds the probability of this payout being picked and the multiplier
// for winning or losing
type Payout struct {
	Probability float64
	Win         int
	Lose        int
}

// ReadConfig reads the configuration from a file into the struct
func ReadConfig(filename string) (Configuration, error) {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	err = validateConfig(config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// validateConfig validates the read configuration to be correct
func validateConfig(config Configuration) error {
	sum := 0.0
	for _, payout := range config.Payouts {
		sum += payout.Probability
	}

	if sum != config.Denominator {
		return fmt.Errorf(
			"Probabilities in config.json add up to %.0f, but should add up to denominator %.0f",
			sum,
			config.Denominator,
		)
	}

	return nil
}
