package isen

import (
	"net/url"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
)

const (
	LoginPage    = "https://ent-toulon.isen.fr/login"
	MainMenuPage = "https://ent-toulon.isen.fr/faces/MainMenuPage.xhtml"
)

var (
	SelfInfoMenuId      aurion.MenuId = []string{"Mon compte", "Mes informations"}
	SelfAgendaMenuId    aurion.MenuId = []string{"Planning", "Mon planning"}
	NotationMenuId      aurion.MenuId = []string{"Scolarité", "Mes notes"}
	NotationClassMenuId aurion.MenuId = []string{"Scolarité", "Mes notes (classe)"}
	AbsenceMenuId       aurion.MenuId = []string{"Scolarité", "Mes absences"}
)

func NotationPage() aurion.ScrapTableOption {
	return aurion.ScrapTableOption{
		Url: "https://ent-toulon.isen.fr/faces/LearnerNotationListPage.xhtml",
		FormOption: url.Values{
			"javax.faces.partial.ajax":    {"true"},
			"javax.faces.source":          {"form:dataTableFavori"},
			"javax.faces.partial.execute": {"form:dataTableFavori"},
			"javax.faces.partial.render":  {"form:dataTableFavori"},
			"form:dataTableFavori":        {"form:dataTableFavori"},

			"form:dataTableFavori_sorting":       {"true"},
			"form:dataTableFavori_skipChildren":  {"true"},
			"form:dataTableFavori_encodeFeature": {"true"},
			"form:dataTableFavori_sortKey":       {"form:dataTableFavori:j_idt113"},
			"form:dataTableFavori_sortDir":       {"-1"},
			"form:table_first":                   {"200"},
			"form:table_rows":                    {"0"},
			"form":                               {"form"},
		},
	}
}

func NotationClassPage() aurion.ScrapTableOption {
	return aurion.ScrapTableOption{
		Url: "https://ent-toulon.isen.fr/faces/ChoixDonnee.xhtml",
		FormOption: url.Values{
			"javax.faces.partial.ajax":    {"true"},
			"javax.faces.source":          {"form:dataTableFavori"},
			"javax.faces.partial.execute": {"form:dataTableFavori"},
			"javax.faces.partial.render":  {"form:dataTableFavori"},
			"form:dataTableFavori":        {"form:dataTableFavori"},

			"form:dataTableFavori_sorting":       {"true"},
			"form:dataTableFavori_skipChildren":  {"true"},
			"form:dataTableFavori_encodeFeature": {"true"},
			"form:dataTableFavori_sortKey":       {"form:dataTableFavori:j_idt113"},
			"form:dataTableFavori_sortDir":       {"-1"},
			"form:table_first":                   {"200"},
			"form:table_rows":                    {"0"},
			"form":                               {"form"},
		},
	}
}

func AbsencePage() aurion.ScrapTableOption {
	return aurion.ScrapTableOption{
		Url: "https://ent-toulon.isen.fr/faces/MesAbsences.xhtml",
		FormOption: url.Values{
			"javax.faces.partial.ajax":    {"true"},
			"javax.faces.source":          {"form:dataTableFavori"},
			"javax.faces.partial.execute": {"form:dataTableFavori"},
			"javax.faces.partial.render":  {"form:dataTableFavori"},
			"form:dataTableFavori":        {"form:dataTableFavori"},

			"form:dataTableFavori_sorting":       {"true"},
			"form:dataTableFavori_skipChildren":  {"true"},
			"form:dataTableFavori_encodeFeature": {"true"},
			"form:dataTableFavori_sortKey":       {"form:dataTableFavori:j_idt153"},
			"form:dataTableFavori_sortDir":       {"-1"},
			"form:table_first":                   {"200"},
			"form:table_rows":                    {"0"},
			"form":                               {"form"},
		},
	}
}

func PersonalInformationsPage() aurion.ScrapTableOption {
	return aurion.ScrapTableOption{
		Url: "https://ent-toulon.isen.fr/faces/TeacherPage.xhtml",
		FormOption: url.Values{
			"javax.faces.partial.ajax":                {"true"},
			"javax.faces.source":                      {"form:tabPanelPrincipalFormulaireSupport"},
			"javax.faces.partial.execute":             {"form:tabPanelPrincipalFormulaireSupport"},
			"javax.faces.partial.render":              {"form:tabPanelPrincipalFormulaireSupport"},
			"form:tabPanelPrincipalFormulaireSupport": {"form:tabPanelPrincipalFormulaireSupport"},

			"form:tabPanelPrincipalFormulaireSupport_activeIndex": {"0"},
			"form:tabPanelPrincipalFormulaireSupport_scrollState": {"0"},
			"form": {"form"},
		},
	}
}

//2_3_3_0_0
