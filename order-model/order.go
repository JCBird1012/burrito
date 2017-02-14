package order

import (
	"encoding/json"
	"fmt"

	"github.com/JCBird1012/chipotle-cli/api"
)

// Mealtypes defines what physical form your torilla will take
var Mealtypes = []string{"Burrito", "Bowl", "Salad", "Tacos"}

// Fillings enumerates the types of meats availible
var Fillings = []string{"Chorizo", "Sofritas", "Veggie", "Chicken", "Steak", "Barbacoa", "Carnitas"}

// Beans enumerates beans availible
var Beans = []string{"BeansBlack", "BeansPinto", "NoBeans"}

// Rice duh
var Rice = []string{"RiceWhite", "RiceBrown", "NoRice"}

// Toppings enumnerates the topings availible
var Toppings = []string{"GuacTopping", "SalsaTopping", "SalsaCorn", "SalsaGreenChili", "SalsaRedChili", "SourCream", "FajitaVeggies", "Cheese", "Lettuce"}

// Tortillas duh
var Tortillas = []string{"TortillaFlour", "TortillaCornSoft", "TortillaCornCrispy"}

// Drinks duh YMMV
var Drinks = []string{"SodaSmall", "SodaLarge", "BottledWater"}

// Sides chips and guake
var Sides = []string{"Chips", "GuacSide", "ChipsAndGuacamole", "ChipsAndSalsaTomato"}

// Order repesenting a standard chipotle order
type Order struct {
	Name     string
	Mealtype string
	Filling  string
	Beans    string
	Rice     string
	Toppings []string
	Drink    string
	Cheese   bool
	StoreID  int
}

// PrettyMenu does things
func PrettyMenu() string {
	var report string

	strMealTypes, _ := json.Marshal(Mealtypes)
	strBeans, _ := json.Marshal(Beans)
	strRice, _ := json.Marshal(Rice)
	strFillings, _ := json.Marshal(Fillings)
	strToppings, _ := json.Marshal(Toppings)
	strTortillas, _ := json.Marshal(Tortillas)
	strDrinks, _ := json.Marshal(Drinks)
	strSides, _ := json.Marshal(Sides)

	report = fmt.Sprintln(
		" + ---- Food Configuration: \n" +
			" - meals:\t" + string(strMealTypes) + "\n" +
			" - fillings:\t" + string(strFillings) + "\n" +
			" - beans:\t" + string(strBeans) + "\n" +
			" - rice:\t" + string(strRice) + "\n" +
			" - toppings:\t" + string(strToppings) + "\n" +
			" - tortillas:\t" + string(strTortillas) + "\n" +
			" + ---- Auxillary Items: " + "\n" +
			" - drinks:\t" + string(strDrinks) + "\n" +
			" - sides:\t" + string(strSides) + "\n")

	return report
}

// ProcessOrder does stuff
func ProcessOrder(order Order) bool {

	// Let user choose the location, or if defined used the
	//  command line location or storeID
	locations := api.GetLocations("12180")

	// Do the picking here
	// <picking>
	// for now just use the first one
	location := locations[0]

	// Prompt the user for their login credentials
	token := api.Login("username", "password")

	strOrder, _ := json.Marshal(order)
	return api.SendOrderToSChipotleByStoreID(location.ID, token, string(strOrder))
}
