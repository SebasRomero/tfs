package internal

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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
	_, err := uploadFiles(flag.Args())
	if err != nil {
		log.Panic(err.Error())
	}
}

func pull() {

	res, err := http.Get("http://localhost:8080/api/v1/pull/12")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res.Body)
}

func help() {
	fmt.Println(types.Help)
}
