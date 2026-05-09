package pipeline

import (
	"compiler/core/source"
	"compiler/internal/context"
	"compiler/internal/frontend/ast"
	"compiler/internal/tokens"
)

// Source text to tokens.
// TODO: replace EOF scaffold with the real lexer.
func lex(module *context.Module) []tokens.Token {
	if module == nil {
		return nil
	}
	pos := source.NewPosition()
	return []tokens.Token{{
		Kind:    tokens.EOF,
		Literal: "",
		Start:   pos,
		End:     pos,
	}}
}

// Tokens to frontend AST.
// Should return partial ASTs after recoverable syntax errors.
func parse(module *context.Module, _ []tokens.Token) *ast.Module {
	if module == nil {
		return nil
	}
	return &ast.Module{
		FilePath: module.FilePath,
		Decls:    make([]ast.Decl, 0),
		Imports:  make([]*ast.ImportDecl, 0),
	}
}

// Collector, resolver, type checker, CTFE, and related semantic passes.
func analyze(_ *context.Module, _ *ast.Module) bool {
	return true
}

// Checked AST/semantic data to high-level IR.
func lowerHIR(module *context.Module, _ *ast.Module) string {
	if module == nil {
		return ""
	}
	return "; hir module " + module.ImportPath + "\n"
}

// HIR to target-independent mid-level IR.
func lowerMIR(module *context.Module, _ string) string {
	if module == nil {
		return ""
	}
	return "; mir module " + module.ImportPath + "\n"
}

// MIR to LLVM IR.
func lowerLLVMIR(module *context.Module, _ string) string {
	if module == nil {
		return ""
	}
	return "; llvm ir module " + module.ImportPath + "\n"
}
