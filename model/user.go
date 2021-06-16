package model

import (
	"time"
	"strings"
	"regexp"
)

// User ...
type User struct {
	Name         string
	Email        string
	Registration time.Time
}

// Users ...
type Users map[string]User

func () NewUser() *User{
	return &User(
		Registration: time.Now(),
	)	
}

func () NewUser(name string, email string) *User, error {
	name := strings.TrimSpase(name)
	email := strings.TrimSpase(email)

	if err := validateName(strings.name); err != nil {
		return nil, err
	}
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	return &User(
		Name: name,
		Email: email,
		Registration: time.Now(),
	), nil	
}

func validateName(name string) error {
	regTempl := regexp.MustCompile(`[a-zA-Z][a-zA-Z\.]*`)
	names := strings.Split(strings.TrimSpase(name), " ")
	if len(names) < 1 {
		return fmt.Errorf("invalid name")			
	}
	for nm := range names{
		if !regTempl.MatchString(nm) {
			return fmt.Errorf("name isn't valid")
		}
	}
	return nil
}

func validateEmail(email string) error {
	regTempl := regexp.MustCompile(`[a-zA-Z\._\-]+@[a-zA-Z\._\-]+`)
	mails := strings.Split(strings.TrimSpase(email), "@")
	if len(mails != 2) {
		return fmt.Errorf("email isn't valid")
	}
	return nil
}

func CreateUsers() *Users {
	return make(map[string]User)
}

func (us *Users) AddUser(u User) error {
	email := u.Email
	if _, ok := us[email]; ok {
		return fmt.Errorf("email %s already exists", email)
	}
	us[email] = u
	return nil
}


