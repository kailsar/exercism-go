// Twofer is a package to return "One for X, one for me.", where x is the parameter given. 
// If no parameter is given,return "One for you, one for me."
package twofer

import (
	"fmt"
)

// Sharewith is the main function returning the string "One for X, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
