package types

var Dst string = "The directory to download the files"

var Help string = `
Usage: ./tfs COMMAND [OPTIONS]

A Temporary File Sharing tool

Commands: 
	push	Command to upload the files
	pull	Command to download the files

Run './tfs COMMAND --help' for more information about the command.
`

var HelpPull string = `
Usage: ./tfs pull <directory path>
`
