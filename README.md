# Sum Calculator for JSON File

This program calculates the sum of numbers in a JSON file using parallel processing with goroutines.

## Requirements

- Go (Golang) installed on your system.

## Usage

1. Clone the repository or download the `sum_calculator.go` file.
2. Navigate to the directory containing the `sum_calculator.go` file in your terminal.
3. Run the program with the following command:

```bash
go run sum_calculator.go <input_file> <num_goroutines>
```

- `<input_file>`: Path to the JSON file containing an array of objects with "a" and "b" keys.
- `<num_goroutines>`: Number of goroutines to use for parallel processing.

### Example

To calculate the sum using 4 goroutines:

```bash
go run sum_calculator.go input.json 4
```

### Help

To display the usage instructions:

```bash
go run sum_calculator.go -h
```

or

```bash
go run sum_calculator.go --help
```

## JSON Input Format

The input JSON file should be an array of objects. Each object should contain two keys: "a" and "b", with integer values ranging from -10 to 10.

Example:

```json
[
    {
        "a": 1,
        "b": 3
    },
    {
        "a": 5,
        "b": -9
    },
    {
        "a": -2,
        "b": 4
    }
    // ... (you can add more objects here)
]
```