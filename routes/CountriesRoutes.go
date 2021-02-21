package routes

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"task/logic"
)

//Countries find all Countries
func Countries(c echo.Context) error {

	res, err := logic.GetAllCountries()

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//DeleteCountry ticket by it's name
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

	res, err := logic.AddNewCountry(name, location, wiki)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}
type Info struct {
	Name string
	Location string
	Wiki string
}
//UpdateCountry add new Country
func UpdateCountry(c echo.Context) error {

	name := c.Param("countryName")

	nameNew := c.QueryParam("name")
	location := c.QueryParam("location")
	wiki := c.QueryParam("wiki")

	info,_:=logic.GetCountryByName(name)
	if(nameNew==" ") {
		for _, con := range info {
			res, err := logic.UpdateCountry(name, con.Name, location, wiki)
			if err != nil {
				log.Println("Unsupported method")
			}
			return c.JSON(http.StatusCreated, res)
		}
	}else if(location==" "){
		for _, con := range info {
			res, err := logic.UpdateCountry(name, nameNew, con.Location, wiki)
			if err != nil {
				log.Println("Unsupported method")
			}
			return c.JSON(http.StatusCreated, res)
		}

	}else if(wiki==" "){
		for _, con := range info {
			res, err := logic.UpdateCountry(name, nameNew, location, con.Wiki)
			if err != nil {
				log.Println("Unsupported method")
			}
			return c.JSON(http.StatusCreated, res)
		}

	}else{
	res, err := logic.UpdateCountry(name, nameNew, location, wiki)
		if err != nil {
			log.Println("Unsupported method")
		}
		return c.JSON(http.StatusCreated, res)
	}
	return nil
}