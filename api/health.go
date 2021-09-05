package api

import (
	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/sysinfo"
)

// Health check
// swagger:model HealthResponse
type healthResponse struct {
	// in: body
	// Required: true
	Status health.Status `json:"status" example:"pass"`
	// Required: true
	Version string `json:"version,omitempty" example:"1"`
	// Required: false
	ReleaseID string `json:"releaseId,omitempty" example:"1.14.2-SNAPSHOT"`
}

func (a *API) initialiseHealth() {
	h := health.New(
		// TODO read from version
		health.Health{Version: "1", ReleaseID: "1.0.0-SNAPSHOT"},
		sysinfo.Health(),
	)

	// swagger:operation GET /health health getHealth
	//
	// Gets health of the API
	//
	// Returns the status of the API using the configured health checks
	// ---
	// responses:
	//   '200':
	//     description: Health response
	//     schema:
	//       "$ref": "#/definitions/HealthResponse"
	a.Router.HandleFunc("/health", h.Handler)
	// a.Router.HandleFunc("/health", h.Handler).Methods(http.MethodGet)
}
