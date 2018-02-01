package account

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"syscall"

	// Other burrito libraries
	"burrito/utils/api"

	// Terminal colors
	"github.com/fatih/color"

	// Useful for hiding typed passwords
	"golang.org/x/crypto/ssh/terminal"

	// .netrc parser for credentials
	"github.com/dickeyxxx/netrc"
)

func Login() {

	username, loggedIn := isUserLoggedIn()

	if loggedIn == 2 {
		fmt.Println("You're already logged in as " + color.CyanString(username))
		os.Exit(1)
	}

	// If .netrc doesn't exist, create it
	if loggedIn == 0 {
		err := createNetrc()

		if err != nil {
			color.Red("There was an error writing to ~/.netrc...")
			os.Exit(1)
		}
	}

	if loggedIn == 0 || loggedIn == 1 {
		username, password := askCreds()

		userCookie, err := api.Login(username, password)

		if err != nil {
			if err.Error() == "EINCORRECTCREDS" {
				color.Red("Authentication failure. Your credentials were incorrect.")
			}
			os.Exit(1)

		}

		netrc := getNetrc()

		if err != nil {
			color.Red("There was an error parsing your ~/.netrc")
		}

		netrc.AddMachine("order.chipotle.com", username, userCookie)
		netrc.Save()
		fmt.Println("Logged in as " + color.CyanString(username))
	}
}

func Logout() {

	username, loggedIn := isUserLoggedIn()

	if loggedIn == 2 {
		netrcFilepath, _ := getNetrcPath()
		n, _ := netrc.Parse(netrcFilepath)

		n.RemoveMachine("order.chipotle.com")

		fmt.Println("The local credentials for " + color.CyanString(username) + " have been removed.")
		os.Exit(1)
	} else if loggedIn == 0 || loggedIn == 1 {
		fmt.Println("No one is currently logged in.")
	}
}

// INTERNAL (PRIVATE) FUNCTIONS
func askCreds() (username string, password string) {

	fmt.Println("Enter your Chipotle credentials.")
	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password (typing will be hidden): ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	password = string(bytePassword)

	return

}

func getNetrc() *netrc.Netrc {

	netrcFilepath, err := getNetrcPath()

	if err != nil {
		color.Red("There was an error determining your .netrc")
		os.Exit(1)
	}

	if os.Stat(netrcFilepath); !os.IsNotExist(err) {
		fmt.Println(err)
		if err != nil {
			return nil
		}

		return netrc
	}
}

func isUserLoggedIn() (username string, loggedIn int) {

	// 0 = .netrc doesn't exist; 1 = .netrc exists, but doesn't contain Chipotle login info for any user; 2 = both .netrc and Chipotle login info exist
	loggedIn = 0
	username = ""

	// Check to see if .netrc exists

	netrc := getNetrc()

	if netrc != nil {

		// Chipotle entry exists
		if netrc.Machine("Chipotle") != nil {
			loggedIn = 2
			username = netrc.Machine("order.chipotle.com").Get("login")
		}

	}

	return
}

func getNetrcPath() (string, error) {
	usr, err := user.Current()

	if err != nil {
		return "", errors.New("EDETERMINENETRC")
	}

	netrcFilepath := filepath.Join(usr.HomeDir, ".netrc")

	return netrcFilepath, nil

}

func createNetrc() error {

	path, err := getNetrcPath()

	if err != nil {
		color.Red("There was an error determining your .netrc")
		os.Exit(1)
	}
	_, err = os.Create(path)
	os.Chmod(path, 0600)

	if err != nil {
		return err
	}

	return nil
}
