package isen

import (
	"encoding/json"
	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
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
