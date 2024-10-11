package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type housingRecord struct {
	//Creating a housing record struct for storing CSV lines as JSON struct field tags
	Value    float64 `json:"value"`
	Income   float64 `json:"income"`
	Age      int     `json:"age"`
	Rooms    int     `json:"rooms"`
	Bedrooms int     `json:"bedrooms"`
	Pop      int     `json:"pop"`
	Hh       int     `json:"hh"`
}

func createHousingList(data [][]string) []housingRecord {
	//convert csv lines to array of structs
	var housingList []housingRecord
	for i, line := range data {
		if i > 0 { //skips line 0, the header
			var rec housingRecord
			for j, field := range line {
				if j == 0 {
					var err error
					rec.Value, err = strconv.ParseFloat(field, 64)
					if err != nil {
						fmt.Printf("Error converting value in line %d: ", i+1)
						fmt.Println(err)
						continue
					}
				} else if j == 1 {
					var err error
					rec.Income, err = strconv.ParseFloat(field, 64)
					if err != nil {
						fmt.Printf("Error converting income in line %d: ", i+1)
						fmt.Println(err)
						continue
					}
				} else if j == 2 {
					var err error
					rec.Age, err = strconv.Atoi(field)
					if err != nil {
						fmt.Printf("Error converting age in line %d: ", i+1)
						fmt.Println(err)
						continue
					}
				} else if j == 3 {
					var err error
					rec.Rooms, err = strconv.Atoi(field)
					if err != nil {
						fmt.Printf("Error converting rooms in line %d: ", i+1)
						fmt.Println(err)
						continue
					}
				} else if j == 4 {
					var err error
					rec.Bedrooms, err = strconv.Atoi(field)
					if err != nil {
						fmt.Printf("Error converting bedrooms in line %d: ", i+1)
						fmt.Println(err)
						continue
					}
				} else if j == 5 {
					var err error
					rec.Pop, err = strconv.Atoi(field)
					if err != nil {
						fmt.Printf("Error converting pop in line %d: ", i+1)
						fmt.Println(err)
						continue
					}
				} else if j == 6 {
					var err error
					rec.Hh, err = strconv.Atoi(field)
					if err != nil {
						fmt.Printf("Error converting jjin line %d: ", i+1)
						fmt.Println(err)
						continue
					}
				}
			}
			housingList = append(housingList, rec)
		}
	}
	return housingList
}

// Returns true if filename already exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func argChecks(inputArg string, outputArg string) (string, string) {
	absOutput, err := filepath.Abs(outputArg)
	if err != nil {
		log.Fatal(err)
	}

	absInput, err := filepath.Abs(inputArg)
	if err != nil {
		log.Fatal(err)
	}

	//Check whether or not the given input filename actually exists and the output also exists.
	if fileExists(outputArg) {
		log.Fatal(absOutput, " already exists.\nPlease choose a different output file name.")
	}

	if !fileExists(inputArg) {
		log.Fatal("Unable to find ", absInput, "\n please ensure your file is named correctly")
	}
	return absInput, absOutput
}

func main() {

	if len(os.Args) != 3 {
		log.Fatal("Please ensure there are 2 arguments provided. The first argumment should be your csv file (including .csv). The second argument should be the desired JSON filename (including.json)")
	}

	inputArg := os.Args[1]
	outputArg := os.Args[2]

	_, absOutput := argChecks(inputArg, outputArg)

	//Open input csv file
	f, err := os.Open(inputArg)
	if err != nil {
		log.Fatal(err)
	}

	//defer closing the csv file
	defer f.Close()

	//Read CSV file through csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//Take data and create a list of structs
	housingList := createHousingList(data)
	//fmt.Printf("%+v\n", housingList)

	//convert the array of structs to JSON
	jsonData, err := json.MarshalIndent(housingList, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//Create JSON file to write data to
	jsonFile, err := os.Create(outputArg)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

	fmt.Print("Successfully converted CSV to JSON. The file is in ", absOutput)
}
