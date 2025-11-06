-- Remove the trigger that deletes location when court is deleted

DROP TRIGGER IF EXISTS trg_delete_court_location ON court;
DROP FUNCTION IF EXISTS delete_court_location();
