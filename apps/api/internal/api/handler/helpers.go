package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"paddletraffic/internal/dto"
)

func ParsePaginationParams(r *http.Request) dto.PaginationParams {
	page := dto.DefaultPage
	pageSize := dto.DefaultPageSize

	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil {
			page = p
		}
	}

	if pageSizeParam := r.URL.Query().Get("pageSize"); pageSizeParam != "" {
		if ps, err := strconv.Atoi(pageSizeParam); err == nil {
			pageSize = ps
		}
	}

	return dto.NewPaginationParams(page, pageSize)
}

func DecodeJSON[T any](r *http.Request) (T, error) {
	var v T
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}
