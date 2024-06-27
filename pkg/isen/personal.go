package isen

import (
	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Address struct {
	Title              string `json:"title,omitempty"`
	Name               string `json:"name,omitempty"`
	StreetAddressLine1 string `json:"streetAddressLine1,omitempty"`
	StreetAddressLine2 string `json:"streetAddressLine2,omitempty"`
	StreetAddressLine3 string `json:"streetAddressLine3,omitempty"`
	PostalCodeCity     string `json:"postalCodeCity,omitempty"`
	Country            string `json:"country,omitempty"`
}

type BacType struct {
	Academy    string `json:"academy,omitempty"`
	Year       string `json:"year,omitempty"`
	Type       string `json:"type,omitempty"`
	Note       string `json:"note,omitempty"`
	Merit      string `json:"merit,omitempty"`
	School     string `json:"school,omitempty"`
	SchoolCode string `json:"schoolCode,omitempty"`
}

type Parent struct {
	Name       string `json:"name,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	Occupation string `json:"occupation,omitempty"`
}

type PersonalInformations struct {
	Title                             string  `json:"title,omitempty"`
	Name                              string  `json:"name,omitempty"`
	FirstName                         string  `json:"firstName,omitempty"`
	MiddleNames                       string  `json:"middleNames,omitempty"`
	Sex                               string  `json:"sex,omitempty"`
	BirthDate                         string  `json:"birthDate,omitempty"`
	BirthPlace                        string  `json:"birthPlace,omitempty"`
	Nationality                       string  `json:"nationality,omitempty"`
	PersonalAddress                   Address `json:"personalAddress,omitempty"`
	ParentsAddress                    Address `json:"parentsAddress,omitempty"`
	ReportSendingAddress              Address `json:"reportSendingAddress,omitempty"`
	BillingAddress                    Address `json:"billingAddress,omitempty"`
	PersonalPhone                     string  `json:"personalPhone,omitempty"`
	ParentsPhone                      string  `json:"parentsPhone,omitempty"`
	ReportSendingEmail                string  `json:"reportSendingEmail,omitempty"`
	ParentsEmail                      string  `json:"parentsEmail,omitempty"`
	PersonalEmail                     string  `json:"personalEmail,omitempty"`
	LastCertificate                   string  `json:"lastCertificate,omitempty"`
	Bac                               BacType `json:"bac,omitempty"`
	Father                            Parent  `json:"father,omitempty"`
	Mother                            Parent  `json:"mother,omitempty"`
	HaveAcknowledgeGlobalRules        *bool   `json:"haveAcknowledgeGlobalRules,omitempty"`
	HaveAcknowledgeStudentLifeCharter *bool   `json:"haveAcknowledgeStudentLifeCharter,omitempty"`
	CanIsenUsePersonalImage           *bool   `json:"canIsenUsePersonalImage,omitempty"`
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
