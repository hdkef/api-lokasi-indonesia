package district

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type districthandler struct {
	ByName *byName
	ByID   *byID
}

func GetDistrictHandler() *districthandler {
	return &districthandler{
		ByName: &byName{},
		ByID:   &byID{},
	}
}

//Endpoint for get district detail by name
func GetDistrictByNameHandler(value *string, ginctx *gin.Context) {

	districtname := *value

	districtFound, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.District).Name == districtname {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(districtFound)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}

//Endpoint for get district detail by id
func GetDistrictByIDHandler(value *string, ginctx *gin.Context) {

	districtid := *value

	districtFound, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.District).ID == districtid {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(districtFound)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}
