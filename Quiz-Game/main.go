package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	score int
	numberProblem int
) 

func main(){
	fileName := flag.String("csv", "problem.csv", `a csv file in the format of 'question,answer' (default "problem.csv") `)
	flag.Parse()

    file, err := os.Open(*fileName)
    if err != nil {

        log.Fatal(err)
    }

    r := csv.NewReader(file)

	for {
        record, err := r.Read()
		numberProblem++

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatal(err)
        }
		
		if answer := runTest(record); answer == true {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", score, numberProblem)
}

func runTest(record []string) bool{
	var inputUser string

	fmt.Printf("Problem #%d: %s = ", numberProblem, record[0])
	fmt.Scanln(&inputUser)

	if inputUser == record[1]{
		
		return true
	}

	return false
}

