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
