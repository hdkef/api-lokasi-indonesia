package province

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type provincehandler struct {
}

func GetProvinceHandler() *provincehandler {
	return &provincehandler{}
}

var _province []models.Province
var _provinceNameKeyMap map[string]models.Province = make(map[string]models.Province)
var _provinceIDKeyMap map[string]models.Province = make(map[string]models.Province)

func init() {
	initProvince()
}

//initProvince is to unmarshall all provinces and saved those in memory,
//there are variable containes slices of all provinces, map with province name key and province id value and map with vice versa
func initProvince() {

	provinces, err := data.UnmarshallProvince(func(i interface{}) (interface{}, bool, bool) {
		province := i.(models.Province)
		_provinceNameKeyMap[province.Name] = province
		_provinceIDKeyMap[province.ID] = province
		return i, true, false
	})

	if err != nil {
		panic(err)
	}
	_province = provinces
}

func GetProvinceByName(name *string) models.Province {
	return _provinceNameKeyMap[*name]
}

func GetProvinceByID(id *string) models.Province {
	return _provinceIDKeyMap[*id]
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

//Endpoint for get province detail by id
func GetProvinceByIDHandler(value *string, ginctx *gin.Context) {
	provinceFound := GetProvinceByID(value)

	respond, err := json.Marshal(provinceFound)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}

//Endpoint for get province detail by name
func GetProvinceByNameHandler(value *string, ginctx *gin.Context) {
	provinceFound := GetProvinceByName(value)

	respond, err := json.Marshal(provinceFound)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}
