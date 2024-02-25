package isen

import (
	"encoding/json"
	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
)

type ScheduleEvent struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Start     string `json:"start,omitempty"`
	End       string `json:"end,omitempty"`
	AllDay    bool   `json:"allDay,omitempty"`
	Editable  bool   `json:"editable,omitempty"`
	ClassName string `json:"className,omitempty"`
}

type ScheduleEventDetails struct {
	Id string `json:"Id,omitempty"`
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
