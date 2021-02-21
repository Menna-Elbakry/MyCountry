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

func UpdateCountryName(Name string, newName string) (string, error) {
	newCon := &model.Country{
		Name: newName,
	}
	return database.UpdateCountryName(Name, newCon)
}
func UpdateCountryLocation(Name string, newLocation string) (string, error) {
	newCon := &model.Country{
		Location: newLocation,
	}
	return database.UpdateCountryLocation(Name, newCon)
}
func UpdateCountryWiki(Name string, newWiki string) (string, error) {
	newCon := &model.Country{
		Wiki: newWiki,
	}
	return database.UpdateCountryWiki(Name, newCon)
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
