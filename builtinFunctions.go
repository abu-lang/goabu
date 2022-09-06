package goabu

import "github.com/hyperjumptech/grule-rule-engine/ast"

// builtinFunctions provides the built-in functions usable in GoAbU's rules.
// The built-in functions consist in the exported methods of the builtInFunctions
// type including the built-in functions provided by Grule.
// The exported methods can be directly called in the rules (e.g. "AbsInt(-1)").
type builtinFunctions struct {
	*ast.BuiltInFunctions
}

func makeBuiltinFunctions(kb *ast.KnowledgeBase, wm *ast.WorkingMemory, dc ast.IDataContext) builtinFunctions {
	return builtinFunctions{
		BuiltInFunctions: &ast.BuiltInFunctions{
			Knowledge:     kb,
			WorkingMemory: wm,
			DataContext:   dc,
		},
	}
}

// AbsInt returns the absolute value of the argument.
func (f builtinFunctions) AbsInt(arg int64) int64 {
	if arg < 0 {
		return -arg
	}
	return arg
}
