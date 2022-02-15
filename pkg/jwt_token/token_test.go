package jwt_token

import "testing"

func TestGenerateAndParseToken(t *testing.T) {

	testSecretKey := "123"
	var userID uint = 1
	token, err := GenerateToken(userID, testSecretKey)

	if err != nil {
		t.Errorf("error on Generate Token Func, %v", err.Error())
	}

	tUserID, err := ParseToken(token, testSecretKey)
	if err != nil {
		t.Errorf("error on Parse Token Func, %v", err.Error())
	}

	if tUserID != userID {
		t.Errorf("user id is %v, must be %v", tUserID, userID)
	}

}