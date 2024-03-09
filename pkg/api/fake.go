package api

import "github.com/AYDEV-FR/ISEN-Api/pkg/isen"

var fakeAbs = []isen.Absence{
	{
		Date:     "05/12/2001",
		Reason:   "Excused Absence",
		Duration: "4:00",
		Hours:    "08:00 - 12:00",
		Course:   "Herbology class",
		Teachers: []string{"Professor Sprout"},
		Subject:  "Herbology of Mandrake Root",
	}, {

		Date:     "05/12/2001",
		Reason:   "Unexcused absence",
		Duration: "2:00",
		Hours:    "08:00 - 12:00",
		Course:   "Potion class",
		Teachers: []string{"Severus Snape", "Horace Slughorn"},
		Subject:  "Love filter potion",
	},
}

var fakeNotes = []isen.Notation{
	{
		Date:     "05/12/2001",
		Code:     "21_HOGWART_S3_HERBOLOGY",
		Note:     "19",
		Comments: "Good job !",
		Teachers: []string{"Professor Sprout"},
	},
	{
		Date:     "05/12/2001",
		Code:     "21_HOGWAR_Defense_Against_the_Dark_Arts ",
		Note:     "18",
		Comments: "Your Stunning Spell could be better",
		Teachers: []string{"Quirinus Quirrell", "Dolores Umbridge", "Remus Lupin"},
	},
}

var fakeAgenda = []isen.ScheduleEvent{
	{
		Id:        "1",
		Title:     "08:00 - 10:00 - Herbology Class - Professor - Sprout - GreenHouse - TD - (02h00) -  - 154632",
		Start:     "2001-12-05T08:00:00+0100",
		End:       "2001-12-05T10:00:00+0100",
		Editable:  true,
		ClassName: "TD",
	},
	{
		Id:        "2",
		Title:     "10:00 - 12:00 - Potion Class - Severus - Snape - Alchemy Classroom - Magistral Course - (02h00) -  - 597864",
		Start:     "2001-12-05T10:00:00+0100",
		End:       "2001-12-05T12:00:00+0100",
		Editable:  true,
		ClassName: "CM",
	},
}
