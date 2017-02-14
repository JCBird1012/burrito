package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/JCBird1012/chipotle-cli/api"
	"github.com/JCBird1012/chipotle-cli/config"
	"github.com/JCBird1012/chipotle-cli/order-model"
	"github.com/JCBird1012/chipotle-cli/utils"
	"github.com/jessfraz/weather/geocode"
	"github.com/jroimartin/gocui"
	"github.com/op/go-logging"
)

const ()

var (
	version       bool
	geo           geocode.Geocode
	location      string
	guac          bool
	license       bool
	test          bool
	isInteractive bool
	userToken     string
	multi         bool
	fName         string
	fFilling      string
	fBeans        string
	fRice         string
	fToppings     []string
	fStrToppings  string
	fDrink        string
	fCheese       bool
	fStoreID      int
	fMealType     string
	fShowMenu     bool
)

const (
	defaultServerURI string = "https://thing.com"
)

// Setup logging
var log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}🌮 %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init() {
	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)

	flag.BoolVar(&isInteractive, "interactive", false, "interactive meal generation")
	flag.BoolVar(&isInteractive, "i", false, "interactive meal generation (alias)")
	flag.BoolVar(&test, "test", false, "do test thing")
	flag.BoolVar(&license, "license", false, "print out license information")
	flag.StringVar(&location, "location", "", "zipcode, city or state, defaults to ip location")

	flag.BoolVar(&guac, "guac", false, "Accept guac prompt")
	flag.BoolVar(&guac, "yes", false, "Accept guac prompt (alias)")
	flag.BoolVar(&guac, "y", false, "Accept guac prompt(alias)")

	// Parse burrito Options
	// fName         string
	// fFilling      string
	// fBeans        string
	// fRice         string
	// fToppings     []string
	// fDrink        string
	// fCheese       bool
	// fStoreID      int
	flag.StringVar(&fName, "name", "", "Sets the order display name")
	flag.StringVar(&fFilling, "filling", "", "Sets the filling type")
	flag.StringVar(&fFilling, "meat", "", "Sets the filling type (alias)")
	flag.StringVar(&fBeans, "beans", "", "Sets the bean type")
	flag.StringVar(&fRice, "rice", "", "Sets the rice type")
	flag.StringVar(&fStrToppings, "toppings", "SalsaTopping", "Sets the topping type")
	flag.StringVar(&fDrink, "drink", "", "Sets the drink type")
	flag.BoolVar(&fCheese, "cheese", false, "Sets the cheese preference")
	flag.IntVar(&fStoreID, "storeID", 0, "Sets the store id to order from")
	flag.StringVar(&fMealType, "meal", "Burrito", "Sets the meat type")
	flag.BoolVar(&fShowMenu, "menu", true, "Shows the availible menu")

	flag.Usage = func() {
		fmt.Printf("Usage: chipotle-cli [Mealtype][<options> ...]\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(os.Args) < 1 {
		usageAndExit("Insufficient flags", 0)
	}
	if fShowMenu {
		fmt.Println(order.PrettyMenu())
		os.Exit(0)
	}
	if test {
		log.Debug("This is a test\n")
		userToken = api.Login("username", "password")
		log.Debug("userToken is: %s", userToken)

		fmt.Printf("This is something else %s \n", os.Args[2])
		if utils.IsInArray(os.Args[2], order.Mealtypes) {
			log.Debug("WE FOUND IT\n")
		} else {
			log.Debug("we didnt find it\n")
		}
	}
	// begin parse tree
	if utils.IsInArray(fMealType, order.Mealtypes) {
		// They used the command to orbder a single meal
		// see if things are defined or else pick sensible defaults
		if len(fName) <= 0 {
			fName = "Anonymous_" + fMealType

			log.Warning("You didn't specify a name for your " + fMealType + "\n" +
				"\t - using sane default, \"" + fName + "\"")
		}
		if len(fFilling) <= 0 || !utils.IsInArray(fFilling, order.Fillings) {
			fFilling = order.Fillings[0]

			log.Warning(" You didn't specify a filling for your " + fMealType + "\n" +
				"\t - using sane default, \"" + fFilling + "\"")
		}
		if len(fBeans) <= 0 || !utils.IsInArray(fBeans, order.Beans) {
			fBeans = order.Beans[0]

			log.Warning(" You didn't specify a bean type for your " + fMealType + "\n" +
				"\t - using sane default, \"" + fBeans + "\"")
		}
		if len(fRice) <= 0 || !utils.IsInArray(fRice, order.Rice) {
			fRice = order.Rice[0]

			log.Warning(" You didn't specify a rice type for your " + fMealType + "\n" +
				"\t - using sane default, \"" + fRice + "\"")
		}
		if len(fStrToppings) <= 0 {
			fToppings = []string{order.Toppings[0]}

			log.Warning(" You didn't specify any toppings  for your " + fMealType + "\n" +
				"\t - using sane default, \"" + fToppings[0] + "\"")
		} else {
			fToppings = strings.Split(fStrToppings, ",")
		}
		if len(fDrink) <= 0 {
			log.Warning(" you didn't specify a drink for your meal" + "\n" +
				"\t - that's fine... I hope it's not spicy")
		}
		if fStoreID == 0 {
			// TODO:nag for root store id, probaly not legit etc
		}

		myOrder := order.Order{
			Name:     fName,
			Mealtype: fMealType,
			Filling:  fFilling,
			Beans:    fBeans,
			Rice:     fRice,
			Toppings: fToppings,
			Drink:    fDrink,
			Cheese:   fCheese,
			StoreID:  fStoreID}
		myJSONObj, err := json.MarshalIndent(myOrder, "", "    ")
		if err != nil {
			fmt.Print(err)
		}
		log.Notice("One " + myOrder.Name + ", coming up!" + "\n" +
			string(myJSONObj))
		c := utils.AskForConfirmation("Is this order correct?")
		if c {
			log.Notice("sweet... ordering.")]
			order.send(myOrder)
		} else {
			log.Error("order is not correct, sad")
		}

	} else {
		// They didn't , check if it's a multi and expect the appropriate params
		if os.Args[1] == "multi" {
			// Yep it's multi, are the params there?
			if os.Args[2] == "path" {
				// yep, the params are there, does it open a valid file location?
				// TODO: test file validity
			}
		}
		// They didnt go for multi, maybe they want an interactive shell?
		if isInteractive {
			fmt.Printf("is interactive: %t\n", isInteractive)
			interactive()
		}
	}
	// Huh, not interactive, and we're fresh out of ideas..
	// let's just print some stuff out and hope they go away
	// flag.Usage()
}

func main() {

	// log.Debugf("debug %s", Password("secret"))
	// log.Info("info")
	// log.Notice("notice")
	// log.Warning("warning")
	// log.Error("err")
	// log.Critical("crit")
	if version {
		fmt.Printf("Chipotle CLI version %s,build %s", config.VERSION, config.GITCOMMIT)
		return
	}

	var err error
	if location == "" {
		geo, err = geocode.Autolocate()
		if err != nil {
			fmt.Print(err)
		}
	}
	log.Debug("Zipcode is : " + geo.PostalCode)
}

//with respect to jessfraz/weather/blog/master/main.go:117
func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprint(os.Stderr, message)
		fmt.Fprint(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprint(os.Stderr, "\n")
	os.Exit(exitCode)
}

func interactive() {
	// g, err := gocui.NewGui(gocui.OutputNormal)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// defer g.Close()
	//
	// g.SetManagerFunc(layout)
	//
	// if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
	// 	log.Panicln(err)
	// }
	//
	// if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
	// 	log.Panicln(err)
	// }
}
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if maxX < 10 || maxY < 10 {
		fmt.Printf("Your terminal is too small for interactive mode :(\n")
		return nil
	}
	if leftpane, err := g.SetView("leftpane", 0, 0, 20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		leftpane.Title = "Orders"
		fmt.Fprintln(leftpane, "Buuurito!")
		fmt.Fprint(leftpane, "anotha one")
	}
	if rightpane, err := g.SetView("rightpane", 21, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		rightpane.Title = "details"
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
