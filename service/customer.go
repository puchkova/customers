package service

import (
	"customers/model"
	"customers/storage"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"regexp"
	"strings"
	"time"
)

var customers []model.Customer

func RetrieveQueryParameters(c echo.Context) (string, string, string, string, string, string) {
	firstname := c.QueryParam("firstname")
	lastname := c.QueryParam("lastname")
	birthdate := c.QueryParam("birthdate")
	gender := c.QueryParam("gender")
	email := c.QueryParam("email")
	address := c.QueryParam("address")
	return firstname, lastname, birthdate, gender, email, address
}

func GetAllCustomers() interface{} {
	DB := storage.GetDBInstance()
	return DB.Find(&customers)
}

func GetFilteredCustomers(query string, name string) *gorm.DB {
	DB := storage.GetDBInstance()
	return DB.Where(query, "%"+name+"%").Find(&customers)
}

func ParseTimeString(birthdate string) (time.Time, error) {
	return time.Parse("2006-01-02", birthdate)
}

func TrimAndUpperCaseString(name string) string {
	return strings.ToUpper(strings.TrimSpace(name))
}

func IsGenderValid(gender string) bool {
	switch gender {
	case model.Male.String(), model.Female, "":
		return true
	}
	return false
}

func FieldIsRequired(field string) bool {
	if len(strings.TrimSpace(field)) != 0 {
		return true
	}
	return false
}

func IsValid(str string, min int, max int) bool {
	if len(strings.TrimSpace(str)) == 0 || len(str) > min && len(str) < max {
		return true
	}
	return false
}

func IsEmailValid(email string) bool {
	var regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9]" +
		"(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return regexpEmail.MatchString(email)
}

func IsBirthdateValid(t time.Time, min int, max int) bool {
	if CalculateAge(t, time.Now()) < min || CalculateAge(t, time.Now()) > max {
		return false
	}
	return true
}

func CalculateAge(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}
