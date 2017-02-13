package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/JCBird1012/chipotle-cli/api"
	"github.com/JCBird1012/chipotle-cli/config"
	"github.com/JCBird1012/chipotle-cli/orderModel"
	"github.com/JCBird1012/chipotle-cli/utils"
	"github.com/jessfraz/weather/geocode"
	"github.com/jroimartin/gocui"
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
	fDrink        string
	fCheese       bool
	fStoreID      int
)

const (
	defaultServerURI string = "https://thing.com"
)

func init() {
	flag.BoolVar(&isInteractive, "interactive", false, "interactive meal generation")
	flag.BoolVar(&isInteractive, "i", false, "interactive meal generation (alias)")
	flag.BoolVar(&test, "test", false, "do test thing")
	flag.BoolVar(&license, "license", false, "print out license information")
	flag.StringVar(&location, "location", "", "zipcode, city or state, defaults to ip location")

	flag.BoolVar(&guac, "guac", false, "Accept guac prompt")
	flag.BoolVar(&guac, "yes", false, "Accept guac prompt (alias)")
	flag.BoolVar(&guac, "y", false, "Accept guac prompt(alias)")
	flag.Usage = func() {
		fmt.Printf("Usage: burrito-get [Mealtype][<options> ...]\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(os.Args) < 1 {
		usageAndExit("Insufficient flags", 0)
	}
	if test {
		fmt.Printf("This is a test\n")
		userToken = api.Login("username", "password")
		fmt.Printf("userToken is: %s", userToken)

		fmt.Printf("This is something else %s \n", os.Args[2])
		if utils.IsInArray(os.Args[2], order.Mealtypes) {
			fmt.Print("WE FOUND IT\n")
		} else {
			fmt.Print("we didnt find it\n")
		}
	}
	// begin parse tree
	if utils.IsInArray(os.Args[1], order.Mealtypes) {
		// They used the command to orbder a single meal
		// see if things are defined or else pick sensible defaults
		if len(fName) <= 0 {
			fName = "anonymous_" + os.Args[1]

			fmt.Println("[WARN] You didn't specify a name for your " + os.Args[1])
			fmt.Println("\t - using sane default, \"" + fName + "\"")
		}
		if len(fFilling) <= 0 || !utils.IsInArray(fFilling, order.Fillings) {
			fFilling = order.Fillings[0]

			fmt.Println("[WARN] You didn't specify a filling for your " + os.Args[1])
			fmt.Println("\t - using sane default, \"" + fFilling + "\"")
		}
		if len(fBeans) <= 0 || !utils.IsInArray(fBeans, order.Beans) {
			fBeans = order.Beans[0]

			fmt.Println("[WARN] You didn't specify a bean type for your " + os.Args[1])
			fmt.Println("\t - using sane default, \"" + fBeans + "\"")
		}
		if len(fRice) <= 0 || !utils.IsInArray(fRice, order.Rice) {
			fRice = order.Rice[0]

			fmt.Println("[WARN] You didn't specify a rice type for your " + os.Args[1])
			fmt.Println("\t - using sane default, \"" + fRice + "\"")
		}
		if len(fToppings) <= 0 {
			fToppings = []string{order.Toppings[0]}

			fmt.Println("[WARN] You didn't specify a for  toppings list for your " + os.Args[1])
			fmt.Println("\t - using sane default, \"" + fToppings[0] + "\"")
		}
		if len(fDrink) <= 0 {
			fmt.Println("[WARN] you didn't specify a drink for your meal")
			fmt.Println("\t - that's fine... I hope it's not spicy")
		}
		if fStoreID == 0 {
			// TODO:nag for root store id, probaly not legit etc
		}
		myOrder := order.Order{
			Name:     fName,
			Mealtype: os.Args[1],
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
		fmt.Println("[INFO] One " + myOrder.Name + ", coming up!")
		fmt.Println(string(myJSONObj))
		fmt.Println("Is this order correct? [y/N]")
		// r  := fmt.Scan(a)

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
	fmt.Printf("Zipcode is : %s\n", geo.PostalCode)
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
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
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
