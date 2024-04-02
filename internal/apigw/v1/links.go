package v1

import (
	"net/http"
)

func newLinksHandler(linksClient linksClient) *linksHandler {
	return &linksHandler{client: linksClient}
}

type linksHandler struct {
	client linksClient
}

func (h *linksHandler) GetLinks(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *linksHandler) PostLinks(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *linksHandler) DeleteLinksId(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

func (h Handler) GetLinksId(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *linksHandler) PutLinksId(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *linksHandler) GetLinksUserUserID(w http.ResponseWriter, r *http.Request, userID string) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}
