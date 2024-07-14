package api

import (
	"net/http"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/AYDEV-FR/ISEN-Api/pkg/isen"
	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	Error string `json:"error" example:"error message"`
}

// AbsencesGet gin Handler for /absences endpoint
//
//	@Summary		List user's absences
//	@Description	Get user's absences
//	@Produce		json
//	@Success		200	{array}		isen.Absence
//	@Failure		400	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/absences [get]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, absences)
}

// AgendaGet gin Handler for /agenda endpoint
//
//	@Summary		Get user's agenda
//	@Description	Returns a list of all user's courses between start and end timestamps.
//	@Produce		json
//	@Param			start	query		string	false	"UNIX Milliseconds when the events in the agenda will begin"
//	@Param			end		query		string	false	"UNIX Milliseconds when the events in the agenda will end"
//	@Success		200		{array}		isen.ScheduleEvent
//	@Failure		400		{object}	HTTPError
//	@Failure		500		{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/agenda [get]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agenda)
}

// EventAgendaGet gin Handler for /agenda/event/{eventId} endpoint
//
//	@Summary		Get informations from a specific event
//	@Description	Get informations of an agenda event from its ID
//	@Produce		json
//	@Param			eventId	path		string	true	"Agenda Event ID"
//	@Success		200		{object}	isen.ScheduleEventDetails
//	@Failure		400		{object}	HTTPError
//	@Failure		500		{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/agenda/event/{eventId} [get]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

// NotationsGet gin Handler for /notations endpoint
//
//	@Summary		List user's notations
//	@Description	Get user's notations
//	@Produce		json
//	@Success		200	{array}		isen.Notation
//	@Failure		400	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/notations [get]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notation)
}

// NotationsClassGet gin Handler for /notations/class endpoint
//
//	@Summary		List user's class notations
//	@Description	Get a list of all user's class notes with min, average and max note
//	@Produce		json
//	@Success		200	{array}		isen.Notation
//	@Failure		400	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/notations/class [get]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notationClass)
}

// PersonalInformationsGet gin Handler for /personal-informations endpoint
//
//	@Summary		List user's personal informations
//	@Description	Get user's personal informations
//	@Produce		json
//	@Success		200	{object}	isen.PersonalInformations
//	@Failure		400	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/personal-informations [get]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, infos)
}

// TokenPost gin Handler for /token endpoint
//
//	@Summary		Get ISEN Token
//	@Description	Get ISEN Token needed for every authenticated request
//	@Accept			json
//	@Produce		plain
//	@Param			account	body		aurion.Login	true	"Account credentials"
//	@Success		200		{string}	string
//	@Failure		400		{object}	HTTPError
//	@Failure		500		{object}	HTTPError
//	@Router			/token [post]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "text/plain", []byte(token))
}
