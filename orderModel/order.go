package order

import (
	"encoding/json"
	"fmt"
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

	strMealTypes, err := json.Marshal(Mealtypes)
	strBeans, err := json.Marshal(Beans)
	strRice, err := json.Marshal(Rice)
	strFillings, err := json.Marshal(Fillings)
	strToppings, err := json.Marshal(Toppings)
	strTortillas, err := json.Marshal(Tortillas)
	strDrinks, err := json.Marshal(Drinks)
	strSides, err := json.Marshal(Sides)
	if err != nil {
		return "Error parsing menu"
	}

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
