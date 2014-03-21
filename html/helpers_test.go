package html

import (
	"testing"
)

func Test_IsBuiltinHelper(t *testing.T) {
	var builtins = map[string]bool{
		"len":       true,
		"print":     true,
		"printf":    true,
		"println":   true,
		"urlquery":  true,
		"js":        true,
		"json":      true,
		"index":     true,
		"html":      true,
		"unescaped": true,
	}

	for fn, ok := range builtins {
		assetBool(IsBuiltinHelper(fn), ok, t)
	}
}

func Test_HelperFuncs(t *testing.T) {
	ClearHelperFuncs()

	assetInt(len(HelperFuncs()), 0, t)
}

func Test_RegisterHelperFunc(t *testing.T) {
	ClearHelperFuncs()

	RegisterHelperFunc("upcase", func(input string) string {
		return input
	})

	assetInt(len(HelperFuncs()), 1, t)
}

func Test_UnregisterHelperFunc(t *testing.T) {
	ClearHelperFuncs()

	_, ok := UnregisterHelperFunc("unregistered helper func")
	assetBool(ok, false, t)

	RegisterHelperFunc("downcase", func(input string) string {
		return input
	})

	_, ok = UnregisterHelperFunc("downcase")
	assetBool(ok, true, t)
}

func Test_ClearHelperFuncs(t *testing.T) {
	ClearHelperFuncs()

	RegisterHelperFunc("camelcase", func(input string) string {
		return input
	})
	assetInt(len(HelperFuncs()), 1, t)

	ClearHelperFuncs()
	assetInt(len(HelperFuncs()), 0, t)
}
