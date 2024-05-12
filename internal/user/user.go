package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetRandomUser() User {

	var user User

	resp, _ := http.Get("https://random-data-api.com/api/v2/users")

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &user)
	return user
}
