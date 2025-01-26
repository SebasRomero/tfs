# TFS - Temporary File Sharing

## About

This CLI application will help you to quickly share files and get them immediately.  

## Installation

Download the appropriate binary for your operating system:
   - [Linux (amd64)](https://github.com/sebasromero/tfs/releases/download/0.1.0/tfs)
   - [macOS (amd64)](https://github.com/sebasromero/tfs/releases/download/0.1.0/tfs_ios)
   - [Windows (amd64)](https://github.com/sebasromero/tfs/releases/download/0.1.0/tfs.exe)

Make the binary executable (if needed):


## How to use
To upload files:

```bash
./tfs push <files>
```

To get files:

```shell
./tfs pull <directory destination> <directory code>
```

## Example
```shell
./tfs push my_pdf.pdf text.txt
```

```shell
./tfs pull ~/Desktop 4f3vs2
```