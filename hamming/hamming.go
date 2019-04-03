// hamming is a package that calculates the Hamming distance between two strands of 
// DNA: that is, the number of nucleotides that are different at the same index of
// the DNA strands.
package hamming

import "fmt"

// Distance is a function that checks that the two given strings are of equal length,
// and then calculates the number of positions in those strings that are different.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Strings are not equal")
	}
	counter := 0
	for i :=0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			counter ++
		}

	}
return counter, nil
}
