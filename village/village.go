package village

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
