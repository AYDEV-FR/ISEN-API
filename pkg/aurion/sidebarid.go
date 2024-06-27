package aurion

import (
	"encoding/xml"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SidebarId string
type WebscolaappId string

func updateSidebarSubmenu(token Token, viewState ViewState, wId WebscolaappId, mainMenuUrl string) (string, ViewState, error) {
	client := &http.Client{}

	// prepare form values to load every child of a submenu.
	// j_idt52 is a const that doesn't change
	updateFormValue := url.Values{
		"javax.faces.partial.ajax":       {"true"},
		"javax.faces.partial.execute":    {"form:j_idt52"},
		"javax.faces.partial.render":     {"form:sidebar"},
		"javax.faces.source":             {"form:j_idt52"},
		"form:j_idt52":                   {"form:j_idt52"},
		"form":                           {"form"},
		"webscolaapp.Sidebar.ID_SUBMENU": {string(wId)},
		"javax.faces.ViewState":          {string(viewState)},
	}

	// Post on homepage to update submenu
	req, err := http.NewRequest("POST", mainMenuUrl, strings.NewReader(updateFormValue.Encode()))
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Faces-Request", "partial/ajax")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%v", token))

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	// Convert partial response to HTML compatible array
	var partialResponse PartialResponse
	err = xml.Unmarshal(content, &partialResponse)
	if err != nil {
		return "", "", err
	}

	updatedSubmenu := ""
	for _, update := range partialResponse.Changes.Update {
		switch update.ID {
		case "form:sidebar":
			updatedSubmenu = fmt.Sprintf("<html>%s</html>", update.Text)
		case "j_id1:javax.faces.ViewState:0":
			viewState = ViewState(update.Text)
		}
	}

	return updatedSubmenu, viewState, nil
}

func findSubmenuIds(doc *goquery.Selection, selector string, menuId MenuId, submenuIndex int) (SidebarId, WebscolaappId) {
	submenuSidebarId := SidebarId("")
	submenuWebscolaappId := WebscolaappId("")
	// we use the children here, otherwise Find goes into the nested list elements, counting new loaded submenus
	doc.Find(selector).First().Children().Each(func(i int, s *goquery.Selection) {
		if s.Find(".ui-menuitem-text").First().Text() == menuId[submenuIndex] {
			classes, _ := s.Attr("class")
			classesSlice := strings.Split(classes, " ")
			// if a submenu is loaded, front will add this class at the end, making it simple to check if a submenu is fully loaded
			if classesSlice[len(classesSlice)-1] == "enfants-entierement-charges" {
				// if it is the case, we give the current submenu selection and call the same function recursively
				submenuSidebarId, submenuWebscolaappId = findSubmenuIds(s, ".ui-menu-child", menuId, submenuIndex+1)
				// we then add the current submenu id before the next submenu id.
				submenuSidebarId = SidebarId(fmt.Sprintf("%d_", i)) + submenuSidebarId
				return
			} else {
				for _, c := range classesSlice {
					if strings.Contains(c, "submenu_") {
						// if the current element is a submenu and not an endpoint, add its webscolaapp ID to load the submenu
						submenuWebscolaappId = WebscolaappId(c)
						break
					}
				}
				// eventually add the submenu sidebarid, this could be the endpoint
				submenuSidebarId = SidebarId(fmt.Sprintf("%d_", i))
			}
			return
		}
	})

	return submenuSidebarId, submenuWebscolaappId
}

// getSidebarId - Get sidebar ID to dynamically choose the menu ID
func getSidebarId(token Token, viewState ViewState, document io.Reader, menuId MenuId, mainMenuUrl string) (SidebarId, error) {
	doc, err := goquery.NewDocumentFromReader(document)
	if err != nil {
		return "", err
	}

	sidebarId, webscolaappId := findSubmenuIds(doc.Selection, "div[id='form:sidebar'] .ui-menu-list", menuId, 0)

	newViewState := viewState

	// as long as webscolaappId is not empty, this means a new submenu has to be loaded.
	// this while loop load every submenu needed and find the correct sidebarId
	for webscolaappId != "" {
		updatedSubmenu, viewState, err := updateSidebarSubmenu(token, newViewState, webscolaappId, mainMenuUrl)
		if err != nil {
			return "", err
		}

		// as updateSidebarSubmenu updates the MainMenuPage with its newly loaded submenu child, we need to read it again
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(updatedSubmenu))
		if err != nil {
			return "", err
		}

		// find the correct sidebarId with the newly loaded submenu child
		sidebarId, webscolaappId = findSubmenuIds(doc.Selection, "div[id='form:sidebar'] .ui-menu-list", menuId, 0)

		newViewState = viewState
	}

	// the sidebarId will always end with an additional _ so we remove it before returning the value
	return sidebarId[:len(sidebarId)-1], nil
}
