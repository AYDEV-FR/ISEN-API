package aurion

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MenuId []string

// MenuNavigateTo
func MenuNavigateTo(token Token, menuId MenuId, mainMenuUrl string) ([]byte, error) {
	// Set client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

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

	// clone body for two different functions
	bodyBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body1 := io.NopCloser(bytes.NewBuffer(bodyBuf))
	body2 := io.NopCloser(bytes.NewBuffer(bodyBuf))
	// Get viewstate to request destination page
	viewState, err := getViewState(body1)
	if err != nil {
		return nil, err
	}

	sidebarId, err := getSidebarId(token, viewState, body2, menuId, mainMenuUrl)
	if err != nil {
		return nil, err
	}

	// Prepare post on homepage to be redirect on good page
	formData := url.Values{
		"javax.faces.ViewState": {string(viewState)},
		"form:sidebar_menuid":   {string(sidebarId)},
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
