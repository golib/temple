package html

import (
	"testing"
)

func Test_Doctype(t *testing.T) {
	var (
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

		xmls = map[string]string{
			"xml":       `<?xml version="1.0" encoding="utf-8" ?>`,
			"xml utf-8": `<?xml version="1.0" encoding="utf-8" ?>`,
		}

		fallbacks = map[string]string{
			"1.1":    `<!DOCTYPE html>`,
			"mobile": `<!DOCTYPE html>`,
			"basic":  `<!DOCTYPE html>`,
		}
	)

	// xhtml
	for name, doctype := range xdoctypes {
		assetString(doctype, Doctype(name, FORMAT_XHTML), t)
	}

	// html
	for name, doctype := range doctypes {
		assetString(doctype, Doctype(name, FORMAT_HTML), t)
	}

	// xml
	for name, doctype := range xmls {
		assetString(doctype, Doctype(name, FORMAT_XHTML), t)
	}

	// fallback for xhtml with wrong format
	for name, doctype := range fallbacks {
		assetString(doctype, Doctype(name, FORMAT_HTML), t)
	}
}

func Test_NewWrapper(t *testing.T) {
	wrapper := NewWrapper()

	assetString(wrapper.L, "<!--\n//<![CDATA[\n", t)
	assetString(wrapper.R, "\n//]]>\n//-->", t)
}

func Test_NewCommentWrapper(t *testing.T) {
	wrapper := NewCommentWrapper()

	assetString(wrapper.L, "<!--\n", t)
	assetString(wrapper.R, "\n//-->", t)
}

func Test_NewCdataWrapper(t *testing.T) {
	wrapper := NewCdataWrapper()

	assetString(wrapper.L, "\n//<![CDATA[\n", t)
	assetString(wrapper.R, "\n//]]>\n", t)
}

func Test_IsAutoclose(t *testing.T) {
	var autoclose = map[string]bool{
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

	for tag, ok := range autoclose {
		assetBool(IsAutoclose(tag), ok, t)
	}
}

var (
	tagName  = "d"
	tagValue = "div"

	defaultTagValue = "div"
)

func Test_RegisterShortcutTag(t *testing.T) {
	RegisterShortcutTag(tagName, tagValue)
	assetString(ShortcutTag(tagName), tagValue, t)

	RegisterShortcutTag(tagName, "cannot overwrite registered tag")
	assetString(ShortcutTag(tagName), tagValue, t)
}

func Test_ShortcutTag(t *testing.T) {
	RegisterShortcutTag(tagName, tagValue)
	assetString(ShortcutTag(tagName), tagValue, t)

	assetString(ShortcutTag("unregistered tag"), defaultTagValue, t)
}

func Test_IsShortcutTag(t *testing.T) {
	RegisterShortcutTag(tagName, tagValue)
	assetBool(IsShortcutTag(tagName), true, t)

	assetBool(IsShortcutTag("unregistered tag"), false, t)
}

var (
	attrName  = "."
	attrValue = "class"

	defaultAttrValue = ""
)

func Test_RegisterShortcutAttr(t *testing.T) {
	RegisterShortcutAttr(attrName, attrValue)
	assetString(ShortcutAttr(attrName), attrValue, t)

	RegisterShortcutAttr(attrName, "cannot overwrite registered attr")
	assetString(ShortcutAttr(attrName), attrValue, t)
}

func Test_ShortcutAttr(t *testing.T) {
	RegisterShortcutAttr(attrName, attrValue)
	assetString(ShortcutAttr(attrName), attrValue, t)

	assetString(ShortcutAttr("unregistered attr"), defaultAttrValue, t)
}

func Test_IsShortcutAttr(t *testing.T) {
	RegisterShortcutAttr(attrName, attrValue)
	assetBool(IsShortcutAttr(attrName), true, t)

	assetBool(IsShortcutAttr("unregistered attr"), false, t)
}
