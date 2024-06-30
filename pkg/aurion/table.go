package aurion

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type PartialResponse struct {
	XMLName xml.Name `xml:"partial-response"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr"`
	Changes struct {
		Text   string `xml:",chardata"`
		Update []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"update"`
	} `xml:"changes"`
}

type ScrapTableOption struct {
	Url        string
	FormOption url.Values
}

func ScrapTable(token Token, currentPage []byte, pageOptions ScrapTableOption) (string, error) {

	// Set client
	client := &http.Client{}

	// Get view state from currentPage
	reader := bytes.NewReader(currentPage)
	viewState, err := getViewState(reader)
	if err != nil {
		return "", err
	}

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

	content, _ := ioutil.ReadAll(resp.Body)
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

func convertPartialResponseToHTML(partialResponse PartialResponse) string {
	var html string
	for _, update := range partialResponse.Changes.Update {
		switch update.ID {
		case "form:dataTableFavori":
			html = fmt.Sprintf("<html><table>%s</table></html>", update.Text)
		case "form:tabPanelPrincipalFormulaireSupport":
			html = update.Text
		case "form:modaleDetail":
			html = fmt.Sprintf("<html>%s</html>", update.Text)
		}
	}
	return html
}
