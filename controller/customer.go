package controller

import (
	"customers/model"
	"customers/service"
	"customers/storage"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCustomers(c echo.Context) error {

	firstname := c.QueryParam("firstname")
	lastname := c.QueryParam("lastname")

	firstnameToCompare := service.TrimAndUpperCaseString(firstname)
	lastnameToCompare := service.TrimAndUpperCaseString(lastname)

	if len(firstnameToCompare) != 0 {
		query := "upper(firstname) LIKE ?"
		return c.JSON(http.StatusOK, service.GetFilteredCustomers(query, firstnameToCompare))
	}
	if len(lastnameToCompare) != 0 {
		query := "upper(lastname) LIKE ?"
		return c.JSON(http.StatusOK, service.GetFilteredCustomers(query, lastnameToCompare))
	}
	return c.JSON(http.StatusOK, service.GetAllCustomers())

}

func CreateCustomer(c echo.Context) error {
	firstname, lastname, birthdate, gender, email, address := service.RetrieveQueryParameters(c)
	dateTimeType, _ := service.ParseTimeString(birthdate)

	err, done := service.RequiredFieldCheck(c, firstname, lastname, birthdate, gender, email)
	if done {
		return err
	}

	err2, done2 := service.IsDataValid(c, firstname, lastname, birthdate, gender, email, address, dateTimeType)
	if done2 {
		return err2
	}
	DB := storage.GetDBInstance()

	DB.Create(&model.Customer{
		Firstname: firstname,
		Lastname:  lastname,
		Birthdate: dateTimeType,
		Gender:    model.Gender(gender),
		Email:     email,
		Address:   address})

	//return nil
	var message = "The customer is added"
	return c.JSON(http.StatusMethodNotAllowed, message)
}

func UpdateCustomer(c echo.Context) error {
	id := c.QueryParam("id")
	firstname, lastname, birthdate, gender, email, address := service.RetrieveQueryParameters(c)
	dateTimeType, _ := service.ParseTimeString(birthdate)

	err, done := service.IsDataValid(c, firstname, lastname, birthdate, gender, email, address, dateTimeType)
	if done {
		return err
	}

	var customers []model.Customer

	DB := storage.GetDBInstance()
	DB.Model(customers).Where("id = ?", id).Updates(model.Customer{
		Firstname: firstname,
		Lastname:  lastname,
		Birthdate: dateTimeType,
		Gender:    model.Gender(gender),
		Email:     email,
		Address:   address})

	var message = "The customer is updated"
	return c.JSON(http.StatusMethodNotAllowed, message)
}
