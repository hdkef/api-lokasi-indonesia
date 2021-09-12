package district

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type byID struct {
}

//WARNING!!! value is pointer MUST BE COMPARE WITH DEPOINTER FIRST
func (b *byID) FromCity(value *string, ginctx *gin.Context) {

	cityid := *value

	//get all districts from cityid
	districts, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.District).CityID == cityid {
			return c.(interface{}), true, false
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(districts)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}

//WARNING!! THIS OPERATION TAKES O^n time complexity
//WARNING!!! value is pointer MUST BE COMPARE WITH DEPOINTER FIRST
func (b *byID) FromProvince(value *string, ginctx *gin.Context) {

	var alldistricts []models.District

	provinceid := *value

	//get all cities with provinceid
	cities, err := data.UnmarshallCity(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.City).ProvinceID == provinceid {
			return c.(interface{}), true, false
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	//looping every city for district
	for _, v := range cities {
		districts, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
			if c.(models.District).CityID == v.ID {
				return c.(interface{}), true, false
			}
			return c.(interface{}), false, false
		})
		if err != nil {
			utils.ResErr(ginctx, http.StatusInternalServerError, err)
			return
		}
		alldistricts = append(alldistricts, districts...)
	}

	respond, err := json.Marshal(alldistricts)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}
