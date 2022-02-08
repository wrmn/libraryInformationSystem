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

```
go run . <flag> <option>
```

## Command

### Database (-d)

Command for action on database
Action

- Create Specific table

```
 go run . -d create -t <table-name>
```

- Create all table available

```
 go run . -d create -a
```

- Seed Specific table

```
 go run . -d seed -t <table-name>
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

- Create table on database base on user input/file
- Read data from database
- RestAPI Handler
- JWT
- CRUD

## License

[MIT LICENSE](https://en.wikipedia.org/wiki/MIT_License)

Copyright &copy; Aditya Warman
