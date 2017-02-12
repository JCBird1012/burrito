package order

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
}
