-- Trigger to delete location when court is deleted
-- Since courts and locations have a 1-to-1 relationship,
-- deleting a court should also delete its location

CREATE OR REPLACE FUNCTION delete_court_location()
RETURNS TRIGGER AS $$
BEGIN
  -- Delete the location associated with this court
  DELETE FROM location WHERE id = OLD.location_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_delete_court_location
AFTER DELETE ON court
FOR EACH ROW
EXECUTE FUNCTION delete_court_location();
