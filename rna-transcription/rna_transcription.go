// Package strand translates a string of DNA in to RNA
package strand

// ToRNA replaces each character in a string with its RNA equivalent
func ToRNA(dna string) string {
	transcribe := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	var rna string
	for _, rune := range dna {
		rna = rna + string(transcribe[rune])
	}
	return rna
}
