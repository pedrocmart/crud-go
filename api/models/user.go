package models

import (
	"errors"
	"regexp"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Model
	Name        string `gorm:"size:200;not null" json:"name"`
	DOB         string `gorm:"size:10 not null" json:"dob"`
	Address     string `gorm:"size:300" json:"address"`
	Description string `gorm:"size:300" json:"description"`
}

var ErrUserEmptyName = errors.New("user.name cannot be empty")
var ErrUserNameMaxLen = errors.New("user.name max length is 200")
var ErrUserEmptyDOB = errors.New("user.DOB cannot be empty")
var ErrUserWrongFormatDOB = errors.New("user.DOB must be YYYY-MM-DD")
var ErrUserGreaterDOB = errors.New("user.DOB cannot be greater than today")
var ErrUserAddressMaxLen = errors.New("user.address max length is 300")
var ErrUserDescriptionMaxLen = errors.New("user.description max length is 300")

func (p *User) Validate() error {
	if p.Name == "" {
		return ErrUserEmptyName
	}

	//validate max length
	if len(p.Name) > 300 {
		return ErrUserNameMaxLen
	}

	if len(p.Address) > 300 {
		return ErrUserAddressMaxLen
	}

	if len(p.Description) > 300 {
		return ErrUserDescriptionMaxLen
	}

	if p.DOB == "" {
		return ErrUserEmptyDOB
	}

	//validate DOB format
	format := regexp.MustCompile("((19|20)\\d\\d)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$")
	if !format.MatchString(p.DOB) {
		return ErrUserWrongFormatDOB
	}

	//validate if DOB is greater than today
	timeDOB, _ := time.Parse("2006-01-02", p.DOB)
	today := time.Now()
	if today.Before(timeDOB) {
		return ErrUserGreaterDOB
	}

	//validate length

	return nil
}
