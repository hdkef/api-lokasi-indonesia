package data

import (
	"api-lokasi-indonesia/konstant"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jszwec/csvutil"
)

var _province []models.Province
var _provinceNameKeyMap map[string]string = make(map[string]string)
var _provinceIDKeyMap map[string]string = make(map[string]string)

func init() {
	initProvince()
}

//Endpoint for get all provinces
func GetAllProvince() gin.HandlerFunc {
	return func(ginctx *gin.Context) {
		respond, err := json.Marshal(_province)
		if err != nil {
			utils.ResErr(ginctx, http.StatusInternalServerError, err)
			return
		}

		ginctx.Writer.Write(respond)
	}
}

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

func unmarshall(unitType string, filepath string, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {
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

	result, err := loopingCSVWithFilter(unitType, dec, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//will match the loop csv handle with unit type BECAUSE GOLANG HAS NO GENERIC AND IT SUCKS
func loopingCSVWithFilter(unitType string, dec *csvutil.Decoder, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {

	switch unitType {
	case konstant.Province:
		return doLoopingProvince(dec, filter)
	case konstant.City:
		return doLoopingCity(dec, filter)
	case konstant.District:
		return doLoopingDistrict(dec, filter)
	case konstant.Village:
		return doLoopingVillage(dec, filter)
	default:
		return nil, errors.New("NO matching unit type")
	}

}

func UnmarshallProvince(filter func(interface{}) (interface{}, bool, bool)) ([]models.Province, error) {
	result, err := unmarshall(konstant.Province, PROVINCE, filter)
	if err != nil {
		return nil, err
	}
	return result.([]models.Province), nil
}

func UnmarshallCity(filter func(interface{}) (interface{}, bool, bool)) ([]models.City, error) {

	result, err := unmarshall(konstant.City, CITY, filter)

	if err != nil {
		return nil, err
	}

	return result.([]models.City), nil
}

func UnmarshallDistrict(filter func(interface{}) (interface{}, bool, bool)) ([]models.District, error) {

	result, err := unmarshall(konstant.District, DISTRICT, filter)

	if err != nil {
		return nil, err
	}

	return result.([]models.District), nil
}

func UnmarshallVillage(filter func(interface{}) (interface{}, bool, bool)) ([]models.Village, error) {

	result, err := unmarshall(konstant.Village, VILLAGE, filter)

	if err != nil {
		return nil, err
	}

	return result.([]models.Village), nil
}

//initProvince is to unmarshall all provinces and saved those in memory,
//there are variable containes slices of all provinces, map with province name key and province id value and map with vice versa
func initProvince() {

	provinces, err := UnmarshallProvince(func(i interface{}) (interface{}, bool, bool) {
		province := i.(models.Province)
		_provinceNameKeyMap[province.Name] = province.ID
		_provinceIDKeyMap[province.ID] = province.Name
		return i, true, false
	})

	if err != nil {
		panic(err)
	}
	_province = provinces
}

func GetProvinceIDByName(name *string) string {
	return _provinceNameKeyMap[*name]
}
