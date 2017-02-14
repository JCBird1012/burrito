package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// BaseURL does things
const BaseURL string = "https://order.chipotle.com"

type loginInformation struct {
	Username string
	Password string
	Persist  bool
}

// Location is a location
type Location struct {
	ID                 int     `json:"Id"`
	Name               string  `json:"Name"`
	Address            string  `json:"Address"`
	Address2           string  `json:"Address2"`
	City               string  `json:"City"`
	State              string  `json:"State"`
	Country            string  `json:"Country"`
	Zip                string  `json:"Zip"`
	Phone              string  `json:"Phone"`
	Live               bool    `json:"Live"`
	OnlineOrderingLive bool    `json:"OnlineOrderingLive"`
	ComingSoon         bool    `json:"ComingSoon"`
	Latitude           float64 `json:"Latitude"`
	Longitude          float64 `json:"Longitude"`
	Distance           string  `json:"Distance"`
	SpecialMessage     string  `json:"SpecialMessage"`
	BusinessHourText   string  `json:"BusinessHourText"`
}

// Login , does things
func Login(username string, password string) (UserToken string) {
	login := &loginInformation{
		Username: username,
		Password: password,
		Persist:  false}

	loginJSON, _ := json.Marshal(login)

	b := bytes.NewBufferString(string(loginJSON))

	req, _ := http.NewRequest("POST", BaseURL+"/api/customer/login", b)

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		panic("We've encountered an error!")
	}

	body, _ := ioutil.ReadAll(res.Body)

	var response (map[string]interface{})

	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	return response["CustomerToken"].(string)

}

// GetLocations does things
func GetLocations(postalCode string) (response []Location) {
	b := bytes.NewBufferString(`{"Address": ` + postalCode + `, "Radius":50, "StartIndex":1, "ReturnCount": 5}`)

	req, _ := http.NewRequest("POST", BaseURL+"/api/restaurant/restaurantssearch", b)

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		panic("We've encountered an error!")
	}

	body, _ := ioutil.ReadAll(res.Body)

	locations := make([]Location, 0)
	json.Unmarshal(body, &locations)

	return locations
}

// SendOrderToSChipotleByStoreID does things too
func SendOrderToSChipotleByStoreID(storeID int, loginToken string, order string) bool {

	return false
}
