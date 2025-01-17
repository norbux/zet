package err

import "log"

// For is a function to check for and error and issue a log.Fatal in case it's not nil
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
