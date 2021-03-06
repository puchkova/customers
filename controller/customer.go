package controller

import (
	"customers/model"
	"customers/service"
	"customers/storage"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetCustomers(c echo.Context) error {

	firstname := c.QueryParam("firstname")
	lastname := c.QueryParam("lastname")

	firstnameToCompare := service.GetUppercaseAndTrimmedString(firstname)
	lastnameToCompare := service.GetUppercaseAndTrimmedString(lastname)

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
	firstname, lastname, birthdate, gender, email, address := service.GetQueryParameters(c)
	dateTimeType, _ := service.GetTimeString(birthdate)

	err, done := GetRequiredFieldErrorMessages(c, firstname, lastname, birthdate, gender, email)
	if done {
		return err
	}

	err2, done2 := GetInvalidInputErrorMessages(c, firstname, lastname, birthdate, gender, email, address, dateTimeType)
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

	var message = "The customer is added"
	return c.JSON(http.StatusOK, message)
}

func UpdateCustomer(c echo.Context) error {
	id := c.QueryParam("id")
	firstname, lastname, birthdate, gender, email, address := service.GetQueryParameters(c)
	dateTimeType, _ := service.GetTimeString(birthdate)

	err, done := GetInvalidInputErrorMessages(c, firstname, lastname, birthdate, gender, email, address, dateTimeType)
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
	return c.JSON(http.StatusOK, message)
}

func GetInvalidInputErrorMessages(c echo.Context, firstname string, lastname string, birthdate string,
	gender string, email string, address string, dateTimeType time.Time) (error, bool) {
	if !service.IsBirthdateValid(dateTimeType, 18, 60) && len(birthdate) != 0 {
		var errorMessage = "Age should be in the range from 18 to 60 years"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsInputStringValid(firstname, 1, 100) {
		var errorMessage = "Invalid First Name"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsInputStringValid(lastname, 1, 100) {
		var errorMessage = "Invalid Last Name"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsGenderValid(gender) {
		var errorMessage = "Gender should be Male or Female"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsEmailValid(email) && len(email) != 0 {
		var errorMessage = "Invalid email address format"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsInputStringValid(address, 2, 200) {
		var errorMessage = "Invalid address"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	return nil, false
}

func GetRequiredFieldErrorMessages(c echo.Context, firstname string, lastname string,
	birthdate string, gender string, email string) (error, bool) {
	if !service.IsFieldRequired(firstname) {
		var errorMessage = "First Name is required field"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsFieldRequired(lastname) {
		var errorMessage = "Last Name is required field"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsFieldRequired(birthdate) {
		var errorMessage = "Birthdate is required field"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsFieldRequired(gender) {
		var errorMessage = "Gender is required field"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	if !service.IsFieldRequired(email) {
		var errorMessage = "Email is required field"
		return c.JSON(http.StatusMethodNotAllowed, errorMessage), true
	}
	return nil, false
}
