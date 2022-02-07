package main

import (
	"database/sql"
)

var (
	dbAction *string
	// sysAction  *string
	tableName  *string
	allCommand *bool
	task       string
	dmy        string = "2006-01-02 "
	dmyhms     string = "2006-01-02 15:04:05"
	gender            = []string{"M", "F"}
)

type QueryParam struct {
	Db        *sql.DB
	TableName string
}

type SearchParam struct {
	Column string
	Value  interface{}
}

type DbConfig struct {
	Connection string `json:"connection"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Database   string `json:"database"`
}

type AssetRecord struct {
	Id                 int
	AdminId            int
	RegistrationNumber string
	RegistrationDate   string
	Source             string
	CreatedAt          string
	UpdatedAt          string
}
type Book struct {
	Id             int
	RegistrationId int
	SerialNumber   int
	Ddc            string
	DdcOrder       int
	Title          string
	Author         string
	Publisher      string
	Availability   bool
	Price          int
	CreatedAt      string
	UpdatedAt      string
}
type Borrow struct {
	Id            int
	AdminId       int
	MemberId      int
	BookId        int
	FineStatus    bool
	DateOfBorrow  string
	DateOfReturn  string
	DateOfPayment string
}
type Ddc struct {
	Ddc   string
	Group string
	Name  string
}
type Employee struct {
	Id             int
	EmployeeNumber string
	Name           string
	Gender         string
	PlaceOfBirth   string
	DateOfBirth    string
	Address1       string
	Address2       string
	Division       string
	Position       string
}
type Guest struct {
	Id          int
	Name        string
	Gender      string
	Address     string
	Profession  string
	Institution string
}
type Inventory struct {
	Id             int
	RegistrationId int
	SerialNumber   int
	Name           string
	Category       string
	Status         string
	Description    string
}
type Member struct {
	Id            int
	Name          string
	Gender        string
	PlaceOfBirth  string
	DateOfBirth   string
	Address1      string
	Address2      string
	Profession    string
	Institution   string
	PhoneNo       string
	IsWhatsapp    bool
	IdentityNo    string
	IdentityType  string
	IdentityFile  string
	PhotoFile     string
	AgreementFile string
	CreatedAt     string
	VerifiedAt    string
	UpdatedAt     string
}
type User struct {
	Id        int
	Username  string
	Password  string
	LastLogin string
	CreatedAt string
	UpdatedAt string
}
type Visitor struct {
	Id      int
	UserId  int
	GuestId int
	LoginAt string
	Method  string
	Purpose string
}
