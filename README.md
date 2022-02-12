# Library Information System

Library information system is golang based back-end for Padang city library and archives service.

> The initial goal of this project is to fulfill the final project of the course, but it does not rule out the possibility for further development.

# Setup

copy `.env.example` as `.env` and fill the configuration

## Go initialization

```
go mod init
go mod tidy
go get -u
```

## Running app

without build

```
go run . <flag> <option>
```

building

```
go build librarySysfo
```

```
./librarySysfo <flag> <option>
```

## Command

### run (-r)

Command for action on database
Action

- Migrate

```sh
#migrate all
$ go run . -r migrate
#migrate specific table
$ go run . -r migrate -t <table-name>
```

- Seed

```sh
#seed all
$ go run . -r seed
#seed specific table
$ go run . -r seed -t <table-name>
```

- Serve

```sh
$ go run . -r serve
```

## Table Name

- assetRecord
- books
- borrows
- ddc
- employees
- guests
- inventory
- members
- users
- visitors

## ERD

![Erd](https://imgur.com/g1XyOPp.png)

## TODO

- RestAPI Handler
- CRUD

## License

[MIT LICENSE](https://en.wikipedia.org/wiki/MIT_License)

Copyright &copy; Aditya Warman
