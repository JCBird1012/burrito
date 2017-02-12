package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Global variables
const baseURL string = "https://order.chipotle.com"

type loginInformation struct {
	Username string
	Password string
	Persist  bool
}

// Login does things
func Login() string {
	login := &loginInformation{
		Username: "jbirdcaicedo@yahoo.com",
		Password: "burritos123",
		Persist:  false}

	loginJSON, _ := json.Marshal(login)

	b := bytes.NewBufferString(string(loginJSON))

	req, _ := http.NewRequest("POST", baseURL+"/api/customer/login", b)

	req.Header.Add("json", "true")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		fmt.Println("AHHHH! We've encountered an error!")
	}

	body, _ := ioutil.ReadAll(res.Body)

	var response (map[string]interface{})

	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	userCookie := response["CustomerToken"].(string)
	return userCookie
}
