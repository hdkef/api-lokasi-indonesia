package district

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
