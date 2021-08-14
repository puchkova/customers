package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID string `gorm:"primary_key;"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", uuid.NewV4().String())
}

type Customer struct {
	Base
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthdate time.Time `json:"birthdate"`
	Gender    Gender    `json:"gender"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
}

type Gender string

const (
	Male   Gender = "Male"
	Female        = "Female"
)

func (e Gender) String() string {
	genders := [...]string{"Male", "Female"}

	x := string(e)
	for _, v := range genders {
		if v == x {
			return x
		}
	}
	return ""
}

func (c *Customer) MarshalJSON() ([]byte, error) {
	type Alias Customer
	return json.Marshal(&struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Birthdate string `json:"birthdate"`
		Gender    string `json:"gender"`
		Email     string `json:"email"`
		Address   string `json:"address"`
		*Alias
	}{
		Firstname: c.Firstname,
		Lastname:  c.Lastname,
		Birthdate: c.Birthdate.Format("2006-01-02"),
		Gender:    c.Gender.String(),
		Email:     c.Email,
		Address:   c.Address,
		Alias:     (*Alias)(c),
	})
}
