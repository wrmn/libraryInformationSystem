# Library Information System
Library information system is golang based back-end for Padang city library and archives service.

>The initial goal of this project is to fulfill the final project of the course, but it does not rule out the possibility for further development.

## Setup
copy `.env.example` as `.env` and fill the configuration  
### go initialization
```sh
go mod init
go mod tidy
go get -u
```
### running app
```sh
go run .
```

## TODO
- Create table on database base on user input/file
- Read data from database
- RestAPI Handler
- JWT
- CRUD


## License
[MIT LICENSE](https://en.wikipedia.org/wiki/MIT_License)

Copyright &copy; Aditya Warman