package isen

import (
	"strings"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
)

type Notation struct {
	Date          string   `json:"date,omitempty"`
	Code          string   `json:"code,omitempty"`
	Name          string   `json:"name,omitempty"`
	Note          string   `json:"note,omitempty"`
	AbsenceReason string   `json:"absenceReason,omitempty"`
	Comments      string   `json:"comments,omitempty"`
	Teachers      []string `json:"teachers,omitempty"`
}

func GetNotationList(token aurion.Token) ([]Notation, error) {
	var notationsList []Notation = []Notation{}

	currentPage, err := aurion.MenuNavigateTo(token, NotationMenuId, MainMenuPage)
	if err != nil {
		return nil, err
	}

	htmlTable, err := aurion.ScrapTable(token, currentPage, NotationPage())
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(htmlTable)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	doc.Find("tr[role='row']").Each(func(i int, s *goquery.Selection) {
		var note Notation
		s.Find("td[role='gridcell']").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				note.Date = s.Text()
			case 1:
				note.Code = s.Text()
			case 2:
				note.Name = s.Text()
			case 3:
				note.Note = s.Text()
			case 4:
				note.AbsenceReason = s.Text()
			case 5:
				note.Comments = s.Text()
			case 6:
				note.Teachers = strings.Split(s.Text(), ",")
			}
		})
		notationsList = append(notationsList, note)
	})

	return notationsList, err
}
