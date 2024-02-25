package isen

import (
	"encoding/json"
	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ScheduleEvent struct {
	Id        string `json:"Id,omitempty"`
	Title     string `json:"Title,omitempty"`
	Start     string `json:"Start,omitempty"`
	End       string `json:"End,omitempty"`
	AllDay    bool   `json:"AllDay,omitempty"`
	Editable  bool   `json:"Editable,omitempty"`
	ClassName string `json:"ClassName,omitempty"`
}

type ScheduleEventDetails struct {
	Id          aurion.EventId `json:"id,omitempty"`
	Start       string         `json:"start,omitempty"`
	End         string         `json:"end,omitempty"`
	Status      string         `json:"status,omitempty"`
	Subject     string         `json:"subject,omitempty"`
	Type        string         `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	IsPaper     *bool          `json:"isPaper,omitempty"`
	Rooms       []string       `json:"rooms,omitempty"`
	Teachers    []string       `json:"teachers,omitempty"`
	Students    []string       `json:"students,omitempty"`
	Groups      []string       `json:"groups,omitempty"`
	CourseName  string         `json:"courseName,omitempty"`
	Module      string         `json:"module,omitempty"`
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
