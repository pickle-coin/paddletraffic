package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/service"

	"github.com/go-chi/chi/v5"
)

// --- fake service ---

type fakeReportService struct { // implements Report Service
	create dto.ReportSummary
	err    error
}

func (f *fakeReportService) Create(ctx context.Context, in dto.ReportCreate) (dto.ReportSummary, error) {
	return f.create, f.err
}

func newReportService() service.ReportService {
	return &fakeReportService{
		create: dto.ReportSummary{
			ID:             1,
			CourtID:        1,
			CourtsOccupied: 1,
			GroupsWaiting:  1,
		},
		err: nil,
	}
}

// --- helpers ---

func testRouter(h *ReportHandler) http.Handler {
	r := chi.NewRouter()
	h.RegisterRoutes(r)
	return r
}

// --- tests ---

func TestCreateReport_OK(t *testing.T) {
	// service returns success
	reportService := newReportService()
	h := NewReportHandler(reportService)
	r := testRouter(h)

	body := []byte(`{"court_id":1, "courts_occupied":2,"groups_waiting":1}`)

	req := httptest.NewRequest("POST", "/v1/reports/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("got %d, want %d, body=%s", rr.Code, http.StatusCreated, rr.Body.String())
	}
}

func TestCreateReport_BadRequestFormat(t *testing.T) {
	// service won't be called, body is invalid
	svc := &fakeReportService{}
	h := NewReportHandler(svc)
	r := testRouter(h)

	req := httptest.NewRequest("POST", "/v1/reports/1", bytes.NewReader([]byte(`{not valid json`)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("got %d, want %d, body=%s", rr.Code, http.StatusBadRequest, rr.Body.String())
	}

}
func TestCreateReport_BadRequestJson(t *testing.T) {
	// service won't be called, body is invalid
	svc := &fakeReportService{}
	h := NewReportHandler(svc)
	r := testRouter(h)

	req := httptest.NewRequest("POST", "/v1/reports/1", bytes.NewReader([]byte(`{"not": "valid shape"}`)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("got %d, want %d, body=%s", rr.Code, http.StatusBadRequest, rr.Body.String())
	}
}

func TestCreateReport_ServiceError(t *testing.T) {
	svc := &fakeReportService{err: errors.New("db failed")}
	h := NewReportHandler(svc)
	r := testRouter(h)

	body := []byte(`{"court_id": 1, "courts_occupied":1,"groups_waiting":0}`)

	req := httptest.NewRequest("POST", "/v1/reports/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("got %d, want %d, body=%s", rr.Code, http.StatusInternalServerError, rr.Body.String())
	}
}

func TestCreateReport_MissingField(t *testing.T) {
	svc := newReportService()
	h := NewReportHandler(svc)
	r := testRouter(h)

	body := []byte(`{"courts_occupied":1}`)

	req := httptest.NewRequest("POST", "/v1/reports/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("got %d, want %d, body=%s", rr.Code, http.StatusBadRequest, rr.Body.String())
	}
}
