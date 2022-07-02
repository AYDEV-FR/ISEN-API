package aurion

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type MenuId string

// MenuNavigateTo
func MenuNavigateTo(token Token, menu_id MenuId, mainMenuUrl string) ([]byte, error) {
	// Set client
	client := &http.Client{}

	// Get homepage
	req, err := http.NewRequest("GET", mainMenuUrl, nil)
	req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%v", token))

	if err != nil {
		return nil, err
	}

	// Do request homepage
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Test if everithing is ok
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Get viewstate to request destination page
	viewState, err := getViewState(resp.Body)
	if err != nil {
		return nil, err
	}

	// Prepare post on homepage to be redirect on good page
	formData := url.Values{
		"javax.faces.ViewState": {string(viewState)},
		"form:sidebar_menuid":   {string(menu_id)},
		"form:sidebar":          {"form:sidebar"},
		"form":                  {"form"},
	}

	// Post
	req, err = http.NewRequest("POST", mainMenuUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%v", token))

	// Do Post
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// func GetMenuId(pageContent []byte, menuPath string) (string, error) {
// 	// Get only #form:sidebar
// 	reader := bytes.NewReader(pageContent)
// 	doc, err := goquery.NewDocumentFromReader(reader)
// 	if err != nil {
// 		return "", err
// 	}

// 	doc.Find("#form:sidebar>ui-menu-list").Each(func(i int, s *goquery.Selection) {
// 		fmt.Printf("%s", s.Text())
// 		s.Find("ul[role='gridcell']").Each(func(i int, s *goquery.Selection) {

// 		})
// 	})

// 	return "", nil
// }
