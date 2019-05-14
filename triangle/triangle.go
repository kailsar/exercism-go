package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota
	Sca
	Iso
	Equ
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a < 0 || b < 0 || c < 0 {
		return
	}
	var k Kind
	return k
}
