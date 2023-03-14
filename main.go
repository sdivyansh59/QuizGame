package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
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
	inputCh := make(chan string)

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

		go func ()  {
			fmt.Printf("Enter your Response: ")
			_,err := fmt.Scan(&input)
			if err != nil {
				log.Fatal(err)
			}
			inputCh <-input
			
		}()
		
		

		select{

		case userRes := <-inputCh:
			// check user response
			if userRes == rec[1] {
				correct++
			}else {
				incorrect++
			}

		case <- time.After(time.Second*8):
			fmt.Println("Sorry: time out")

		}
	
		counter++
	}

	fmt.Printf("Total Question %d Correct: %d  Incorrect: %d Skiped: %d\n",counter, correct, incorrect, (counter-(correct+incorrect)))

}


/* My reference 
1. How to read csv file : https://gosamples.dev/read-csv/



*/