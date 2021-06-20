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
.
├── README.md
├── controllers
│   └── routes.go
├── database
│   └── database.go
├── go.mod
├── go.sum
├── iitkcoin-190349.db
├── main.go
└── models
    └── models.go

3 directories, 8 files
```

## Usage
```bash
cd $GOPATH/src/github.com/<username>
git clone https://github.com/gurbaaz27/iitk-coin.git
cd repo
go run main.go
```

## Endpoints

- `/login`
- `/signup`
- `/secretpage`
- `/reward`
- `/transfer`
- `/balance`


## Models

-  User
	Name     string `json:"name"`
	Rollno   int64  `json:"rollno,string"`
	Password string `json:"password"`

- RewardPayload
    Rollno int64 `json:"rollno,string"`
	Coins  int64 `json:"coins,string"`

- TransferPayload
    SenderRollno   int64 `json:"sender,string"`
	ReceiverRollno int64 `json:"receiver,string"`
	Coins          int64 `json:"coins,string"`
