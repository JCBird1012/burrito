package utils

import "fmt"

// IsInArray linearn search if thing is in array
func IsInArray(thing string, things []string) bool {
	for i := 0; i < len(things); i++ {
		if things[i] == thing {
			return true
		}
		fmt.Printf("%s is %s : %t\n", things[i], thing, things[i] == thing)
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
