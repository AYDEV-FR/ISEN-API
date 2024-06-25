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

// AgendaGet - Returns a list of all user's courses between start and end timestamps.
// start and end must be milliseconds UNIX timestamps. They are not mandatory and have defaults to the first and last day of the week.
func AgendaGet(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token header"})
		return
	}

	if token == "FAKETOKEN" {
		c.JSON(http.StatusOK, fakeAgenda)
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

// EventAgendaGet -
func EventAgendaGet(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token header"})
		return
	}

	eventId := c.Param("eventId")

	if token == "FAKETOKEN" {
		if eventId == "1" {
			c.JSON(http.StatusOK, fakeEvent)
		} else {
			c.JSON(http.StatusOK, "{}")
		}

		return
	}

	event, err := isen.GetPersonalAgendaEvent(aurion.Token(token), aurion.EventId(eventId))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
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
// NotationsClassGet - Returns a list of all user's class notes with min, average and max note
func NotationsClassGet(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token header"})
		return
	}

	if token == "FAKETOKEN" {
		c.JSON(http.StatusOK, fakeNotesClass)
		return
	}

	notationClass, err := isen.GetNotationClassList(aurion.Token(token))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notationClass)
}

// PersonalInformationsGet - Returns user's personal informations
func PersonalInformationsGet(c *gin.Context) {
  token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token header"})
		return
	}

	if token == "FAKETOKEN" {
    c.JSON(http.StatusOK, fakePersonalInformations)
		return
	}

	infos, err := isen.GetPersonalInformations(aurion.Token(token))
  
  if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
  
  c.JSON(http.StatusOK, infos)
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
