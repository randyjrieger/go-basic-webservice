package controllers

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"webservice/models"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

var (
	telemetryClient               appinsights.TelemetryClient
	AppInsightsInstrumentationKey = "1eae044e-9de9-4c76-b775-e6e8e9414241"
)

func init() {
	flag.StringVar(&AppInsightsInstrumentationKey, "instrumentationKey", AppInsightsInstrumentationKey, "set instrumentation key from azure portal")
	telemetryClient = appinsights.NewTelemetryClient(AppInsightsInstrumentationKey)
	/*Set role instance name globally -- this is usually the name of the service submitting the telemetry*/
	telemetryClient.Context().Tags.Cloud().SetRole("hello-world")
	/*turn on diagnostics to help troubleshoot problems with telemetry submission. */
	appinsights.NewDiagnosticsMessageListener(func(msg string) error {
		log.Printf("[%s] %s\n", time.Now().Format(time.UnixDate), msg)
		return nil
	})
}

// specify the type to bind the function to
func (uc *userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	telemetryClient.TrackEvent("List of clients requested.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetUsers())
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	telemetryClient.TrackEvent("Adding new client: " + u.FirstName + " " + u.LastName)
	u, err = models.AddUser(u)
}

// update
func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	telemetryClient.TrackEvent("Client: " + u.FirstName + " " + u.LastName + " is being updated.")
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	telemetryClient.TrackEvent("Client with the following Id is being removed: " + string(id))
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// constructor function - naming convention -> new + Type of object
// create new userController, return address to it
// returns pointer to a userController object

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

// no named variable, just constructing the object and using the pointer to it - allowed to do with structs
// scope - creating userController and returning the address - which should be out of scope if access external to package
// -- Go will see if we are returning address of a local variable and so it will promote it to the level of where this package and the caller can see it
