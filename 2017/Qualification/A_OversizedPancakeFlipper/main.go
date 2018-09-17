package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	path := "./sample.in"
	if len(os.Args) > 2 {
		path = os.Args[1]
	}

	file, err := readFile(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	_ = scanner.Scan() // Skip number of test cases.
	for scanner.Scan() {
		sol := solver{}
		err := sol.parse(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sol.solve()
	}
}

func readFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
