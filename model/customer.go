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

func (c *Customer) GetGender() Gender {
	return c.Gender
}

func (e Gender) String() string {
	genders := [...]string{"Unknown", "Male", "Female"}

	x := string(e)
	for _, v := range genders {
		if v == x {
			return x
		}
	}

	return ""
}

func (c *Customer) MarshalGenderJson() ([]byte, error) {
	type Alias Customer
	return json.Marshal(&struct {
		*Alias
		Gender Gender `json:"gender"`
	}{
		Alias:  (*Alias)(c),
		Gender: Gender(c.GetGender().String()),
	})
}

func (c *Customer) MarshalDateJSON() ([]byte, error) {
	type Alias Customer
	return json.Marshal(&struct {
		*Alias
		Birthdate string `json:"birthdate"`
	}{
		Alias:     (*Alias)(c),
		Birthdate: c.Birthdate.Format("2006-01-02"),
	})
}
