// Package proverb is designed to build a proverb from an array of strings.
package proverb

import "fmt"

// Proverb is a function to build an array of strings, with each line being made up
// of lines in the format "For want of a x the y was lost", except the last line, being
// "And all for the want of an x"
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	poem := make([]string, len(rhyme))
	for i := 0; i < (len(rhyme) - 1); i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		poem[i] = line
	}
	poem[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return poem
}
