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
		Hours:    "08:00 - 10:00",
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
		Code:     "21_HOGWAR_Defense_Against_the_Dark_Arts",
		Note:     "18",
		Comments: "Your Stunning Spell could be better",
		Teachers: []string{"Quirinus Quirrell", "Dolores Umbridge", "Remus Lupin"},
	},
}

var fakePersonalInformations = isen.PersonalInformations{
	Title:       "Mister",
	Name:        "Ronald",
	FirstName:   "WEASLEY",
	MiddleNames: "Bilius",
	Sex:         "M",
	BirthDate:   "01/03/1980",
	BirthPlace:  "The Burrow",
	Nationality: "England",
	PersonalAddress: isen.Address{
		Title:              "Mr.",
		Name:               "Ronald Weasley",
		StreetAddressLine1: "The Burrow",
		StreetAddressLine2: "Devon",
		StreetAddressLine3: "-",
		PostalCodeCity:     "EX11 Ottery St. Cactchpole",
		Country:            "United Kingdom",
	},
	ParentsAddress: isen.Address{
		Title:              "Mr. & Ms.",
		Name:               "Weasley",
		StreetAddressLine1: "The Burrow",
		StreetAddressLine2: "Devon",
		StreetAddressLine3: "-",
		PostalCodeCity:     "EX11 Ottery St. Cactchpole",
		Country:            "United Kingdom",
	},
	ReportSendingAddress: isen.Address{
		Title:              "Mr. & Ms.",
		Name:               "Weasley",
		StreetAddressLine1: "The Burrow",
		StreetAddressLine2: "Devon",
		StreetAddressLine3: "-",
		PostalCodeCity:     "EX11 Ottery St. Cactchpole",
		Country:            "United Kingdom",
	},
	BillingAddress: isen.Address{
		Title:              "Mr. & Ms.",
		Name:               "Weasley",
		StreetAddressLine1: "The Burrow",
		StreetAddressLine2: "Devon",
		StreetAddressLine3: "-",
		PostalCodeCity:     "EX11 Ottery St. Cactchpole",
		Country:            "United Kingdom",
	},
	PersonalPhone:      "+447769578549",
	ParentsPhone:       "+447464751722",
	ReportSendingEmail: "arthur.weasley@mom.gouv.uk",
	ParentsEmail:       "arthur.weasley@mom.gouv.uk",
	PersonalEmail:      "ronald.weasley@poudlard.uk",
	LastCertificate:    "Baccalauréat",
	Bac: isen.BacType{
		Academy:    "London",
		Year:       "1990",
		Type:       "L",
		Note:       "11.54",
		Merit:      "-",
		School:     "City of London School",
		SchoolCode: "FDS78DS5",
	},
	Father: isen.Parent{
		Name:       "Weasley",
		FirstName:  "Arthur",
		Occupation: "Government Officer",
	},
	Mother: isen.Parent{
		Name:       "Weasley",
		FirstName:  "Molly",
		Occupation: "Parenting",
	},
	HaveAcknowledgeGlobalRules:        func() *bool { b := true; return &b }(),
	HaveAcknowledgeStudentLifeCharter: func() *bool { b := true; return &b }(),
	CanIsenUsePersonalImage:           func() *bool { b := true; return &b }(),
}

var fakeNotesClass = []isen.NotationClass{
	{
		Code:         "21_HOGWART_S3_HERBOLOGY",
		Name:         "Semester 3 - Herbology - Examination",
		NotePersonal: "19",
		NoteAverage:  "14.21",
		NoteMin:      "9.64",
		NoteMax:      "19",
		Presence:     "62",
	},
	{
		Code:         "21_HOGWAR_Defense_Against_the_Dark_Arts",
		Name:         "Semester 3 - Defense Against the Dark Arts - Examination",
		NotePersonal: "18",
		NoteAverage:  "10.2",
		NoteMin:      "0",
		NoteMax:      "18.1",
		Presence:     "76",
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

var fakeEvent = isen.ScheduleEventDetails{
	Id:          "1",
	Start:       "Du mercredi 5 décembre 2001 à 08:00",
	End:         "Au mercredi 5 décembre 2001 à 10:00",
	Status:      "REALISE",
	Subject:     "Herbology",
	Type:        "Travaux dirigés",
	Description: "",
	IsPaper:     func() *bool { b := false; return &b }(),
	Rooms:       []string{"Greenhouse"},
	Teachers:    []string{"Sprout Professor"},
	Students:    []string{"GRANGER Hermione", "POTTER Harry", "WEASLEY Ronald"},
	Groups:      []string{"0102YEAR1"},
	CourseName:  "Herbology Class",
	Module:      "Herbology",
}
