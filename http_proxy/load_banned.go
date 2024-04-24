package http_proxy

import (
	"encoding/csv"
	"log"
	"os"
)

func readCSV(filename string) (map[string]bool, error) {
	file, ferr := os.Open(filename)
	if ferr != nil {
		return nil, ferr
	}
	r := csv.NewReader(file)

	websites, rerr := r.ReadAll()
	if rerr != nil {
		return nil, rerr
	}
	var values map[string]bool = make(map[string]bool)
	for i := 1; i < len(websites); i++ {
		values[websites[i][0]] = true
	}
	return values, nil
}

func LoadBanned(filename string) map[string]bool {
	values, err := readCSV(filename)
	if err != nil {
		log.Fatalln(err)
	}
	return values
}
