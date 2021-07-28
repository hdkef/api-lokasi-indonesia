package konstant

import "fmt"

type konstant struct {
}

func GetKonstantObject() *konstant {
	return &konstant{}
}

var Get = "get"
var ByID = "byid"
var ByName = "byname"
var City = "city"
var Province = "province"
var ObjectOne = "objectone"
var ObjectTwo = "objecttwo"
var Value = "value"
var ByWhat = "bywhat"

func (k *konstant) GetPath() string {
	return fmt.Sprintf("/%s/:%s/:%s/:%s/:%s", Get, ObjectOne, ByWhat, ObjectTwo, Value)
}

func (k *konstant) GetAllProvincePath() string {
	return "/provinces"
}
