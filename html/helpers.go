package html

var _HELPER_BUILTINS = map[string]bool{
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

func IsBuiltinHelper(name string) bool {
	if _, ok := _HELPER_BUILTINS[name]; ok {
		return true
	}

	return false
}

var _HELPER_FUNCS = map[string]interface{}{}

func HelperFuncs() map[string]interface{} {
	return _HELPER_FUNCS
}

func RegisterHelperFunc(name string, helper interface{}) {
	if _, ok := _HELPER_FUNCS[name]; ok {
		return
	}

	_HELPER_FUNCS[name] = helper
}

func UnregisterHelperFunc(name string) (interface{}, bool) {
	helper, ok := _HELPER_FUNCS[name]
	if ok {
		delete(_HELPER_FUNCS, name)
	}

	return helper, ok
}

func ClearHelperFuncs() {
	if len(_HELPER_FUNCS) == 0 {
		return
	}

	for name, _ := range _HELPER_FUNCS {
		delete(_HELPER_FUNCS, name)
	}
}
