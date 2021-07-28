package data

import (
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jszwec/csvutil"
)

var province []models.Province
var provinceNameKeyMap map[string]string = make(map[string]string)
var provinceIDKeyMap map[string]string = make(map[string]string)

func init() {
	province, provinceNameKeyMap, provinceIDKeyMap = getProvince()
}

func GetAllProvince() gin.HandlerFunc {
	return func(ginctx *gin.Context) {
		respond, err := json.Marshal(province)
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

func loopingCSVWithFilter(dec *csvutil.Decoder, modelsOf interface{}, filter func(interface{}) (interface{}, bool)) (interface{}, error) {

	var slices []interface{}

	for {
		var city models.City
		err := dec.Decode(&city)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if item, valid := filter(city); valid {
			slices = append(slices, item)
		}
	}

	if len(slices) == 0 {
		return nil, errors.New("NO DATA")
	}

	return slices, nil
}

func UnmarshallProvince() (*csvutil.Decoder, error) {
	return unmarshall(PROVINCE)
}

func UnmarshallCity(filter func(interface{}) (interface{}, bool)) ([]models.City, error) {

	dec, err := unmarshall(CITY)

	if err != nil {
		return nil, err
	}

	cities, err := loopingCSVWithFilter(dec, models.City{}, filter)
	if err != nil {
		return nil, err
	}

	return cities.([]models.City), nil

	// var cities []models.City

	// for {
	// 	var city models.City
	// 	err = dec.Decode(&city)
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	if correctcity, valid := filter(city); valid {
	// 		cities = append(cities, correctcity)
	// 	}
	// }

	// if len(cities) == 0 {
	// 	return nil, errors.New("NO DATA")
	// }

	// return cities, nil
}

func UnmarshallDistrict() (*csvutil.Decoder, error) {
	return unmarshall(DISTRICT)
}

func UnmarshallVillage() (*csvutil.Decoder, error) {
	return unmarshall(VILLAGE)
}

func readFile(filepath string) ([]byte, error) {

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return file, nil
}

//Province is saved in memory (see above at init()), so no need to unmarshall repeatedly
//this func return all provinces in slice, map with key is province name and value is provinceid, and map with key is province id and value is provincename
func getProvince() ([]models.Province, map[string]string, map[string]string) {

	dec, err := UnmarshallProvince()

	if err != nil {
		panic(err)
	}

	var provinces []models.Province
	var provinceNameKeyMap map[string]string = make(map[string]string)

	for {
		var province models.Province
		err = dec.Decode(&province)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		provinces = append(provinces, province)
		provinceNameKeyMap[province.Name] = province.ID
		provinceIDKeyMap[province.ID] = province.Name
	}

	return provinces, provinceNameKeyMap, provinceIDKeyMap
}

func GetProvinceIDByName(name *string) string {
	return provinceNameKeyMap[*name]
}
