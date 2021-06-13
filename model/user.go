package model

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// User ...
type User struct {
	Name         string
	Email        string
	Registration string
}

func (u *User) String() string {
	return fmt.Sprintf("[user=%s; email=%s; registration=%s]", u.Name, u.Email, u.Registration)
}

// Users ...
type Users map[string]*User

func (u *Users) String() string {
	output := ""
	for _, us := range *u {
		output += fmt.Sprintf("%s\n", us.String())
	}
	return output
}

func NewUser(name string, email string) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	return &User{
		name,
		email,
		time.Now().UTC().Format("02.01.2006"),
	}, nil
}

func (u *Users) ToArray() []*User {
	result := make([]*User, 0)
	for _, currUser := range *u {
		result = append(result, currUser)
	}
	return result
}

func validateName(name string) error {
	regTempl := regexp.MustCompile(`^[a-zA-z][a-zA-Z]*\.*\s*[a-zA-z][a-zA-Z]*\.*$`)
	names := strings.TrimSpace(name)
	if len(names) < 1 {
		return fmt.Errorf("name can't be empty")
	}
	if !regTempl.MatchString(names) {
		return fmt.Errorf("name must be composed of English letters, dots and spaces")
	}
	return nil
}

func validateEmail(email string) error {
	regTempl := regexp.MustCompile(`^[a-zA-Z][\w\.\-_]*$`)
	mails := strings.Split(strings.TrimSpace(email), "@")
	if len(mails) != 2 {
		return fmt.Errorf("email isn't valid")
	}
	addr, host := mails[0], mails[1]
	if !regTempl.MatchString(addr) || !regTempl.MatchString(host) {
		return fmt.Errorf("email isn't valid")
	}
	return nil
}

func NewUsers() Users {
	u := make(map[string]*User)
	return u
}

func (us Users) AddUser(u *User) error {
	email := u.Email
	if _, ok := us[email]; ok {
		return fmt.Errorf("email %s already exists", email)
	}
	us[email] = u
	return nil
}
