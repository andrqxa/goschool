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
	Registration time.Time
}

func (u *User) String() string {
	return fmt.Sprintf("[user=%s; email=%s; registration=%s]", u.Name, u.Email, u.Registration.UTC().Format("02.01.2006"))
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
		time.Now(),
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
	regTempl := regexp.MustCompile(`[a-zA-Z][a-zA-Z\.]*`)
	names := strings.Split(strings.TrimSpace(name), " ")
	if len(names) < 1 {
		return fmt.Errorf("name can't be empty")
	}
	for _, nm := range names {
		if !regTempl.MatchString(nm) {
			return fmt.Errorf("name must be composed of English letters, dots and spaces")
		}
	}
	return nil
}

func validateEmail(email string) error {
	regTempl := regexp.MustCompile(`[a-zA-Z0-9\._\-]+`)
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
