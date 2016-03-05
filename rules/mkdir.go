package rules

import (
	"github.com/jessevdk/go-flags"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
	"github.com/nopecmd/nope/parse"
	"strings"
)

const mkdirBaseCommand = "mkdir"

func isMatchMkdir(cmd models.Command) bool {
	return cmd.BaseCommand == mkdirBaseCommand
}

func getUndoMkdir(cmd models.Command) string {
	var mkdirFlags struct {
		Mode         string `short:"m"`
		CreateInterm bool   `short:"p"`
		Verbose      bool   `short:"v"`
	}
	filteredTokens, err := flags.ParseArgs(&mkdirFlags, cmd.Tokens[1:])
	if err != nil {
		panic(err)
	}
	var dirPaths = parse.GetFilePathsFromTokens(filteredTokens)

	return rmBaseCommand + " -rf " + strings.Join(dirPaths, " ")
}

func init() {
	match.AddRule(isMatchMkdir, getUndoMkdir)
}
