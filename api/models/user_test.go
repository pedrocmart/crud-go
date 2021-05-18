package models

import (
	"testing"
)

func TestUserModel_Validate(t *testing.T) {
	p := &User{}
	p.Name = "John"
	p.Address = "123, Main Road"
	p.Description = "First User"
	p.DOB = "2001-12-01"

	if err := p.Validate(); err != nil {
		t.Error(err)
	}

	//validate empty DOB
	p.DOB = ""
	err := p.Validate()
	expected := "user.DOB cannot be empty"
	if err.Error() != expected {
		t.Errorf("expected: %s, got: %s", expected, err.Error())
	}

	//validate wrong date format for DOB
	p.DOB = "2020-13-32"
	err = p.Validate()
	expected = "user.DOB must be YYYY-MM-DD"
	if err.Error() != expected {
		t.Errorf("expected: %s, got: %s", expected, err.Error())
	}

	//validate DOB greater than today
	p.DOB = "2040-10-01"
	err = p.Validate()
	expected = "user.DOB cannot be greater than today"
	if err.Error() != expected {
		t.Errorf("expected: %s, got: %s", expected, err.Error())
	}
}
