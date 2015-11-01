# Poker Spinner
A simulator for the poker "Spin & Go" format.

## Usage
- Copy `config.json.example` to `config.json`
- Set your parameters in `config.json`
- Run the binary in the same directory, e.g. `poker-spinner.exe`
- Set your win rate and how many runs you want to simulate as arguments

Example run:
```
$ .\poker-spinner.exe --winrate 40.5 --runs 10000
```

For a win rate of 40.5 percent and 10,000 runs (each with as many tournaments as
specified in the config).

## Configuration
- `tournaments` defines the number of simulated tournaments per run.
 - The ROI per run is calculated from the total of these tournaments.

- `payouts` are all possible multipliers and their payouts.
 - The probability is given over 1,000,000. So a probability of `0` to `3` actually means `0.000003`.
 - The difference between `probabilityStart` and `probabilityEnd` is the actual probability over one million.
 - The start of one probability must be equal to the end of another probability (unless it is `0`).
 - Probabilities can not "overlap" in start to end!
 - `win` and `lose` are the payout multipliers for winning and not winning (second and third place).
 - If `lose` is not given, they are set to `0`.