# TFS - Temporary File Sharing

## About

This CLI application will help you to quickly share files and get them immediately.  

### How it works?
 After download the CLI tool, to update or push any files use the push command, with this command you will be making a request via HTTPS to the API, this API will be in charge of encrypt using [Block Cipher](https://en.wikipedia.org/wiki/Block_cipher) algorithm and then, sending these files to an S3 Bucket, returning a "directory code" that we will be using later to get or pull these files.

The process getting these files would be similar, calling with the CLI tool we will be doing a request to the API passing as commands the directory where we want to download the files and the "directory code" obtained before, this call will get the files from the S3 Bucket, decrypting them and being download into your system.

## Installation

Download the appropriate binary for your operating system:
   - [Linux (amd64)](https://github.com/sebasromero/tfs/releases/download/0.2.0/tfs)
   - [macOS (amd64)](https://github.com/sebasromero/tfs/releases/download/0.2.0/tfs_ios)
   - [Windows (amd64)](https://github.com/sebasromero/tfs/releases/download/0.2.0/tfs.exe)

Make the binary executable (if needed):

```bash
chmod +x tfs
```

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
