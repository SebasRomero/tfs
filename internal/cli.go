package internal

import (
	"flag"
	"fmt"

	"github.com/sebasromero/tfs/internal/types"
)

func Cli() {
	flag.Parse()

	if flag.Arg(0) == "push" {
		push()
	} else if flag.Arg(0) == "pull" {
		pull()
	} else {
		help()
	}

}

func push() {
	fmt.Print("push")
}

func pull() {
	fmt.Print("pull")
}

func help() {
	fmt.Println(types.Help)
}
