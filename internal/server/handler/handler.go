package handler

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Handler struct {
	X1 float64 `json:"x1"`
	X2 float64 `json:"x2"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) LineS(w http.ResponseWriter, r *http.Request) {

	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil || a == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := strconv.Atoi(r.URL.Query().Get("c"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.solve(a, b, c)

	body, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	return
}

func (h *Handler) solve(a, b, c int) {
	des := b*b + 4*c*a

	h.X1 = (math.Sqrt(float64(des)) - float64(b)) / 2 * float64(a)
	h.X2 = (math.Sqrt(float64(des)) - float64(b)) / 2 * float64(a)

}
