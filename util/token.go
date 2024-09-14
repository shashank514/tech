package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func GenereateToken(auth string) string {
	userExpertDate := time.Now().AddDate(0, 0, 10)
	tokenExpertDate := userExpertDate.AddDate(0, 0, 1)
	body := `{"auth":"` + auth + `","userExpert":"` + cast.ToString(userExpertDate) + `","tokenExpert":"` + cast.ToString(tokenExpertDate) + `"}`
	return base64.StdEncoding.EncodeToString([]byte(body))
}

func GenereateAuthToken(email string) string {
	body := `{"email":"` + email + `"}`
	return base64.StdEncoding.EncodeToString([]byte(body))
}

func DecodeToken(token string) (details map[string]interface{}, err error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(decodedBytes, &details)
	if err != nil {
		fmt.Println("Error in unmarshalling the body ", err)
		return nil, err
	}
	return details, nil
}
