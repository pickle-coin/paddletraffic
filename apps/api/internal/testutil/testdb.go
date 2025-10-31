package testutil

import (
	"database/sql"
	"testing"
)

// StartPostgresAndMigrate is a placeholder for an integration-test database bootstrapper.
// Replace with Testcontainers + migrations when implementing integration tests.
func StartPostgresAndMigrate(t *testing.T) (*sql.DB, func()) {
	t.Helper()
	t.Skip("integration test DB not implemented yet")
	return nil, func() {}
}

// SeedCourts is a placeholder for inserting seed court data for repository tests.
func SeedCourts(t *testing.T, db *sql.DB, num int) []int64 {
	t.Helper()
	t.Skip("seeding not implemented yet")
	return nil
}
