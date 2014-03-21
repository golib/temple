package html

import (
	"regexp"
)

var (
	rxml = regexp.MustCompile(`^xml(\s+(.+?))?$`)

	xdoctypes = map[string]string{
		"1.1":          `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`,
		"5":            `<!DOCTYPE html>`,
		"html":         `<!DOCTYPE html>`,
		"strict":       `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`,
		"frameset":     `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">`,
		"mobile":       `<!DOCTYPE html PUBLIC "-//WAPFORUM//DTD XHTML Mobile 1.2//EN" "http://www.openmobilealliance.org/tech/DTD/xhtml-mobile12.dtd">`,
		"basic":        `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN" "http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">`,
		"transitional": `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`,
	}

	doctypes = map[string]string{
		"5":            `<!DOCTYPE html>`,
		"html":         `<!DOCTYPE html>`,
		"strict":       `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
		"frameset":     `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN" "http://www.w3.org/TR/html4/frameset.dtd">`,
		"transitional": `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">`,
	}
)

// Doctype returns the actual string of shortcut.
func Doctype(name, format string) string {
	var (
		dt = ""
		ok = false
	)

	if rxml.MatchString(name) {
		if format == FORMAT_HTML {
			panic("Invalid xml directive with html format")
		}

		encoding := rxml.FindAllString(name, 1)
		if len(encoding) == 1 {
			encoding = append(encoding, "utf-8")
		}

		dt = `<?xml version="1.0" encoding="` + encoding[1] + `" ?>`
		ok = true
	} else {
		switch format {
		case FORMAT_HTML:
			dt, ok = doctypes[name]
		case FORMAT_XHTML:
			dt, ok = xdoctypes[name]
		}
	}

	if ok == false {
		dt = doctypes["5"]
	}

	return dt
}

type Wrapper struct {
	L string
	R string
}

func NewWrapper() *Wrapper {
	return &Wrapper{"<!--\n//<![CDATA[\n", "\n//]]>\n//-->"}
}

func NewCommentWrapper() *Wrapper {
	return &Wrapper{"<!--\n", "\n//-->"}
}

func NewCdataWrapper() *Wrapper {
	return &Wrapper{"\n//<![CDATA[\n", "\n//]]>\n"}
}

var _AUTOCLOSES = map[string]bool{
	"base":     true,
	"basefont": true,
	"bgsound":  true,
	"link":     true,
	"meta":     true,
	"area":     true,
	"br":       true,
	"embed":    true,
	"img":      true,
	"keygen":   true,
	"wbr":      true,
	"input":    true,
	"menuitem": true,
	"param":    true,
	"source":   true,
	"track":    true,
	"hr":       true,
	"col":      true,
	"frame":    true,
}

// Whether is the tag autoclose?
func IsAutoclose(name string) bool {
	if _, ok := _AUTOCLOSES[name]; ok {
		return true
	}

	return false
}

var (
	rword      = regexp.MustCompile(`[\w-]`)
	defaultTag = "div"

	_SHORTCUT_TAGS  = map[string]string{}
	_SHORTCUT_ATTRS = map[string]string{}
)

func RegisterShortcutTag(shortcut, name string) {
	if _, ok := _SHORTCUT_TAGS[shortcut]; ok {
		return
	}

	_SHORTCUT_TAGS[shortcut] = name
}

func ShortcutTag(shortcut string) string {
	if tag, ok := _SHORTCUT_TAGS[shortcut]; ok {
		return tag
	}

	// default to div
	return defaultTag
}

func IsShortcutTag(shortcut string) bool {
	if _, ok := _SHORTCUT_TAGS[shortcut]; ok {
		return true
	}

	return false
}

func RegisterShortcutAttr(shortcut, name string) {
	if rword.MatchString(shortcut) {
		panic("Only special characters are accepted for attribute shortcuts")
	}

	if _, ok := _SHORTCUT_ATTRS[shortcut]; ok {
		return
	}

	// use default tag for attr shortcut
	RegisterShortcutTag(shortcut, defaultTag)

	_SHORTCUT_ATTRS[shortcut] = name
}

func ShortcutAttr(shortcut string) string {
	if attr, ok := _SHORTCUT_ATTRS[shortcut]; ok {
		return attr
	}

	// default to blank
	return ""
}

func IsShortcutAttr(shortcut string) bool {
	if _, ok := _SHORTCUT_ATTRS[shortcut]; ok {
		return true
	}

	return false
}
