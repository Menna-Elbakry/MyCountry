package logic

import (
	database "task/database/implementation"
	"task/model"
)

func AddNewCountry(ConName string, Location string, Wiki string) (string, error) {
	newCon := &model.Country{
		Name:     ConName,
		Location: Location,
		Wiki:     Wiki,
	}
	return database.AddNewCountry(newCon)

}

func UpdateCountry(Name string,ConName string, Location string, Wiki string) (string, error) {
	newCon := &model.Country{
		Name:     ConName,
		Location: Location,
		Wiki:     Wiki,
	}
	return database.UpdateCountry(Name ,newCon)

}


func GetAllCountries() ([]model.Country, error) {
	return database.GetAllCountries()
}
func GetCountryByName(name string) ([]model.Country, error) {
	return database.GetCountryByName(name)
}

func DeleteCountry(ConName string) (string, error) {
	newCon := &model.Country{
		Name: ConName,
	}
	return database.DeleteCountry(newCon)
}
