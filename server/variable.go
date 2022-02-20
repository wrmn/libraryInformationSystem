package server

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type response struct {
	Status  string
	Message interface{}
	Code    int
}

var key []byte

type credential struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	IsEmployee bool   `json:"isEmployee"`
	Division   int    `json:division,omitempty`
	jwt.RegisteredClaims
}

type tokenCred struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type responseParam struct {
	W      http.ResponseWriter
	Body   []byte
	Status int
}

type userIn struct {
	Username string
	Password string
}

type accountRegis struct {
	AccountData accountData `json:"account"`
	RegisData   regisData   `json:"data"`
}

type accountData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type regisData struct {
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	PlaceOfBirth  string `json:"placeOfBirth"`
	DateOfBirth   string `json:"dateOfBirth"`
	Address1      string `json:"address1"`
	Address2      string `json:"address2"`
	Profession    int    `json:"profession"`
	Institution   string `json:"institution"`
	PhoneNo       string `json:"phoneNo"`
	IsWhatsapp    bool   `json:"isWhatsapp"`
	IdentityNo    string `json:"identityNo"`
	IdentityType  int    `json:"identityType"`
	IdentityFile  string `json:"identityFile"`
	PhotoFile     string `json:"photoFile"`
	AgreementFile string `json:"agreementFile"`
}

type book struct {
	Id             int
	RegistrationId int
	SerialNumber   int
	DdcNo          string
	DdcOrder       int
	Title          string
	Author         string
	Publisher      string
	Availability   bool
	Status         int
	Price          int
}
