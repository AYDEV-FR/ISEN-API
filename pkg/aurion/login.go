package aurion

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Token string

type Login struct {
	Username string `json:"username" example:"ronald.weasley" extensions:"x-order=1"`
	Password string `json:"password" example:"i<3hermione" extensions:"x-order=2"`
}

// GetToken - Return token/cookie with username and password given
func GetToken(username string, password string, loginPage string) (Token, error) {
	data := url.Values{
		"username": {username},
		"password": {password},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("POST", loginPage, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusFound {
		return "", errors.New("login ou mot de passe invalide")
	}

	cookies := res.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "JSESSIONID" && cookie.Value != "" {
			return Token(cookie.Value), nil
		}
	}

	return "", errors.New("token not found")
}
