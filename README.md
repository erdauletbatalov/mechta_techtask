# Sum Calculator for JSON File

This program calculates the sum of numbers in a JSON file.

## Input

The input is a JSON file in the form of an array of objects. Each object contains two keys: "a" and "b". The values of these keys are integers ranging from -10 to 10.

Example input:
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
    ...
]
```

The file contains 1,000,000 objects.

## Requirements

- The program should read the JSON file and compute the sum of all numbers.
- The computation should be parallelized using goroutines.
- The number of goroutines for parallel processing should be specified as a command-line argument when running the program.
- The program should print the total result to the console.
  
## Usage

To run the program, use the following command:

```bash
go run main.go -goroutines=<number_of_goroutines> -file=<path_to_json_file>
```

- `<number_of_goroutines>`: Number of goroutines for parallel processing.
- `<path_to_json_file>`: Path to the JSON file containing the data.

Example:

```bash
go run main.go -goroutines=10 -file=data.json
```

## Implementation

The program uses Go language and leverages goroutines for parallel processing. The sum is calculated by concurrently processing the array of objects from the JSON file.


# .docx file with description:

Написать программу для вычисления суммы чисел в файле

Дан json файл в виде массива объектов как показано ниже:
```
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
...
]
```

Количество объектов = 1,000,000
значения чисел в диапазоне [-10,10]


Программа должна считать файл и вычислить сумму всех чисел. Для этого она должна:
⁃ Нужно распараллелить вычисление по горутинам
⁃ Количество горутин, для параллельной обработки нужно получить при запуске программы через аргумент
⁃ Вывести общий результат в консоль
⁃ Опубликовать решение в github.com и предоставить доступ
