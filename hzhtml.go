package hzutils

import "html"

// HTMLPre return html of string.
// @param shtml
// @return string
func HTMLPre(shtml string) string {
	return `<html>` + html.EscapeString(shtml) + `</html>`
}
