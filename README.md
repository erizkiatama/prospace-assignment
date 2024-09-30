# Intergalactic Converter

## System Design Solution

I tried to make the solution as modular as possible. There are 3 core functionality here in the code, they are:

- `database`, module for keeping the all the mapping that we need, including unit to roman numerals and credits per currency.
- `calculator`, module for anything about calculation. Mainly responsible to convert the units given into a quantity that human could read and comparing two units given whether its greater, less, or equals.
- `parser`, module for parsing the input into its respective business logic. The input is restricted and limited that will be explained further below.

There is also `main` module for accepting the input and main business logic after parsing the input.

I made the `database` and `calculator` module with an interface, introducing loose coupling and high cohesion in the codebase. This makes the code more modular and easier to maintain and test.
The `parser` is not need for any dependencies, so we could made it with no interfaces.

## Limit and Restriction
There are several limits and restrictions for this solution.

- The inputs are limited to the defined words and sentences in the instruction below. As for now, we don't support free inputs. The structure of the inputs must be noted as fail to notice this will result on invalid input.
- The units and credits define in the inputs will be reset when the program exited. As for now we don't need any persistence database for this solution.
- The credits will always be a floating point with 2 digits after point.
- The queries are assumed just like the sample inputs, so I create the instructions following that.

## Instructions

### Inputs
- For assigning:
  - Units to romans -> `{unit} is {roman}`, where `{unit}` is the units you want to assign and `{roman}` is the roman numeral.
  - Units with currency credits -> `{units} {currency} is {total} Credits`, where `{units}` are the units you want to assign and `{currency}` is the currency and `{total}` is the total credits.
- For calculating:
  - Roman numerals -> `how much is {units} ?`, where `{units}` is the units you want to calculate.
  - Credits -> `how many Credits is {units} {currency}`, where `{units}` is the units and `{currency}` is the currency you want to calculate.
- For comparing:
  - Roman numerals -> `Is {firstUnits} larger than {secondUnits}?`, where `{firstUnits}` and `{secondUnits}` are the units you want to compare.
  - Credits -> `Does {firstUnits} {firstCurrency} has more Credits than {secondUnits} {secondCurrency} ?`, where `{firstUnits}` and `{secondUnits}` are the units and `{firstCurrency}` and `{secondCurrency}` are the currencies you want to compare.

We provide the sample input on `test.txt` if you need. Please be aware of words per words because failing to follow this input instructions will make your inputs invalid.

### How to run the program
- Please install go first if you haven't yet, version 1.22.0 if you could as the `go.mod` is using that version
- Run `go build -o intergalactic-converter`
- Then run `./intergalactic-converter` or `./intergalactic-converter < test.txt` if using the text file. You could change the `test.txt` with your text file.
- Start inputting the query, if you are not using txt files, the output will be shown after you input an empty line `""` or just press enter when empty.
- Done
