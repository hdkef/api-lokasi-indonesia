package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jszwec/csvutil"
)

func csvPath(fname string) string {
	return filepath.Join("data", fname)
}

var PROVINCE = csvPath("provinces.csv")
var DISTRICT = csvPath("districts.csv")
var CITY = csvPath("cities.csv")
var VILLAGE = csvPath("villages.csv")

func openFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func unmarshall(filepath string) (*csvutil.Decoder, error) {
	file, err := openFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return dec, nil
}

func UnmarshallCity() (*csvutil.Decoder, error) {
	return unmarshall(CITY)
}

func UnmarshallDistrict() (*csvutil.Decoder, error) {
	return unmarshall(DISTRICT)
}

func UnmarshallVillage() (*csvutil.Decoder, error) {
	return unmarshall(VILLAGE)
}
