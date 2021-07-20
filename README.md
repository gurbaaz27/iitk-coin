# IITK Coin
## SnT Project 2021, Programming Club 

This repository contains the code for the IITKCoin, a vision of a pseudo-currency for use in the IITK Campus.

 project done so far.

### Relevant Links

- [Midterm Evaluation presentation](https://docs.google.com/presentation/d/1kriN-7A3v1RlXUDL5NETX3roJKRMJInptkWofIxY8dg/edit?usp=sharing)
- [Midterm Documentation](https://docs.google.com/document/d/1bvOWH4k0U-l2pQ1jLWIDzOkJ2wbHNW4jJw7tMWkUV6o/edit?usp=sharing)

## Table Of Content
- [Development Environment](#development-environment)
- [Directory Structure](#directory-structure)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Models](#models)

## Development Environment

```bash
- go version: go1.16.4 linux/amd64    # https://golang.org/dl/
- system: 5.4.72-microsoft-standard-WSL2 x86_64    # https://docs.microsoft.com/en-us/windows/wsl/install-win10
- text editor: VSCode    	# https://code.visualstudio.com/download
- terminal: Zsh     		# https://ohmyz.sh/
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
#, or build the program and run the executable
go build
./iitk-coin
```

Output should look like

```
2021/06/20 23:59:40 User Database opened and table created (if not existed) successfully!
2021/06/20 23:59:40 Transaction Database opened and table created (if not existed) successfully!
2021/06/20 23:59:40 Wallet Database opened and table created (if not existed) successfully!
2021/06/20 23:59:40 Serving at 8080
```

## Endpoints
POST requests take place via `JSON` requests. A typical usage would look like

```bash
curl -d '<json-request>' -H 'Content-Type: application/json' http://localhost:8080/<endpoint>
```

- `/login` : `POST`
```json
{"name":"<name>", "rollno":"<rollno>", "password":"<password>"}
```

- `/signup` : `POST`
```json
{"rollno":"<rollno>", "password":"<password>"}
```

- `/reward` : `POST`
```json
{"rollno":"<rollno>", "coins":"<coins>"}
```

- `/transfer` : `POST`
```json
{"sender":"<senderRollNo>", "receiver":"<receiverRollNo>", "coins":"<coins>"}
```

GET requests:

- `/secretpage` : `GET`
```bash
curl http://localhost:8080/secretpage
```

- `/balance` : `GET`
```bash
curl http://localhost:8080/balance?rollno=<rollno>
```

## Models

-  User
```go
	Name     string `json:"name"`
	Rollno   string  `json:"rollno"`
	Password string `json:"password"`
```

- RewardPayload
```go
	Rollno string `json:"rollno"`
	Coins  int64 `json:"coins,string"`
```

- TransferPayload
```go
	SenderRollno   string `json:"sender"`
	ReceiverRollno string `json:"receiver"`
	Coins          int64 `json:"coins,string"`
```
