package routes

import (
	"log"
	"net/http"
	"task/logic"

	"github.com/labstack/echo/v4"
)

//Countries to display all Countries
func Countries(c echo.Context) error {

	res, err := logic.GetAllCountries()

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//DeleteCountry to remove country by it's name
func DeleteCountry(c echo.Context) error {

	name := c.Param("countryName")

	res, err := logic.DeleteCountry(name)
	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//AddCountry add new Country
func AddCountry(c echo.Context) error {

	name := c.QueryParam("name")
	location := c.QueryParam("location")
	wiki := c.QueryParam("wiki")
	if name == "" || location == "" || wiki == "" {
		log.Print("Please Enter Data Correctly")
		return nil
	}
	res, err := logic.AddNewCountry(name, location, wiki)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//UpdateCountryName to modify name
func UpdateCountryName(c echo.Context) error {

	name := c.Param("countryName")

	nameNew := c.QueryParam("name")
	if nameNew == "" {
		log.Print("Please Enter Name")
		return nil
	}
	res, err := logic.UpdateCountryName(name, nameNew)
	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)

}

//UpdateCountryLocation to modify location
func UpdateCountryLocation(c echo.Context) error {

	name := c.Param("countryName")

	location := c.QueryParam("location")

	res, err := logic.UpdateCountryLocation(name, location)
	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//UpdateCountryWiki to modify wiki
func UpdateCountryWiki(c echo.Context) error {
	name := c.Param("countryName")

	wiki := c.QueryParam("wiki")
	res, err := logic.UpdateCountryWiki(name, wiki)
	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)

}
