package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//fmt.Println("Quiz Game!")
    
	// open file 
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	//close the file at the end 
	defer f.Close()

	// read csv value using csv.Reader()
	csvReader := csv.NewReader(f)

	var correct , incorrect , counter int
	
	for {
		rec,err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// do something  with read line
		//fmt.Printf("%+v\n", rec)
		fmt.Printf("Question No: %d  %v \n",counter+1, rec[0])
		var input  string

		fmt.Printf("Enter your Response: ")
		fmt.Scan(&input)
	
		// check 
		if input == rec[1] {
			correct++
		}else {
			incorrect++
		}

		counter++
	}

	fmt.Printf("Your score\n Correct: %d  Incorrect: %d \n", correct, incorrect)




}


/* My reference 
1. How to read csv file : https://gosamples.dev/read-csv/



*/