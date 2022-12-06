package youtube

import (
	"regexp"

	"github.com/rwxrob/to"
)

// Example: https://youtu.be/8st1NhaKDCA
var LinkExp = regexp.MustCompile(`([^<]|^)(https://youtu\.be\/\S+)`)

// Linkify converts any YouTube URL within the input into a valid
// markdown URL wrapped with angle brackets and prefixed with the
// optionally passed prefix (default is 📺). Pass an empty string if no
// prefix is wanted. Any input that to.String accepts is valid. Returns
// empty string if no arguments are passed.
func Linkify(args ...any) string {

	var in, pre string
	ln := len(args)

	switch {

	case ln >= 2:
		pre = to.String(args[1])
		fallthrough

	case ln == 1:

		pre = "📺 "
		in = to.String(args[0])

	case ln == 0:
		return ""

	}

	return LinkExp.ReplaceAllString(in, `$1`+pre+`<$2>`)
}
