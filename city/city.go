package city

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cityhandler struct {
	ByName *byName
	ByID   *byID
}

func GetCityHandler() *cityhandler {
	return &cityhandler{
		ByName: &byName{},
		ByID:   &byID{},
	}
}

//Endpoint for get city detail by id
func GetCityByIDHandler(value *string, ginctx *gin.Context) {

	cityid := *value

	cityFound, err := data.UnmarshallCity(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.City).ID == cityid {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(cityFound)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}

//Endpoint for get city detail by name
func GetCityByNameHandler(value *string, ginctx *gin.Context) {

	cityname := *value

	cityFound, err := data.UnmarshallCity(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.City).Name == cityname {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(cityFound)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}
