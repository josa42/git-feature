package main

import (
	docopt "github.com/docopt/docopt-go"
	"github.com/josa42/git-feature/feature"
	stringutils "github.com/josa42/go-stringutils"
)

func main() {
	usage := stringutils.TrimLeadingTabs(`
		Usage:
		  git-feature <name>
		  git-feature list    [<filter>]
		  git-feature sync    [<name>]
		  git-feature publish [<name>]
		  git-feature delete  [<name>]

		Options:
		  -h --help          Show this screen.
		  --version          Show version.
  `)

	arguments, _ := docopt.Parse(usage, nil, true, "Git Cleanup 0.1.0", false)

	cmdType := "openOrCreate"
	cmdTypes := []string{
		"list",
		"sync",
		"publish",
		"delete",
	}

	for _, key := range cmdTypes {
		if arguments[key] == true {
			cmdType = key
		}
	}

	switch cmdType {
	case "openOrCreate":
		feature.CheckoutOrCreate()
	case "list":
		feature.List()
	case "sync":
		feature.Sync()
	case "publish":
		feature.Publish()
	case "delete":
		feature.Delete()
	}
}
