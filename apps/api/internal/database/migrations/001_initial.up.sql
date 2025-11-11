-- Function to set the updated_at timestamp.
-- Can be used on any table that has an updated_at column.
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- ====================================================================================================
-- Location
-- ====================================================================================================

CREATE TABLE IF NOT EXISTS location (
  id            BIGSERIAL PRIMARY KEY,

  -- required
  address_line  TEXT            NOT NULL,
  country_code  TEXT            NOT NULL,
  timezone      TEXT            NOT NULL,
  lat           NUMERIC(9,6)    NOT NULL,
  lon           NUMERIC(9,6)    NOT NULL,

  -- optional
  region        TEXT,
  postal_code   TEXT,
  place_id      TEXT,

  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  -- constraints derived from schema
  CONSTRAINT country_code_len_chk   CHECK (char_length(country_code) = 2),
  CONSTRAINT lat_range_chk          CHECK (lat >= -90 AND lat <= 90),
  CONSTRAINT lon_range_chk          CHECK (lon >= -180 AND lon <= 180),
  -- generous constraints to avoid malformed data
  CONSTRAINT address_line_len_chk   CHECK (char_length(address_line) <= 255),
  CONSTRAINT region_len_chk         CHECK (char_length(region) <= 100),
  CONSTRAINT postal_code_len_chk    CHECK (char_length(postal_code) <= 30),
  CONSTRAINT place_id_len_chk       CHECK (char_length(place_id) <= 255),
  CONSTRAINT timezone_len_chk       CHECK (char_length(timezone) <= 60)
);

CREATE TRIGGER trg_update_location_timestamp
BEFORE UPDATE ON location
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- ====================================================================================================
-- Court
-- ====================================================================================================

CREATE TABLE IF NOT EXISTS court (
  id            BIGSERIAL PRIMARY KEY,

  location_id   BIGINT  NOT NULL REFERENCES location(id) ON DELETE CASCADE,
  name          TEXT    NOT NULL,
  court_count   INT     NOT NULL,

  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT court_count_chk    CHECK (court_count >= 0 AND court_count <= 1000),
  CONSTRAINT name_len_chk       CHECK (char_length(name) <= 255)
);

CREATE TRIGGER trg_update_court_timestamp
BEFORE UPDATE ON court
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- ====================================================================================================
-- Court Status
-- ====================================================================================================

CREATE TABLE IF NOT EXISTS court_status (
  -- This primary key is the same as the court id, so a court can only have one status.
  -- This makes querying easier, and also makes court status updates more performant.
  court_id          BIGINT PRIMARY KEY,

  groups_waiting    INT NOT NULL DEFAULT 0,
  courts_occupied   INT NOT NULL DEFAULT 0,

  created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT groups_waiting_chk     CHECK (groups_waiting >= 0 AND groups_waiting <= 1000),
  CONSTRAINT courts_occupied_chk    CHECK (courts_occupied >= 0 AND courts_occupied <= 1000),
  CONSTRAINT court_id_fk FOREIGN KEY (court_id) REFERENCES court(id) ON DELETE CASCADE
);

CREATE TRIGGER trg_update_court_status_timestamp
BEFORE UPDATE ON court_status
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- ====================================================================================================
-- Report
-- ====================================================================================================

CREATE TABLE IF NOT EXISTS report (
  id                BIGSERIAL PRIMARY KEY,

  court_id          BIGINT  NOT NULL REFERENCES court(id),
  courts_occupied   INT     NOT NULL,
  groups_waiting    INT     NOT NULL,

  created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  -- The 1000 magic number is to prevent malicious data,
  -- as there should never be more than 1000 courts occupied or groups waiting.
  CONSTRAINT courts_occupied_chk    CHECK (courts_occupied >= 0 AND courts_occupied <= 1000),
  CONSTRAINT groups_waiting_chk     CHECK (groups_waiting >= 0 AND groups_waiting <= 1000)
);

CREATE TRIGGER trg_update_report_timestamp
BEFORE UPDATE ON report
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
