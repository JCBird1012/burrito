package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// IsInArray linearn search if thing is in array
func IsInArray(thing string, things []string) bool {
	for i := 0; i < len(things); i++ {
		if things[i] == thing {
			return true
		}
		// fmt.Printf("%s is %s : %t\n", things[i], thing, things[i] == thing)
	}
	return false
}

// IndexOfThing duh
func IndexOfThing(thing string, things []string) int {
	for i := 0; i < len(things); i++ {
		if things[i] == thing {
			return i
		}
		i = i + 1
	}
	return -1
}

// AllInArray returns the things not in the array
func AllInArray(things []string, others []string) []string {
	// TODO: use a map because this is slow (I mean it's not but it it is)
	var report []string
	for i := 0; i < len(things); i++ {
		if !IsInArray(things[i], others) {
			report = append(report, things[i])
		}
	}
	return report
}

// AskForConfirmation does things
// with respect to m4ng0squ4sh/confirm.go https://gist.github.com/m4ng0squ4sh/3dcbb0c8f6cfe9c66ab8008f55f8f28b
func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
