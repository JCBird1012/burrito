package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/JCBird1012/chipotle-cli/api"
	"github.com/JCBird1012/chipotle-cli/config"
	"github.com/jessfraz/weather/geocode"
	"github.com/jroimartin/gocui"
)

const ()

var (
	thing         string
	serverURI     string
	version       bool
	geo           geocode.Geocode
	location      string
	guac          bool
	license       bool
	test          bool
	isInteractive bool
	userToken     string
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
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.StringVar(&thing, "thing", "", "do thing and exit")
	flag.StringVar(&serverURI, "serverURI", defaultServerURI, "URL of the Chipotle API")
	flag.BoolVar(&guac, "guac", false, "Accept guac prompt")
	flag.BoolVar(&guac, "yes", false, "Accept guac prompt (alias)")
	flag.BoolVar(&guac, "y", false, "Accept guac prompt(alias)")

	flag.Usage = func() {
		fmt.Printf("This is the burrito CLI.\n")
		fmt.Printf("Usage: burrito-get [<options>][<arguments> ...]\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NFlag() < 1 {
		flag.Usage()
	}

	if serverURI == "" {
		usageAndExit("Enter a Chipotle API endpoint please or leave it blank", 0)
	}

	if test {
		fmt.Printf("This is a test\n")
		userToken = api.Login()
		fmt.Printf("userToken is: %s", userToken)
	}
	if isInteractive {
		fmt.Printf("is interactive: %t\n", isInteractive)
		interactive()
	}
}

func main() {
	if version {
		fmt.Printf("Chipotle CLI version %s,build %s", config.VERSION, config.GITCOMMIT)
		return
	}
	if len(thing) > 0 {

		fmt.Printf("we got a thing!\n")
		fmt.Printf("length of thing is: %d", len(thing))
		fmt.Printf(thing)
	} else {
		// fmt.Printf("we didn't get a thing!\n")
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
