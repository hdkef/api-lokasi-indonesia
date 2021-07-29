package data

import (
	"api-lokasi-indonesia/models"
	"errors"
	"io"

	"github.com/jszwec/csvutil"
)

func doLoopingProvince(dec *csvutil.Decoder, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {
	slicesProvince := []models.Province{}

	for {
		var unit models.Province

		err := dec.Decode(&unit)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if item, valid, isItFinish := filter(unit); valid {
			slicesProvince = append(slicesProvince, item.(models.Province))
			if isItFinish {
				break
			}
		}
	}

	if len(slicesProvince) == 0 {
		return nil, errors.New("NO DATA")
	}

	return slicesProvince, nil
}

func doLoopingCity(dec *csvutil.Decoder, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {
	slicesCity := []models.City{}

	for {
		var unit models.City

		err := dec.Decode(&unit)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if item, valid, isItFinish := filter(unit); valid {
			slicesCity = append(slicesCity, item.(models.City))
			if isItFinish {
				break
			}
		}
	}

	if len(slicesCity) == 0 {
		return nil, errors.New("NO DATA")
	}

	return slicesCity, nil
}

func doLoopingDistrict(dec *csvutil.Decoder, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {
	slicesDistrict := []models.District{}

	for {
		var unit models.District

		err := dec.Decode(&unit)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if item, valid, isItFinish := filter(unit); valid {
			slicesDistrict = append(slicesDistrict, item.(models.District))
			if isItFinish {
				break
			}
		}
	}

	if len(slicesDistrict) == 0 {
		return nil, errors.New("NO DATA")
	}

	return slicesDistrict, nil
}

func doLoopingVillage(dec *csvutil.Decoder, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {
	slicesVillage := []models.Village{}

	for {
		var unit models.Village

		err := dec.Decode(&unit)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if item, valid, isItFinish := filter(unit); valid {
			slicesVillage = append(slicesVillage, item.(models.Village))
			if isItFinish {
				break
			}
		}
	}

	if len(slicesVillage) == 0 {
		return nil, errors.New("NO DATA")
	}

	return slicesVillage, nil
}
