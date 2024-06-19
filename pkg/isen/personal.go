package isen

import "github.com/AYDEV-FR/ISEN-Api/pkg/aurion"

type PersonalInformations struct {
	Title string `json:"title,omitempty"`
}

func GetPersonalInformations(token aurion.Token) (PersonalInformations, error) {
	// var personalInformations PersonalInformations = PersonalInformations{}

	page, err := aurion.MenuNavigateTo(token, SelfInfoMenuId, MainMenuPage)
	if err != nil {
		return PersonalInformations{}, err
	}

	htmlTable, err := aurion.ScrapTable(token, page, PersonalInformationsPage())
	if err != nil {
		return PersonalInformations{}, err
	}
	print(htmlTable)
	return PersonalInformations{}, err
}
