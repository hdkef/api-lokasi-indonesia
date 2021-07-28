package forwarder

import (
	"api-lokasi-indonesia/city"
	"api-lokasi-indonesia/konstant"
	"fmt"

	"github.com/gin-gonic/gin"
)

var cityhandler = city.GetCityHandler()
var handler map[string]func(value *string, ginctx *gin.Context) = make(map[string]func(value *string, ginctx *gin.Context))

func fillHandlerMap() {
	handler[fmt.Sprintf("%s%s%s", konstant.City, konstant.ByID, konstant.Province)] = cityhandler.ByID.FromProvince
	handler[fmt.Sprintf("%s%s%s", konstant.City, konstant.ByName, konstant.Province)] = cityhandler.ByName.FromProvince
}

func init() {
	fillHandlerMap()
}

func Forward() gin.HandlerFunc {
	return func(ginctx *gin.Context) {
		fullpath, value := extractParams(ginctx)
		next := handler[fullpath]
		next(&value, ginctx)
	}
}

func extractParams(ginctx *gin.Context) (fullpath string, Value string) {
	ObjectOne := ginctx.Params.ByName(konstant.ObjectOne)
	ByWhat := ginctx.Params.ByName(konstant.ByWhat)
	ObjectTwo := ginctx.Params.ByName(konstant.ObjectTwo)
	Value = ginctx.Params.ByName(konstant.Value)
	return fmt.Sprintf("%s%s%s", ObjectOne, ByWhat, ObjectTwo), Value
}
