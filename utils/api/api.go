/*
Package api handles interaction with Chipotle's Ordering API
*/

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// BaseURL is the base URL for Chipotle ordering
const BaseURL string = "https://order.chipotle.com"

type loginInformation struct {
	Username string
	Password string
	Persist  bool
}

// The location struct denotes an API response representing a Chipotle location
type Location struct {
	Id                 int     `json:"Id"`
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

// The order struct denotes an API response representing a complete Chipotle order
type Order struct {
	OrderId   int    `json:"OrderId"`
	OrderName string `json:"OrderName"`
	OrderDate string `json:"OrderDate"`
	OrderType string `json:"OrderType"`
	Meals     []struct {
		MealId   int    `json:"MealId"`
		Entree   string `json:"Entree"`
		MealType string `json:"MealType"`
		Items    []struct {
			Type     string `json:"Type"`
			Quantity int    `json:"Quantity"`
			Portion  string `json:"Portion"`
		} `json:"Items"`
		Instructions        string `json:"Instructions"`
		Name                string `json:"Name"`
		HasUnavailableItems bool   `json:"HasUnavailableItems"`
	} `json:"Meals"`
	HasUnavailableItems bool `json:"HasUnavailableItems"`
}

/*
   POSTs the Chipotle API with user credentials and returns a API cookie (used as a token in future requests)

   IMPORTANT NOTE: userToken needs to be treated as a password, as it can be used to preform actions on behalf of the user

*/
func Login(username string, password string) (userToken string, err error) {
	login := &loginInformation{
		Username: username,
		Password: password,
		Persist:  false}

	loginJSON, err := json.Marshal(login)

	if err != nil {
		return "", errors.New("EJSONMARSHAL")
	}

	b := bytes.NewBufferString(string(loginJSON))

	req, _ := http.NewRequest("POST", BaseURL+"/api/customer/login", b)

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		return "", errors.New("EHTTPREQUEST")
	}

	body, _ := ioutil.ReadAll(res.Body)

	var response (map[string]interface{})

	if err := json.Unmarshal(body, &response); err != nil {
		return "", errors.New("EJSONUNMARSHAL")
	}

	if response["Message"] != "" && response["Message"] == "There was an error saving your data, please try again." {
		return "", errors.New("EINCORRECTCREDS")
	} else {
		var USER_COOKIE string = response["CustomerToken"].(string)
		return "st=" + USER_COOKIE, nil
	}
}

/*
   POSTs the Chipotle API with parameters that ultimately return nearby Chipotle locations.

   IMPORTANT NOTE: The API can get very slow with large values of returnCount; we should eventually limit the size of it

*/

func GetLocations(postalCode string, radius string, returnCount string) (response []Location, err error) {
	b := bytes.NewBufferString(`{"Address": ` + postalCode + `, "Radius": ` + radius + `, "StartIndex": 1, "ReturnCount": ` + returnCount + `}`)

	req, _ := http.NewRequest("POST", BaseURL+"/api/restaurant/restaurantssearch", b)

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		return nil, errors.New("EHTTPREQUEST")
	}

	body, _ := ioutil.ReadAll(res.Body)

	locations := make([]Location, 0)
	json.Unmarshal(body, &locations)

	return locations, nil
}

/*
   POSTs the Chipotle API with parameters that ultimately return the user's recent orders.

*/

func GetRecentOrders(userToken string, locationID int) (response []Order) {
	req, _ := http.NewRequest("GET", BaseURL+"/api/order/recents/"+strconv.Itoa(locationID), nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", userToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	recentOrders := make([]Order, 0)
	json.Unmarshal(body, &recentOrders)

	return recentOrders

}

/*
func main() {
	userToken := Login("jbirdcaicedo@yahoo.com", "q*HGB@71^Rd9%w4bm#^C6mhApdpx$fQL")
	locationID := GetLocations("12180")[0].Id
	recentOrders := GetRecentOrders(userToken, locationID)
	fmt.Printf("%s", recentOrders[1].OrderName)
}
*/
