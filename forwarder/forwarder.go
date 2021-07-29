package forwarder

import (
	"api-lokasi-indonesia/city"
	"api-lokasi-indonesia/district"
	"api-lokasi-indonesia/konstant"
	"api-lokasi-indonesia/village"
	"fmt"

	"github.com/gin-gonic/gin"
)

var cityhandler = city.GetCityHandler()
var districthandler = district.GetDistrictHandler()
var villagehandler = village.GetVillageHandler()
var allhandler map[string]func(value *string, ginctx *gin.Context) = make(map[string]func(value *string, ginctx *gin.Context))

func fillHandlerMap() {
	allhandler[fmt.Sprintf("%s%s%s", konstant.City, konstant.ByID, konstant.Province)] = cityhandler.ByID.FromProvince
	allhandler[fmt.Sprintf("%s%s%s", konstant.City, konstant.ByName, konstant.Province)] = cityhandler.ByName.FromProvince
	allhandler[fmt.Sprintf("%s%s%s", konstant.District, konstant.ByID, konstant.City)] = districthandler.ByID.FromCity
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
