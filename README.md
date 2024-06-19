# PrintBits

* printbits is program that takes a single argument (a number) and prints its binary representation without a newline at the end. 
* If the argument is not a number or if the number of arguments is not exactly one, the program prints 00000000.

## Usage

* To run the program, use the following command:

```bash
$ go run . <number>
```

### Examples

```bash
$ go run . 1
00000001$
```

```bash
$ go run . 192
11000000$
```

```bash
$ go run . a
00000000$
```

```bash
$ go run . 1 1
$
```

```bash
$ go run .
$
```