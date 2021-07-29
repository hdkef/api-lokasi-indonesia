package forwarder

import (
	"api-lokasi-indonesia/city"
	"api-lokasi-indonesia/district"
	"api-lokasi-indonesia/konstant"
	"api-lokasi-indonesia/province"
	"api-lokasi-indonesia/village"
	"fmt"

	"github.com/gin-gonic/gin"
)

var cityhandler = city.GetCityHandler()
var districthandler = district.GetDistrictHandler()
var villagehandler = village.GetVillageHandler()
var provincehandler = province.GetProvinceHandler()
var allhandler map[string]func(value *string, ginctx *gin.Context) = make(map[string]func(value *string, ginctx *gin.Context))

func fillHandlerMap() {
	allhandler[fmt.Sprintf("%s%s%s", konstant.Province, konstant.ByID, konstant.Province)] = province.GetProvinceByIDHandler
	allhandler[fmt.Sprintf("%s%s%s", konstant.Province, konstant.ByName, konstant.Province)] = province.GetProvinceByNameHandler

	allhandler[fmt.Sprintf("%s%s%s", konstant.City, konstant.ByID, konstant.Province)] = cityhandler.ByID.FromProvince
	allhandler[fmt.Sprintf("%s%s%s", konstant.City, konstant.ByName, konstant.Province)] = cityhandler.ByName.FromProvince

	allhandler[fmt.Sprintf("%s%s%s", konstant.District, konstant.ByID, konstant.City)] = districthandler.ByID.FromCity
	allhandler[fmt.Sprintf("%s%s%s", konstant.District, konstant.ByID, konstant.Province)] = districthandler.ByID.FromProvince
	allhandler[fmt.Sprintf("%s%s%s", konstant.District, konstant.ByName, konstant.City)] = districthandler.ByName.FromCity
	allhandler[fmt.Sprintf("%s%s%s", konstant.District, konstant.ByName, konstant.Province)] = districthandler.ByName.FromProvince

	allhandler[fmt.Sprintf("%s%s%s", konstant.Village, konstant.ByID, konstant.District)] = villagehandler.ByID.FromDistrict
	allhandler[fmt.Sprintf("%s%s%s", konstant.Village, konstant.ByID, konstant.City)] = villagehandler.ByID.FromCity
	allhandler[fmt.Sprintf("%s%s%s", konstant.Village, konstant.ByID, konstant.Province)] = villagehandler.ByID.FromProvince
	allhandler[fmt.Sprintf("%s%s%s", konstant.Village, konstant.ByName, konstant.District)] = villagehandler.ByName.FromDistrict
	allhandler[fmt.Sprintf("%s%s%s", konstant.Village, konstant.ByName, konstant.City)] = villagehandler.ByName.FromCity
	allhandler[fmt.Sprintf("%s%s%s", konstant.Village, konstant.ByName, konstant.Province)] = villagehandler.ByName.FromProvince
}

func init() {
	fillHandlerMap()
}

func Forward() gin.HandlerFunc {
	return func(ginctx *gin.Context) {
		fullpath, value := extractParams(ginctx)
		executeHandler(fullpath, &value, ginctx)
	}
}

func executeHandler(fullpath string, value *string, ginctx *gin.Context) {
	allhandler[fullpath](value, ginctx)
}

func extractParams(ginctx *gin.Context) (fullpath string, Value string) {
	ObjectOne := ginctx.Params.ByName(konstant.ObjectOne)
	ByWhat := ginctx.Params.ByName(konstant.ByWhat)
	ObjectTwo := ginctx.Params.ByName(konstant.ObjectTwo)
	Value = ginctx.Params.ByName(konstant.Value)
	return fmt.Sprintf("%s%s%s", ObjectOne, ByWhat, ObjectTwo), Value
}
