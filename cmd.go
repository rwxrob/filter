package filter

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/filter/youtube"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `filter`,
	Aliases:   []string{`filt`},
	Summary:   `collection of common UNIX filters`,
	Commands:  []*Z.Cmd{help.Cmd, youtube.Cmd},
	Copyright: `Copyright 2022 Robert S Muhlestein`,
	Version:   `v0.1.0`,
	License:   `Apache-2.0`,
	Site:      `rwxrob.tv`,
	Source:    `git@github.com:rwxrob/filter.git`,
	Issues:    `github.com/rwxrob/filter/issues`,
}
