package district

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/province"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type byName struct {
}

//WARNING!!! value is pointer MUST BE COMPARE WITH DEPOINTER FIRST
func (b *byName) FromCity(value *string, ginctx *gin.Context) {

	cityname := *value

	//get one city from cityname
	city, err := data.UnmarshallCity(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.City).Name == cityname {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})

	//get all districts from cityFound.ID
	districts, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.District).CityID == city[0].ID { //First item of slices
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
func (b *byName) FromProvince(value *string, ginctx *gin.Context) {

	var alldistricts []models.District

	//find province from province name
	provinceFound := province.GetProvinceByName(value)

	//get all cities with provinceid
	cities, err := data.UnmarshallCity(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.City).ProvinceID == provinceFound.ID {
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
