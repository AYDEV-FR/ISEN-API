package isen

import (
	"strings"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
)

type Absence struct {
	Date     string   `json:"date,omitempty" example:"05/12/2001" extensions:"x-order=1"`
	Reason   string   `json:"reason,omitempty" example:"Unexcused absence" extensions:"x-order=2"`
	Duration string   `json:"duration,omitempty" example:"2:00" extensions:"x-order=3"`
	Hours    string   `json:"hours,omitempty" example:"08:00 - 10:00" extensions:"x-order=4"`
	Course   string   `json:"course,omitempty" example:"Potion class" extensions:"x-order=5"`
	Teachers []string `json:"teachers,omitempty" example:"Severus Snape,Horace Slughorn" extensions:"x-order=6"`
	Subject  string   `json:"subject,omitempty" example:"Love filter potion" extensions:"x-order=7"`
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
		if absence.Date != "" {
			absencesList = append(absencesList, absence)
		}
	})

	return absencesList, err
}
