package village

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type villagehandler struct {
	ByName *byName
	ByID   *byID
}

func GetVillageHandler() *villagehandler {
	return &villagehandler{
		ByName: &byName{},
		ByID:   &byID{},
	}
}

//Endpoint for get village detail by id
func GetVillageByIDHandler(value *string, ginctx *gin.Context) {

	villageid := *value

	villageFound, err := data.UnmarshallVillage(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.Village).ID == villageid {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(villageFound)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}

//Endpoint for get village detail by id
func GetVillageByNameHandler(value *string, ginctx *gin.Context) {

	villagename := *value

	villageFound, err := data.UnmarshallVillage(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.Village).Name == villagename {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(villageFound)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}
