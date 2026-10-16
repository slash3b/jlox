package main

import (
	"bufio"
	"fmt"
	"os"
)

type Lox struct {
	hasError bool
}

// it is good engineering practice to separate the code
// that generates error from the code that reports error
func errorExit(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %q: %v\n", msg, err.Error())
	} else {
		fmt.Fprintf(os.Stderr, "err: %s\n", msg)
	}

	os.Exit(1)
}

func main() {
	fmt.Println("this is phlox language\n")

	switch len(os.Args) {
	case 2:
		b, err := os.ReadFile(os.Args[1])
		if err != nil {
			errorExit("could not read file", err)
		}

		sc := NewScanner(string(b))

		err = sc.Scan()
		if err != nil {
			errorExit("scan error", err)
		}

		fmt.Println("scan is done")
		for _, v := range sc.Tokens {
			fmt.Printf("token: %q %q %q\n", v.Type, v.Lexeme, v.Literal)
		}

		fmt.Println("conten length: ", len(string(b)))

		return
	case 1:
		sc := bufio.NewScanner(os.Stdin)
		err := runLine(sc)
		if err != nil {
			errorExit("scan err", err)
		}

		return
	}

	errorExit("exactly 1 or no argument is expected", nil)
}

func runLine(sc *bufio.Scanner) error {
	fmt.Print("> ")

	for sc.Scan() {
		if sc.Err() != nil {
			return sc.Err()
		}

		fmt.Println(sc.Text())
		fmt.Print("> ")
	}

	return nil
}
