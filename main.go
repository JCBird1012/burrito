package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JCBird1012/chipotle-cli/config"
	"github.com/jessfraz/weather/geocode"
)

var (
	thing     string
	serverURI string
	version   bool
	geo       geocode.Geocode
	location  string
)

const (
	defaultServerURI string = "https://thing.com"
)

func init() {
	fmt.Printf("This is the burrito CLI.\n")
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.StringVar(&thing, "thing", "", "do thing and exit")
	flag.StringVar(&serverURI, "server URI", defaultServerURI, "URL of the Chipotle API")
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NFlag() < 1 {
		flag.Usage()
	}

	if serverURI == "" {
		usageAndExit("Enter a Chipotle API endpoint please or leave it blank", 0)
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
