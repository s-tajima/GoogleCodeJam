package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	path := "./sample.in"
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}

	file, err := readFile(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Number of test cases.
	num, _ := strconv.Atoi(scanner.Text())
	for i := 1; i <= num; i++ {
		scanner.Scan()
		sol := solver{}
		err := sol.parse(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Case #%d: %s\n", i, sol.solve())
	}
}

func readFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
