package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_URL string = "https://order.chipotle.com"

type loginInformation struct {
	Username string
	Password string
	Persist  bool
}

func login() {
	login := &loginInformation{
		Username: "jbirdcaicedo@yahoo.com",
		Password: "burritos123",
		Persist:  false}

	loginJSON, _ := json.Marshal(login)

	b := bytes.NewBufferString(string(loginJSON))

	req, _ := http.NewRequest("POST", BASE_URL+"/api/customer/login", b)

	req.Header.Add("json", "true")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		fmt.Println("We've encountered an error!")
	}

	body, _ := ioutil.ReadAll(res.Body)

	var response (map[string]interface{})

	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	USER_COOKIE = response["CustomerToken"].(string)
	return USER_COOKIE
}
