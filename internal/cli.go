package internal

import (
	"flag"
	"fmt"
	"log"

	"github.com/sebasromero/tfs/internal/types"
)

func Cli() {
	helpFlag := flag.Bool("help", true, "Show help")
	flag.Parse()

	if flag.Arg(0) == "push" {
		push()
	} else if flag.Arg(0) == "pull" {
		if flag.Arg(1) == "--help" || flag.Arg(1) == "-h" {
			helpPull()
		} else {
			pull(flag.Arg(1))
		}
	} else if *helpFlag {
		help()
	} else {
		help()
	}

}

func push() {
	_, err := uploadFiles(flag.Args())
	if err != nil {
		log.Panic(err.Error())
	}
}

func pull(dst string) {
	err := getFiles(dst)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func help() {
	fmt.Println(types.Help)
}

func helpPull() {
	fmt.Println(types.HelpPull)
}
