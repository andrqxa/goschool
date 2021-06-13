package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateName(t *testing.T) {
	namesTrue := make([]string, 0)
	namesTrue = append(namesTrue, "Aaaaa")
	namesTrue = append(namesTrue, "Aaa Bbbb")
	namesTrue = append(namesTrue, "Aaaa.Bbb")
	namesTrue = append(namesTrue, "Aaaa. Bbb")
	namesTrue = append(namesTrue, "Aaaa.Bbb.")
	for _, v := range namesTrue {
		assert.Equal(t, validateName(v), nil)
	}

	namesFalse := make([]string, 0)
	namesFalse = append(namesFalse, "Aaaa#BBB")
	namesFalse = append(namesFalse, "Aaa1BBB")
	namesFalse = append(namesFalse, "AAA BBB ..")
	for _, v := range namesFalse {
		assert.NotEqual(t, validateName(v), nil)
	}
}

func TestValidateEmail(t *testing.T) {
	emailTrue := make([]string, 0)
	emailTrue = append(emailTrue, "Aaaaa@gmail.com")
	emailTrue = append(emailTrue, "Aa11aaa@gmail.com")
	emailTrue = append(emailTrue, "Aaa1.aa@gmail.com")
	emailTrue = append(emailTrue, "A-aaa@gmail.com")
	emailTrue = append(emailTrue, "Aaa.a-a@gmail.com")
	for _, v := range emailTrue {
		assert.Equal(t, validateEmail(v), nil)
	}

	emailFalse := make([]string, 0)
	emailFalse = append(emailFalse, "aaa")
	emailFalse = append(emailFalse, "aaa@bbb@ccc")
	emailFalse = append(emailFalse, "aaa!%@gmail.com")
	for _, v := range emailFalse {
		assert.NotEqual(t, validateEmail(v), nil)
	}
}
