package isen

import (
	"strings"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
)

type Absence struct {
	Date     string   `json:"date,omitempty"`
	Reason   string   `json:"reason,omitempty"`
	Duration string   `json:"duration,omitempty"`
	Hours    string   `json:"hours,omitempty"`
	Course   string   `json:"course,omitempty"`
	Teachers []string `json:"teachers,omitempty"`
	Subject  string   `json:"subject,omitempty"`
}

func GetAbsenceList(token aurion.Token) ([]Absence, error) {
	var absencesList []Absence = []Absence{}

	page, err := aurion.MenuNavigateTo(token, AbsenceMenuId, MainMenuPage)
	if err != nil {
		return nil, err
	}

	htmlTable, err := aurion.ScrapTable(token, page, AbsencePage())
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(htmlTable)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	doc.Find("tr[role='row']").Each(func(i int, s *goquery.Selection) {
		var absence Absence
		s.Find("td[role='gridcell']").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				absence.Date = s.Text()
			case 1:
				absence.Reason = s.Text()
			case 2:
				absence.Duration = s.Text()
			case 3:
				absence.Hours = s.Text()
			case 4:
				absence.Course = s.Text()
			case 5:
				absence.Teachers = strings.Split(s.Text(), " - ")
			case 6:
				absence.Subject = s.Text()
			}
		})
		absencesList = append(absencesList, absence)
	})

	return absencesList, err
}
