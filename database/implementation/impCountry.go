package database

import (
	"log"
	"task/database/table"
	"task/model"

	"github.com/go-pg/pg/orm"
)

//GetAllCountries to select all countries exist
func GetAllCountries() ([]model.Country, error) {
	var conEntity []table.Country
	findErr := db.Model(&conEntity).Select()
	if findErr != nil {
		log.Printf("Error While Getting Countries Reason:  %v\n", findErr)
		return nil, findErr
	}
	var countries []model.Country
	for _, con := range conEntity {
		countries = append(countries, con.MapToModule())
		log.Printf("Get Countries Successful . \n Name: %v ,Location:%v \n Wiki:%v", con.Name, con.Location,con.Wiki)
	}
	return countries, nil
}
//GetCountryByName to get country by name
func GetCountryByName(name string) ([]model.Country, error) {
	var conEntity []table.Country
	findErr := db.Model(&conEntity).
		Where("name=?", name).
		Select()
	if findErr != nil {
		log.Printf("Error While Getting Country Reason:  %v\n", findErr)
		return nil, findErr
	}
	var country []model.Country
	for _, con := range conEntity {
		country = append(country, con.MapToModule())
		log.Printf("Get Country Successful . \n Name: %v ,Location:%v \n Wiki:%v", con.Name, con.Location,con.Wiki)
	}
	return country, nil
}

//AddNewCountry to database
func AddNewCountry(con *model.Country) (string, error) {
	_, addErr := db.Model(table.Country{}.Fill(con)).Insert()
	if addErr != nil {
		log.Printf("Error While Adding New Country Reason:  %v\n", addErr)
		return " ", addErr
	}
	log.Printf("New Country Added Successfully .\n Country Name : %s", con.Name)
	return con.Name, nil
}

//DeleteCountry to delete country
func DeleteCountry(con *model.Country) (string, error) {
	var conEntity table.Country
	_, delErr := db.Model(&conEntity).Where("name=?", con.Name).
		Returning("name").
		Delete()
	if delErr != nil {
		log.Printf("Error While Deleting Country. Reason: %v \n", delErr)
		return " ", delErr
	}
	log.Printf("Successfully Deleted Country : %v \n", conEntity.Name)
	return conEntity.Name, nil
}
//UpdateCountry to delete country
func UpdateCountry(Name string ,con *model.Country) (string, error) {
	var conEntity table.Country
	_, updErr := db.Model(&conEntity).Where("name=?",Name).
		Set("name=?,location=?,wiki=?",con.Name,con.Location,con.Wiki).
		Returning("name").
		Update()
	if updErr != nil {
		log.Printf("Error While Updating Country. Reason: %v \n", updErr)
		return " ", updErr
	}
	log.Printf("Successfully Updating Country : %v \n", conEntity.Name)
	return conEntity.Name, nil
}

//CreateCountryTable to create table country in database
func CreateCountryTable() error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&table.Country{}, opts)
	if createErr != nil {
		log.Printf("Error While Creating Table Country Reason: \n %v", createErr)
		return createErr
	}
	log.Printf("Table Country Created Successfully.\n")
	return nil
}
