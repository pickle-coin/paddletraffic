-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS report;
DROP TABLE IF EXISTS court_status;
DROP TABLE IF EXISTS court;
DROP TABLE IF EXISTS location;

-- Drop the trigger function
DROP FUNCTION IF EXISTS set_updated_at();
