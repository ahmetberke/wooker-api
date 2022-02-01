package utils

import (
	"testing"
)

func TestGenerateUsernameFromEmail(t *testing.T) {
	username := GenerateUsernameFromEmail("ahmet@gmail.com", "@gmail.com")
	if username != "ahmet" {
		t.Errorf("username is %v, must be ahmet", username)
	}
	username = GenerateUsernameFromEmail("ahmet@gmail.com")
	if username != "ahmet" {
		t.Errorf("username is %v, must be ahmet", username)
	}
}