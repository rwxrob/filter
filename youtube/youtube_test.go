package youtube_test

import (
	"fmt"

	"github.com/rwxrob/filter/youtube"
)

func ExampleLinkExp() {

	fmt.Println(youtube.LinkExp.MatchString(`https://youtu.be/8st1NhaKDCA`))
	fmt.Println(youtube.LinkExp.MatchString(`<https://youtu.be/8st1NhaKDCA>`))
	fmt.Println(youtube.LinkExp.MatchString(`https://youtube/8st1NhaKDCA`))
	fmt.Println(youtube.LinkExp.MatchString(`https://youtu.be/`))

	// Output:
	// true
	// false
	// false
	// false

}

func ExampleLinkify() {

	text := `
    This is a paragraph with https://youtu.be/inpara in it.

		* https://youtu.be/inlist

		https://youtu.be/ownblock

`

	fmt.Printf("%q\n", youtube.Linkify(text))

	// Output:
	// "\n    This is a paragraph with 📺 <https://youtu.be/inpara> in it.\n\n\t\t* 📺 <https://youtu.be/inlist>\n\n\t\t📺 <https://youtu.be/ownblock>\n\n"

}

func ExampleLinkify_ignore_Angled() {

	text := `
    This is a paragraph with <https://youtu.be/inpara> in it.
`

	fmt.Println(youtube.Linkify(text))

	// Output:
	// This is a paragraph with <https://youtu.be/inpara> in it.

}
