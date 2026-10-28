// Package diff computes the differences between two sequences of lines.
package diff

// Op is the kind of change an Edit represents.
type Op int

const (
	Equal  Op = iota // line is in both inputs
	Delete           // line is only in the old input
	Insert           // line is only in the new input
)

// String makes Ops readable in test failures and debug output.
func (o Op) String() string {
	switch o {
	case Equal:
		return "Equal"
	case Delete:
		return "Delete"
	case Insert:
		return "Insert"
	}
	return "Unknown"
}

// Edit is one line of a diff: what happened to it, and its text.
type Edit struct {
	Op   Op
	Text string
}

// Diff compares a and b index by index. It cannot detect inserted or
// removed lines; that needs a real algorithm such as LCS.
func Diff(a, b []string) []Edit {
	var edits []Edit

	for i := 0; i < max(len(a), len(b)); i++ {
		inA := i < len(a)
		inB := i < len(b)

		if inA && inB && a[i] == b[i] {
			edits = append(edits, Edit{Equal, a[i]})
			continue
		}
		if inA {
			edits = append(edits, Edit{Delete, a[i]})
		}
		if inB {
			edits = append(edits, Edit{Insert, b[i]})
		}
	}

	return edits
}
