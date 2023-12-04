# advent-of-code
My solutions to the Advent of Code challenges, written in Go.

## Project Structure
The solutions are stored in their own directory with naming format `calendar/$YEAR/$DAY`.  
Each directory contains the following files:
- `input.txt` - puzzle input
- `main.go` - puzzle solutions (go source code)
- `README.md` - puzzle instructions
- `Makefile` - common operations such as building and running the puzzle solutions

## Automations
By referencing the root [Makefile](Makefile), you can execute the following commands:

To generate the folder structure for today:
```bash
make generate
```

To generate the folder structure for a specific date:
```bash
make generate YEAR=2024 DAY=1
```

To lint the Go source code:
```bash
make lint
```