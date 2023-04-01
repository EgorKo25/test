package handler

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Handler struct {
	X1     float64 `json:"x1,omitempty"`
	X2     float64 `json:"x2,omitempty"`
	Result string  `json:"result,omitempty"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) LineS(w http.ResponseWriter, r *http.Request) {

	h.Clear()

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

	err = h.solve(float64(a), float64(b), float64(c))
	if err != nil {
		h.Result = "D less zero"

		body, _ := json.Marshal(h)

		_, _ = w.Write(body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		return
	}

	log.Println(h)
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

func (h *Handler) Clear() {
	h.X1, h.X2 = 0, 0
	h.Result = ""
}

func (h *Handler) solve(a, b, c float64) error {
	des := b*b - 4*c*a

	if des < 0 {
		return errors.New("D less 0")
	}

	des = math.Sqrt(des)
	log.Println(des)

	h.X1 = (-b + des) / 2 * a
	h.X2 = (-b - des) / 2 * a

	return nil
}
