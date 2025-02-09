# noisemixer

Mix random characters with the string provided. This tool takes an input string and intersperses it with random characters, making it more complex by adding random-length strings between each character of your input.

## Features

- Adds random characters between each character of your input string
- Uses cryptographically secure random generation for added security
- Configurable maximum length for the random strings
- Supports a wide range of characters including letters, numbers, and special characters

## Installation

To install noisemixer, you need to have Go installed on your system. Then:

```bash
go install github.com/jempe/noisemixer@latest
```

The binary will be installed to your `$GOPATH/bin` directory (usually `~/go/bin`). Make sure this directory is in your system's PATH.

## Usage

You can use noisemixer by piping text to it or typing directly when prompted:

```bash
# Using echo
echo "Hello" | ./noisemixer

# Or just run it and type your input
./noisemixer
```

### Options

- `-length`: Sets the maximum length of random strings inserted between characters (default: 25)
  ```bash
  echo "Hello" | ./noisemixer -length 10
  ```

## How It Works

1. The program reads input from stdin
2. For each character in the input:
   - Prints the original character
   - Generates a random number between 2 and the specified maximum length
   - Creates a random string of that length using a charset of letters, numbers, and special characters
   - Prints the random string
3. The process continues until all input characters have been processed

## Example

Input:
```
echo "Hi" | ./noisemixer
```

Possible output:
```
Hx#4kj!iK9p2i$mN8v
```

In this example, random characters were inserted between 'H' and 'i'.

## Building From Source

```bash
# Get the source code
git clone https://github.com/jempe/noisemixer.git

# Change to the project directory
cd noisemixer

# Build the project
go build

# Run the program
./noisemixer
```

## License

[Add your license information here]
