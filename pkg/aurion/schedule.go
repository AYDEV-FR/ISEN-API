package aurion

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type FormId string

type ScrapScheduleOption struct {
	Start string
	End   string
	// Can be month, agendaWeek or agendaDay
	View     string
	Date     string
	Week     string
	Location string
}

func getFormId(document io.Reader) (FormId, error) {
	doc, err := goquery.NewDocumentFromReader(document)
	if err != nil {
		return "", err
	}
	formId := doc.Find("div[class=schedule]").First().AttrOr("id", "")
	return FormId(formId), nil
}

func CalendarPageOption(scheduleOptions ScrapScheduleOption) ScrapScheduleOption {
	var now time.Time
	var start time.Time
	var end time.Time
	var week int

	localLoc, _ := time.LoadLocation("Europe/Paris")
	if scheduleOptions.Date == "" {
		now = time.Now().In(localLoc)
	} else {
		now, _ = time.Parse("02/01/2006", scheduleOptions.Date)
	}

	if scheduleOptions.Week == "" {
		week = now.YearDay()/7 + 1
		if now.Weekday() == time.Sunday {
			week = week + 1
		}
		scheduleOptions.Week = fmt.Sprintf("%02d-%04d", week, now.Year())
	}
	fmt.Println(scheduleOptions.Week)

	if scheduleOptions.View == "" {
		scheduleOptions.View = "agendaWeek"
	}
	switch scheduleOptions.View {
	case "agendaWeek":
		if now.Weekday() == time.Sunday {
			start = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, localLoc)
			end = time.Date(now.Year(), now.Month(), now.Day()+7, 0, 0, 0, 0, localLoc)
		} else {
			start = time.Date(now.Year(), now.Month(), now.Day()-(int(now.Weekday())-1), 0, 0, 0, 0, localLoc)
			end = time.Date(now.Year(), now.Month(), now.Day()+(7-int(now.Weekday())), 0, 0, 0, 0, localLoc)
		}
	case "agendaDay":
		if now.Weekday() == time.Sunday {
			start = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, localLoc)
			end = time.Date(now.Year(), now.Month(), now.Day()+1, 23, 59, 59, 0, localLoc)
		} else {
			start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, localLoc)
			end = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, localLoc)
		}
	}
	//TODO: month view

	return ScrapScheduleOption{
		Start:    strconv.FormatInt(start.UnixMilli(), 10),
		End:      strconv.FormatInt(end.UnixMilli(), 10),
		View:     scheduleOptions.View,
		Date:     scheduleOptions.Date,
		Week:     scheduleOptions.Week,
		Location: scheduleOptions.Location,
	}
}

func CalendarPage(formId FormId, options ScrapScheduleOption) ScrapTableOption {

	return ScrapTableOption{
		Url: "https://ent-toulon.isen.fr/faces/Planning.xhtml",
		FormOption: url.Values{
			"javax.faces.partial.ajax":    {"true"},
			"javax.faces.source":          {string(formId)},
			"javax.faces.partial.execute": {string(formId)},
			"javax.faces.partial.render":  {string(formId)},
			"form":                        {"form"},
			string(formId):                {string(formId)},

			string(formId + "_start"): {options.Start},
			string(formId + "_end"):   {options.End},
			string(formId + "_view"):  {options.View},

			"form:date_input": {options.Date},
			"form:week":       {options.Week},
		},
	}
}

func ScrapSchedule(token Token, scheduleOptions ScrapScheduleOption, currentPage []byte) (string, error) {

	// Set client
	client := &http.Client{}

	// Get view state from currentPage
	reader := bytes.NewReader(currentPage)
	viewState, err := getViewState(reader)
	if err != nil {
		return "", err
	}
	// Get form id from currentPage
	reader = bytes.NewReader(currentPage)
	formId, err := getFormId(reader)
	if err != nil {
		return "", err
	}

	pageOptions := CalendarPage(formId, CalendarPageOption(scheduleOptions))

	formData := pageOptions.FormOption
	formData.Add("javax.faces.ViewState", string(viewState))

	req, err := http.NewRequest("POST", pageOptions.Url, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Faces-Request", "partial/ajax")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%v", token))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert partial response to HTML compatible array
	var partialResponse PartialResponse
	xml.Unmarshal(content, &partialResponse)

	return convertPartialResponseToJson(partialResponse, formId), nil
}

func convertPartialResponseToJson(partialResponse PartialResponse, formId FormId) string {
	var json string

	for _, update := range partialResponse.Changes.Update {
		if update.ID == string(formId) {
			// Convert string data to parsable json data
			json = update.Text[len("{\"events\" : ") : len(update.Text)-1]
		}
	}
	return json
}
