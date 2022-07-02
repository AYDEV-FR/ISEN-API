package aurion

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type ViewState string

// getViewState - Get viewstate from http response
func getViewState(document io.Reader) (ViewState, error) {
	doc, err := goquery.NewDocumentFromReader(document)
	if err != nil {
		return "", err
	}
	viewState := doc.Find("input[name='javax.faces.ViewState']").First().AttrOr("value", "")
	return ViewState(viewState), nil
}
