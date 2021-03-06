package routes

import (
	"github.com/labstack/echo/v4"
)

//Server to connect to http
func Server() {

	e := echo.New()

	e.GET("/countries", Countries)

	e.POST("/addCountry", AddCountry)

	e.DELETE("/deleteCountry/:countryName", DeleteCountry)

	e.PUT("/updateCountryName/:countryName", UpdateCountryName)
	e.PUT("/updateCountryLocation/:countryName", UpdateCountryLocation)
	e.PUT("/updateCountryWiki/:countryName", UpdateCountryWiki)

	e.Logger.Fatal(e.Start(":8081"))

}
