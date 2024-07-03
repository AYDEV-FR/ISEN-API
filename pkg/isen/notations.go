package isen

import (
	"strings"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
)

type Notation struct {
	Date          string   `json:"date,omitempty" example:"05/12/2001" extensions:"x-order=1"`
	Code          string   `json:"code,omitempty" example:"21_HOGWAR_Defense_Against_the_Dark_Arts" extensions:"x-order=2"`
	Name          string   `json:"name,omitempty" example:"Semester 3 - Defense Against the Dark Arts - Examination" extensions:"x-order=3"`
	Note          string   `json:"note,omitempty" example:"18" extensions:"x-order=4"`
	AbsenceReason string   `json:"absenceReason,omitempty" example:"" extensions:"x-order=5"`
	Comments      string   `json:"comments,omitempty" example:"Your Stunning Spell could be better" extensions:"x-order=6"`
	Teachers      []string `json:"teachers,omitempty" example:"Quirinus Quirrell,Dolores Umbridge,Remus Lupin" extensions:"x-order=7"`
}

type NotationClass struct {
	Code         string `json:"name,omitempty" example:"21_HOGWAR_Defense_Against_the_Dark_Arts" extensions:"x-order=1"`
	Name         string `json:"code,omitempty" example:"Semester 3 - Defense Against the Dark Arts - Examination" extensions:"x-order=2"`
	NotePersonal string `json:"notePersonal,omitempty" example:"18" extensions:"x-order=3"`
	NoteAverage  string `json:"noteAverage,omitempty" example:"10.2" extensions:"x-order=4"`
	NoteMin      string `json:"noteMin,omitempty" example:"0" extensions:"x-order=5"`
	NoteMax      string `json:"noteMax,omitempty" example:"18.1" extensions:"x-order=6"`
	Presence     string `json:"presence,omitempty" example:"76" extensions:"x-order=7"`
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

func GetNotationClassList(token aurion.Token) ([]NotationClass, error) {
	var notationsClassList []NotationClass

	currentPage, err := aurion.MenuNavigateTo(token, NotationClassMenuId, MainMenuPage)
	if err != nil {
		return nil, err
	}

	htmlTable, err := aurion.ScrapTable(token, currentPage, NotationClassPage())
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(htmlTable)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	doc.Find("tr[role='row']").Each(func(i int, s *goquery.Selection) {
		var noteClass NotationClass
		s.Find("td[role='gridcell']").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				noteClass.Code = s.Text()
			case 1:
				noteClass.Name = s.Text()
			case 2:
				noteClass.NotePersonal = s.Text()
			case 3:
				noteClass.NoteAverage = s.Text()
			case 4:
				noteClass.NoteMin = s.Text()
			case 5:
				noteClass.NoteMax = s.Text()
			case 6:
				noteClass.Presence = s.Text()
			}
		})
		if noteClass.Name != "" {
			notationsClassList = append(notationsClassList, noteClass)
		}
	})

	return notationsClassList, err
}
