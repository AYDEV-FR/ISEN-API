package isen

import (
	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Address struct {
	Title              string `json:"title,omitempty" example:"Mr." extensions:"x-order=1"`
	Name               string `json:"name,omitempty" example:"Ronald Weasley" extensions:"x-order=2"`
	StreetAddressLine1 string `json:"streetAddressLine1,omitempty" example:"The Burrow" extensions:"x-order=3"`
	StreetAddressLine2 string `json:"streetAddressLine2,omitempty" example:"Devon" extensions:"x-order=4"`
	StreetAddressLine3 string `json:"streetAddressLine3,omitempty" example:"-" extensions:"x-order=5"`
	PostalCodeCity     string `json:"postalCodeCity,omitempty" example:"EX11 Ottery St. Cactchpole" extensions:"x-order=6"`
	Country            string `json:"country,omitempty" example:"United Kingdom" extensions:"x-order=7"`
}

type BacType struct {
	Academy    string `json:"academy,omitempty" example:"London" extensions:"x-order=1"`
	Year       string `json:"year,omitempty" example:"1990" extensions:"x-order=2"`
	Type       string `json:"type,omitempty" example:"L" extensions:"x-order=3"`
	Note       string `json:"note,omitempty" example:"11.54" extensions:"x-order=4"`
	Merit      string `json:"merit,omitempty" example:"-" extensions:"x-order=5"`
	School     string `json:"school,omitempty" example:"City of London School" extensions:"x-order=6"`
	SchoolCode string `json:"schoolCode,omitempty" example:"FDS78DS5" extensions:"x-order=7"`
}

type Parent struct {
	Name       string `json:"name,omitempty" example:"Weasley" extensions:"x-order=1"`
	FirstName  string `json:"firstName,omitempty" example:"Arthur" extensions:"x-order=2"`
	Occupation string `json:"occupation,omitempty" example:"Government Officer" extensions:"x-order=3"`
}

type PersonalInformations struct {
	Title                             string  `json:"title,omitempty" example:"Mister" extensions:"x-order=10"`
	Name                              string  `json:"name,omitempty" example:"Ronald" extensions:"x-order=11"`
	FirstName                         string  `json:"firstName,omitempty" example:"WEASLEY" extensions:"x-order=12"`
	MiddleNames                       string  `json:"middleNames,omitempty" example:"Bilius" extensions:"x-order=13"`
	Sex                               string  `json:"sex,omitempty" example:"M" extensions:"x-order=14"`
	BirthDate                         string  `json:"birthDate,omitempty" example:"01/03/1980" extensions:"x-order=15"`
	BirthPlace                        string  `json:"birthPlace,omitempty" example:"The Burrow" extensions:"x-order=16"`
	Nationality                       string  `json:"nationality,omitempty" example:"England" extensions:"x-order=17"`
	PersonalAddress                   Address `json:"personalAddress,omitempty" extensions:"x-order=18"`
	ParentsAddress                    Address `json:"parentsAddress,omitempty" extensions:"x-order=19"`
	ReportSendingAddress              Address `json:"reportSendingAddress,omitempty" extensions:"x-order=20"`
	BillingAddress                    Address `json:"billingAddress,omitempty" extensions:"x-order=21"`
	PersonalPhone                     string  `json:"personalPhone,omitempty" example:"+447769578549" extensions:"x-order=22"`
	ParentsPhone                      string  `json:"parentsPhone,omitempty" example:"+447464751722" extensions:"x-order=23"`
	ReportSendingEmail                string  `json:"reportSendingEmail,omitempty" example:"arthur.weasley@mom.gouv.uk" extensions:"x-order=24"`
	ParentsEmail                      string  `json:"parentsEmail,omitempty" example:"arthur.weasley@mom.gouv.uk" extensions:"x-order=25"`
	PersonalEmail                     string  `json:"personalEmail,omitempty" example:"ronald.weasley@poudlard.uk" extensions:"x-order=26"`
	LastCertificate                   string  `json:"lastCertificate,omitempty" example:"Baccalauréat" extensions:"x-order=27"`
	Bac                               BacType `json:"bac,omitempty" extensions:"x-order=28"`
	Father                            Parent  `json:"father,omitempty" extensions:"x-order=29"`
	Mother                            Parent  `json:"mother,omitempty" extensions:"x-order=30"`
	HaveAcknowledgeGlobalRules        *bool   `json:"haveAcknowledgeGlobalRules,omitempty" example:"true" extensions:"x-order=31"`
	HaveAcknowledgeStudentLifeCharter *bool   `json:"haveAcknowledgeStudentLifeCharter,omitempty" example:"true" extensions:"x-order=32"`
	CanIsenUsePersonalImage           *bool   `json:"canIsenUsePersonalImage,omitempty" example:"true" extensions:"x-order=33"`
}

func FulfillAddress(address *Address, i int, content string, hasStreetLine3 bool) {
	if hasStreetLine3 {
		switch i {
		case 0:
			address.Title = content
		case 1:
			address.Name = content
		case 2:
			address.StreetAddressLine1 = content
		case 3:
			address.StreetAddressLine2 = content
		case 4:
			address.StreetAddressLine3 = content
		case 5:
			address.PostalCodeCity = content
		case 6:
			address.Country = content
		}
	} else {
		switch i {
		case 0:
			address.Title = content
		case 1:
			address.Name = content
		case 2:
			address.StreetAddressLine1 = content
		case 3:
			address.StreetAddressLine2 = content
		case 4:
			address.PostalCodeCity = content
		case 5:
			address.Country = content
		}
	}
}

func FulfillParent(parent *Parent, i int, content string) {
	switch i {
	case 0:
		parent.Name = content
	case 1:
		parent.FirstName = content
	case 2:
		parent.Occupation = content
	}
}

func GetPersonalInformations(token aurion.Token) (PersonalInformations, error) {
	var personalInformations = PersonalInformations{}
	personalInformations.PersonalAddress = Address{}
	personalInformations.ParentsAddress = Address{}
	personalInformations.ReportSendingAddress = Address{}
	personalInformations.BillingAddress = Address{}
	personalInformations.Bac = BacType{}
	personalInformations.Father = Parent{}
	personalInformations.Mother = Parent{}

	page, err := aurion.MenuNavigateTo(token, SelfInfoMenuId, MainMenuPage)
	if err != nil {
		return PersonalInformations{}, err
	}

	htmlForm, err := aurion.ScrapTable(token, page, PersonalInformationsPage())
	if err != nil {
		return PersonalInformations{}, err
	}

	reader := strings.NewReader(htmlForm)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return PersonalInformations{}, err
	}

	skippedStrings := 0
	doc.Find(".colonne2").Each(func(i int, s *goquery.Selection) {
		content := strings.TrimSpace(s.Text())
		skipped := false

		if strings.Contains(content, "$(function()") || strings.Contains(s.Text(), "\n") {
			skippedStrings++
			skipped = true
		}

		if !skipped {
			i -= skippedStrings
			switch i {
			case 0:
				personalInformations.Title = content
			case 1:
				personalInformations.Name = content
			case 2:
				personalInformations.FirstName = content
			case 3:
				personalInformations.MiddleNames = content
			case 4:
				personalInformations.Sex = content
			case 5:
				personalInformations.BirthDate = content
			case 6:
				personalInformations.BirthPlace = content
			case 7:
				personalInformations.Nationality = content
			case 8, 9, 10, 11, 12, 13, 14:
				FulfillAddress(&(personalInformations.PersonalAddress), i-8, content, true)
			case 15, 16, 17, 18, 19, 20, 21:
				FulfillAddress(&(personalInformations.ParentsAddress), i-15, content, true)
			case 22, 23, 24, 25, 26, 27, 28:
				FulfillAddress(&(personalInformations.ReportSendingAddress), i-22, content, true)
			case 29, 30, 31, 32, 33, 34:
				FulfillAddress(&(personalInformations.BillingAddress), i-29, content, false)
			case 35:
				personalInformations.PersonalPhone = content
			case 36:
				personalInformations.ParentsPhone = content
			case 37:
				personalInformations.ReportSendingEmail = content
			case 38:
				personalInformations.ParentsEmail = content
			case 39:
				personalInformations.PersonalEmail = content
			case 40:
				personalInformations.LastCertificate = content
			case 41:
				personalInformations.Bac.Academy = content
			case 42:
				personalInformations.Bac.Year = content
			case 43:
				personalInformations.Bac.Type = content
			case 44:
				personalInformations.Bac.Note = content
			case 45:
				personalInformations.Bac.Merit = content
			case 46:
				personalInformations.Bac.School = content
			case 47:
				personalInformations.Bac.SchoolCode = content
			case 48, 49, 50:
				FulfillParent(&(personalInformations.Father), i-48, content)
			case 51, 52, 53:
				FulfillParent(&(personalInformations.Mother), i-51, content)
			case 54:
				condition := new(bool)
				*condition = content == "☑"
				personalInformations.HaveAcknowledgeGlobalRules = condition
			case 55:
				condition := new(bool)
				*condition = content == "☑"
				personalInformations.HaveAcknowledgeStudentLifeCharter = condition
			case 56:
				condition := new(bool)
				*condition = content == "Accepte"
				personalInformations.CanIsenUsePersonalImage = condition
			}
		}
	})
	return personalInformations, err
}
