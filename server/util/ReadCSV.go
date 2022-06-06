package util

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	if _, err := csvReader.Read(); err != nil {
		panic(err)
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}
