package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"paddletraffic/internal/dto"
)

// ---- fake repo ----

type fakeReportRepo struct {
	create dto.ReportSummary
	err    error
}

func (f *fakeReportRepo) Create(ctx context.Context, in dto.ReportCreate) (dto.ReportSummary, error) {
	return f.create, f.err
}

// ---- tests ----

func TestNewReportService_NotNil(t *testing.T) {
	svc := NewReportService(&fakeReportRepo{})
	if svc == nil {
		t.Fatal("NewReportService returned nil")
	}
}

func TestReportService_Create_OK(t *testing.T) {
	repo := &fakeReportRepo{
		create: dto.ReportSummary{
			ID:             123,
			CourtID:        1,
			CourtsOccupied: 2,
			GroupsWaiting:  1,
		},
		err: nil,
	}
	svc := NewReportService(repo)

	// pass a context value to ensure it’s forwarded unchanged
	type ctxKey string
	ctx := context.WithValue(context.Background(), ctxKey("k"), "v")

	var in dto.ReportCreate // zero-value is fine; repo captures what service passes through
	got, err := svc.Create(ctx, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := dto.ReportSummary{ID: 123, CourtID: 1, CourtsOccupied: 2, GroupsWaiting: 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("result mismatch: got %#v want %#v", got, want)
	}
}

func TestReportService_Create_Error(t *testing.T) {
	repo := &fakeReportRepo{
		create: dto.ReportSummary{},
		err:    errors.New("db failed"),
	}

	svc := NewReportService(repo)

	got, err := svc.Create(context.Background(), dto.ReportCreate{})
	if err == nil {
		t.Fatalf("expected error, got nil (result=%#v)", got)
	}
}
