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
type EventId string

type ScrapScheduleOption struct {
	Start string
	End   string
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

	localLoc, _ := time.LoadLocation("Europe/Paris")
	now = time.Now().In(localLoc)

	if now.Weekday() == time.Sunday {
		now = now.AddDate(0, 0, 1)
	}

	if scheduleOptions.Start == "" {
		start = time.Date(now.Year(), now.Month(), now.Day()-(int(now.Weekday())-1), 0, 0, 0, 0, localLoc)
	} else {
		startMilli, _ := strconv.ParseInt(scheduleOptions.Start, 10, 64)
		start = time.UnixMilli(startMilli)
	}
	if scheduleOptions.End == "" {
		end = time.Date(now.Year(), now.Month(), now.Day()+(7-int(now.Weekday())), 23, 59, 59, 0, localLoc)
	} else {
		endMilli, _ := strconv.ParseInt(scheduleOptions.End, 10, 64)
		end = time.UnixMilli(endMilli)
	}

	return ScrapScheduleOption{
		Start: strconv.FormatInt(start.UnixMilli(), 10),
		End:   strconv.FormatInt(end.UnixMilli(), 10),
	}
}

func CalendarPage(formId FormId, options ScrapScheduleOption) ScrapTableOption {

	return ScrapTableOption{
		Url: "https://ent.isen-mediterranee.fr/faces/Planning.xhtml",
		FormOption: url.Values{
			"javax.faces.partial.ajax":    {"true"},
			"javax.faces.source":          {string(formId)},
			"javax.faces.partial.execute": {string(formId)},
			"javax.faces.partial.render":  {string(formId)},
			"form":                        {"form"},
			string(formId):                {string(formId)},

			string(formId + "_start"): {options.Start},
			string(formId + "_end"):   {options.End},
		},
	}
}

func CalendarEventPage(formId FormId, eventId EventId) ScrapTableOption {
	return ScrapTableOption{
		Url: "https://ent.isen-mediterranee.fr/faces/Planning.xhtml",
		FormOption: url.Values{
			"javax.faces.partial.ajax":          {"true"},
			"javax.faces.source":                {string(formId)},
			"javax.faces.partial.execute":       {string(formId)},
			"javax.faces.partial.render":        {"form:modaleDetail"},
			"form":                              {"form"},
			"javax.faces.behavior.event":        {"eventSelect"},
			"javax.faces.partial.event":         {"eventSelect"},
			string(formId) + "_selectedEventId": {string(eventId)},
		},
	}
}

func ScrapSchedule(token Token, scheduleOptions ScrapScheduleOption, currentPage []byte) (string, error) {

	// Set client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

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
	err = xml.Unmarshal(content, &partialResponse)
	if err != nil {
		return "", err
	}

	return convertPartialResponseToJson(partialResponse, formId), nil
}

func ScrapScheduleEvent(token Token, eventId EventId, currentPage []byte) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

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

	pageOptions := CalendarEventPage(formId, eventId)

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

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert partial response to HTML compatible array
	var partialResponse PartialResponse
	err = xml.Unmarshal(content, &partialResponse)
	if err != nil {
		return "", err
	}

	return convertPartialResponseToHTML(partialResponse), nil
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
