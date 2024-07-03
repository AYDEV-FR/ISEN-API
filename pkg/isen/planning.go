package isen

import (
	"encoding/json"
	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ScheduleEvent struct {
	Id        string `json:"id,omitempty" example:"1" extensions:"x-order=1"`
	Title     string `json:"title,omitempty" example:"08:00 - 10:00 - Herbology Class - Professor - Sprout - GreenHouse - TD - (02h00) -  - 154632" extensions:"x-order=2"`
	Start     string `json:"start,omitempty" example:"2001-12-05T08:00:00+0100" extensions:"x-order=3"`
	End       string `json:"end,omitempty" example:"2001-12-05T10:00:00+0100" extensions:"x-order=4"`
	AllDay    bool   `json:"allDay,omitempty" example:"false" extensions:"x-order=5"`
	Editable  bool   `json:"editable,omitempty" example:"true" extensions:"x-order=6"`
	ClassName string `json:"className,omitempty" example:"TD" extensions:"x-order=7"`
}

type ScheduleEventDetails struct {
	Id          aurion.EventId `json:"id,omitempty" example:"1" extensions:"x-order=10"`
	Start       string         `json:"start,omitempty" example:"Du mercredi 5 décembre 2001 à 08:00" extensions:"x-order=11"`
	End         string         `json:"end,omitempty" example:"Au mercredi 5 décembre 2001 à 10:00" extensions:"x-order=12"`
	Status      string         `json:"status,omitempty" example:"REALISE" extensions:"x-order=13"`
	Subject     string         `json:"subject,omitempty" example:"Herbology" extensions:"x-order=14"`
	Type        string         `json:"type,omitempty" example:"Travaux dirigés" extensions:"x-order=15"`
	Description string         `json:"description,omitempty" example:"" extensions:"x-order=16"`
	IsPaper     *bool          `json:"isPaper,omitempty" example:"false" extensions:"x-order=17"`
	Rooms       []string       `json:"rooms,omitempty" example:"Greenhouse" extensions:"x-order=18"`
	Teachers    []string       `json:"teachers,omitempty" example:"Sprout Professor" extensions:"x-order=19"`
	Students    []string       `json:"students,omitempty" example:"GRANGER Hermione,POTTER Harry,WEASLEY Ronald" extensions:"x-order=20"`
	Groups      []string       `json:"groups,omitempty" example:"0102YEAR1" extensions:"x-order=21"`
	CourseName  string         `json:"courseName,omitempty" example:"Herbology Class" extensions:"x-order=22"`
	Module      string         `json:"module,omitempty" example:"Herbology" extensions:"x-order=23"`
}

func GetPersonalAgenda(token aurion.Token, scheduleOptions aurion.ScrapScheduleOption) ([]ScheduleEvent, error) {
	var planning []ScheduleEvent
	page, err := aurion.MenuNavigateTo(token, SelfAgendaMenuId, MainMenuPage)
	if err != nil {
		return nil, err
	}

	scheduleEventsJson, err := aurion.ScrapSchedule(token, scheduleOptions, page)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(scheduleEventsJson), &planning)
	if err != nil {
		return nil, err
	}

	return planning, err
}

func GetPersonalAgendaEvent(token aurion.Token, eventId aurion.EventId) (ScheduleEventDetails, error) {
	var eventDetails ScheduleEventDetails
	page, err := aurion.MenuNavigateTo(token, SelfAgendaMenuId, MainMenuPage)
	if err != nil {
		return ScheduleEventDetails{}, err
	}

	scheduleEventsHtml, err := aurion.ScrapScheduleEvent(token, eventId, page)
	if err != nil {
		return ScheduleEventDetails{}, err
	}

	reader := strings.NewReader(scheduleEventsHtml)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ScheduleEventDetails{}, err
	}

	doc.Find("div[class='ui-grid-row']").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			eventDetails.Status = s.Find("div").Last().Text()
		case 1:
			eventDetails.Subject = s.Find("div").Last().Text()
		case 2:
			eventDetails.Type = s.Find("div").Last().Text()
		case 3:
			eventDetails.Description = s.Find("div").Last().Text()
		case 4:
			isPaperText := s.Find("div").Last().Text()
			if isPaperText == "Non" {
				f := new(bool)
				*f = false
				eventDetails.IsPaper = f
			} else if isPaperText == "Oui" {
				t := new(bool)
				*t = true
				eventDetails.IsPaper = t
			}
		}
	})

	doc.Find("table").Each(func(tableIndex int, s *goquery.Selection) {
		s.Find("tr[role='row']").Each(func(rowIndex int, s *goquery.Selection) {
			s.Find("td[role='gridcell']").Each(func(gridIndex int, s *goquery.Selection) {
				switch tableIndex {
				case 0:
					if rowIndex == 0 {
						eventDetails.Start += s.Text() + " "
					} else if rowIndex == 1 {
						eventDetails.End += s.Text() + " "
					}
				case 1:
					if gridIndex == 1 {
						// here we retrieve only "Libellé", not "Code"
						eventDetails.Rooms = append(eventDetails.Rooms, s.Text())
					}
				case 2:
					if gridIndex == 0 {
						eventDetails.Teachers = append(eventDetails.Teachers, s.Text()+" ")
					} else if gridIndex == 1 {
						eventDetails.Teachers[rowIndex-1] += s.Text()
					}
				case 3:
					if gridIndex == 0 {
						eventDetails.Students = append(eventDetails.Students, s.Text()+" ")
					} else if gridIndex == 1 {
						eventDetails.Students[rowIndex-1] += s.Text()
					}
				case 4:
					eventDetails.Groups = append(eventDetails.Groups, s.Text())
				case 5:
					if gridIndex == 0 {
						eventDetails.CourseName = s.Text()
					} else if gridIndex == 1 {
						eventDetails.Module = s.Text()
					}
				}
			})
		})
	})
	eventDetails.Start = strings.TrimSpace(eventDetails.Start)
	eventDetails.End = strings.TrimSpace(eventDetails.End)
	return eventDetails, err
}
