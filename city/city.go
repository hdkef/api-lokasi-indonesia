package city

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
