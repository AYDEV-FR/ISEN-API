package api

import (
	"net/http"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/AYDEV-FR/ISEN-Api/pkg/isen"
	"github.com/gin-gonic/gin"
)

// AbsencesGet - Return absence list
func AbsencesGet(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token header"})
		return
	}
	if token == "FAKETOKEN" {
		c.JSON(http.StatusOK, fakeAbs)
		return
	}

	absences, err := isen.GetAbsenceList(aurion.Token(token))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, absences)
}

// AgendaGet -
func AgendaGet(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token header"})
		return
	}

	if token == "FAKETOKEN" {
		c.JSON(http.StatusOK, "TODO")
		return
	}

	queryParams := c.Request.URL.Query()
	scheduleOptions := aurion.ScrapScheduleOption{
		Start: queryParams.Get("start"),
		End:   queryParams.Get("end"),
	}

	agenda, err := isen.GetPersonalAgenda(aurion.Token(token), scheduleOptions)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agenda)
}

// NotationsGet - Returns a list of all user's notes
func NotationsGet(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token header"})
		return
	}

	if token == "FAKETOKEN" {
		c.JSON(http.StatusOK, fakeNotes)
		return
	}

	notation, err := isen.GetNotationList(aurion.Token(token))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notation)
}

// TokenPost -
func TokenPost(c *gin.Context) {
	var loginCredentials aurion.Login
	err := c.BindJSON(&loginCredentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "syntax error, please see documentation /v1/"})
		return
	}

	if loginCredentials.Username == "ronald.weasley" && loginCredentials.Password == "i<3hermione" {
		c.Data(http.StatusOK, "text/plain", []byte("FAKETOKEN"))
		return
	}

	token, err := aurion.GetToken(loginCredentials.Username, loginCredentials.Password, isen.LoginPage)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "text/plain", []byte(token))
}
