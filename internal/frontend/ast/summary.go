package ast

import "fmt"

// Compact declaration text for CLI output.
func DeclSummary(decl Decl) string {
	if decl == nil {
		return "<nil decl>"
	}
	return fmt.Sprintf("%T", decl)
}
