# Poker Spinner
A simulator for the poker "Spin & Go" format.

It runs `x` tournaments per run and `y` runs. So a total of `x * y` tournaments.
Each run has a certain return of investemnt (ROI) over all tournaments in that run.
The ROIs of all runs are grouped and printed when the simulation is done.

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
- `rounding` defines the precision for the ROI outputs
- `denominator` only exists to check the config for correctness. It has to equal the sum of all `payouts`' probabilities
- `payouts` are all possible multipliers and their payouts.
 - The probability is given over the `denominator`. So a probability of `3` actually means `3/denimonator`.
 - `win` and `lose` are the payout multipliers for winning and not winning (second and third place).
 - If `lose` is not given, it is set to `0`.