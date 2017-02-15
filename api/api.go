package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const BASE_URL string = "https://order.chipotle.com"

type loginInformation struct {
	Username string
	Password string
	Persist  bool
}

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

func Login(username string, password string) (userToken string) {
	login := &loginInformation{
		Username: username,
		Password: password,
		Persist:  false}

	loginJSON, _ := json.Marshal(login)

	b := bytes.NewBufferString(string(loginJSON))

	req, _ := http.NewRequest("POST", BASE_URL+"/api/customer/login", b)

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

	var USER_COOKIE string = response["CustomerToken"].(string)
	return "st=" + USER_COOKIE
}

func GetLocations(postalCode string) (response []Location) {
	b := bytes.NewBufferString(`{"Address": ` + postalCode + `, "Radius":50, "StartIndex":1, "ReturnCount": 5}`)

	req, _ := http.NewRequest("POST", BASE_URL+"/api/restaurant/restaurantssearch", b)

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

func GetRecentOrders(userToken string, locationID int) (response []Order) {
	req, _ := http.NewRequest("GET", BASE_URL+"/api/order/recents/"+strconv.Itoa(locationID), nil)

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
