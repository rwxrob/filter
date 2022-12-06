package youtube

import (
	_ "embed"
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `youtube`,
	Aliases:   []string{`yt`},
	Summary:   `collection of common UNIX filters for YouTube`,
	Copyright: `Copyright 2022 Robert S Muhlestein`,
	Version:   `v0.1.0`,
	License:   `Apache-2.0`,
	Site:      `rwxrob.tv`,
	Source:    `git@github.com:rwxrob/filter.git`,
	Issues:    `github.com/rwxrob/filter/issues`,
	Commands:  []*Z.Cmd{help.Cmd, LinkifyCmd},
}

//go:embed linkify.md
var linkifyDesc string

var LinkifyCmd = &Z.Cmd{
	Name:        `linkify`,
	Aliases:     []string{`link`},
	Usage:       `[help|PATH]`,
	Summary:     `add angle brackets and optional prefix to YT links`,
	Copyright:   `Copyright 2022 Robert S Muhlestein`,
	Version:     `v0.1.0`,
	License:     `Apache-2.0`,
	Site:        `rwxrob.tv`,
	Source:      `git@github.com:rwxrob/filter.git`,
	Issues:      `github.com/rwxrob/filter/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: linkifyDesc,

	Call: func(x *Z.Cmd, args ...string) error {
		buf := Z.ArgsOrIn(args)
		fmt.Print(Linkify(buf))
		return nil
	},
}
