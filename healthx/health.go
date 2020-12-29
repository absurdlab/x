package healthx

import (
	"github.com/absurdlab/x/httpx"
	"go.uber.org/fx"
	"net/http"
	"sync"
)

// NewHandler is the constructor for handler.
func NewHandler(in struct {
	fx.In
	Resources []*Resource `group:"health_resource"`
}) *Handler {
	return &Handler{resources: in.Resources}
}

type Handler struct {
	resources []*Resource
}

func (h *Handler) Status(rw http.ResponseWriter, r *http.Request) {
	var status = new(struct {
		Status    string    `json:"status"`
		Resources []*Report `json:"resources,omitempty"`
	})
	status.Status = UP

	var wg sync.WaitGroup
	wg.Add(len(h.resources))

	rc := make(chan *Report, len(h.resources))
	for _, resource := range h.resources {
		res := resource
		go func() {
			defer wg.Done()
			doc := res.Report(r.Context())
			rc <- &doc
		}()
	}

	wg.Wait()
	close(rc)

	for doc := range rc {
		status.Resources = append(status.Resources, doc)
		if doc.Status != UP {
			status.Status = DOWN
		}
	}

	httpStatus := 200
	if status.Status != UP {
		httpStatus = 503
	}

	httpx.Render(rw,
		httpx.SetApplicationJSON(),
		httpx.SetNoCache(),
		httpx.SetStatus(httpStatus),
		httpx.SetJSONPayload(status),
	)
}

type Report struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
