package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	store := persistence.NewInMemoryStore(time.Second)
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, cachePage(store, time.Minute, route.HandlerFunc))
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello ISEN!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/v1/",
		Index,
	},

	{
		"AbsencesGet",
		http.MethodGet,
		"/v1/absences",
		AbsencesGet,
	},

	{
		"AgendaGet",
		http.MethodGet,
		"/v1/agenda",
		AgendaGet,
	},

	{
		"AgendaEventGet",
		http.MethodGet,
		"/v1/agenda/event/:eventId",
		EventAgendaGet,
	},

	{
		"NotationsGet",
		http.MethodGet,
		"/v1/notations",
		NotationsGet,
	},

	{
		"TokenPost",
		http.MethodPost,
		"/v1/token",
		TokenPost,
	},
}
