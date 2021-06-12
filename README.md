# IITK Coin
## SnT Project 2021, Programming Club 

This repository contains the code for the IITK Coin project done by Gurbaaz Singh Nandra.

## Table Of Content
- [Development Environment](#development-environment)
- [Directory Structure](#directory-structure)
- [Usage](#usage)

## Development Environment

```bash
- go version: go1.16.4 linux/amd64    # https://golang.org/dl/
- system: 5.4.72-microsoft-standard-WSL2 x86_64    # https://docs.microsoft.com/en-us/windows/wsl/install-win10
- text editor: VSCode    # https://code.visualstudio.com/download
- terminal: Zsh     # https://code.visualstudio.com/download
```

## Directory Structure
```
..
├── README.md
├── Task-1
│   ├── iitkcoin-190349.db
│   └── main.go
├── Task-2
│   ├── database.go
│   ├── main.go
│   └── routes.go
├── go.mod
└── go.sum

2 directories, 8 files
```

## Usage
```bash
cd $GOPATH/src/github.com/<username>
git clone https://github.com/gurbaaz27/iitk-coin.git
go get github.com/mattn/go-sqlite3
cd Task-<no>/
go run main.go
```
