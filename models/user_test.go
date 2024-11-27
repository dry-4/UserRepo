package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		expected string
	}{
		{"Valid User", User{Name: "Muskan Shrivastava", Email: "muskan@gmail.com", Age: 20}, ""},
		{"Empty Name", User{Name: "", Email: "muskan@gmail.com", Age: 25}, "Invalid user name.."},
		{"Empty Email", User{Name: "Muskan Shrivastava", Email: "", Age: 26}, "Invalid user email.."},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.user.Validate()
			if testCase.expected == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, testCase.expected)
			}
		})
	}
}
