package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"paddletraffic/internal/model"
	"paddletraffic/internal/service"
)

type LocationController struct {
	service *service.LocationService
}

func NewLocationController(service *service.LocationService) *LocationController {
	return &LocationController{service: service}
}

func (c *LocationController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/locations", c.handleLocations)
	mux.HandleFunc("/locations/", c.handleLocation)
}

func (c *LocationController) handleLocations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.GetAll(w, r)
	case http.MethodPost:
		c.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *LocationController) handleLocation(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path
	path := strings.TrimPrefix(r.URL.Path, "/locations/")
	if path == "" {
		c.handleLocations(w, r)
		return
	}

	id, err := strconv.ParseUint(path, 10, 32)
	if err != nil {
		http.Error(w, "Invalid location ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		c.GetByID(w, r, uint(id))
	case http.MethodPut:
		c.Update(w, r, uint(id))
	case http.MethodDelete:
		c.Delete(w, r, uint(id))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *LocationController) Create(w http.ResponseWriter, r *http.Request) {
	var location model.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.Create(&location); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(location)
}

func (c *LocationController) GetByID(w http.ResponseWriter, r *http.Request, id uint) {
	location, err := c.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

func (c *LocationController) GetAll(w http.ResponseWriter, r *http.Request) {
	locations, err := c.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

func (c *LocationController) Update(w http.ResponseWriter, r *http.Request, id uint) {
	var location model.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	location.ID = id
	if err := c.service.Update(&location); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

func (c *LocationController) Delete(w http.ResponseWriter, r *http.Request, id uint) {
	if err := c.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
