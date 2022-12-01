# Advent of Code 2021
Just my solutions for [Advent of Code 2019](https://adventofcode.com/2020). Starting from day 4 because previous were in Elixir.

Uses [Cobra](https://github.com/spf13/cobra) as CLI framework.

Requirements installation:
~~~~
> go mod tidy
~~~~

(2021-11) How to run specific puzzle (input files included in /input-files):
~~~~
> go run main.go day*.go day <day:1-25> <part:1/2> <input-file>
~~~~


Because of TDD approach, tests are also included:
~~~~
> go test
~~~~
