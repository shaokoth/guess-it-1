# guess-it-1
# Guess-It Statistics Calculator

## Overview
Given a number as standard input, this Go program reads from standard input (stdin) and calculates the following statistical metrics:

- **Mean**: The average of the provided numbers.
- **Standard Deviation**: A measure of the amount of variation or dispersion of the numbers.
- **Prediction Interval**: A statistical range where a future observation is likely to fall.

The results from the prediction interval are then rounded and displayed to the user.

## Features

- Reads input values from stdin.
- Computes the mean of the input values.
- Computes the standard deviation.
- Calculates the prediction interval based on the mean and standard deviation.

## Prerequisites

- **Go**: You need to have Go installed on your machine. You can download it from [the official website](https://golang.org/dl/).

## Installation

1. **Clone the repository** :

    ```bash
    git clone https://learn.zone01kisumu.ke/git/shaokoth/guess-it-1
    cd guess-it-1
    ```

3. **Run the program** using:

    ```bash
    go run .
    ```

## Usage

To use this program:
1. Move into student folder:
    ```
    cd student
    ```

1. Run the program:

    ```bash
    go run .
    ```

2. Provide floating-point numbers as input by typing the numbers directly into the terminal:
    ```bash
    go run .
    ```

    and then enter the numbers one by one, pressing `Enter` after each.

## Error Handling

- If you input a non-numeric value, the program will output an error message and terminate.
- If no data is provided, the program will notify you that there was an error parsing data.

## Contributing

If you want to contribute to this project, feel free to submit a pull request or open an issue on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
