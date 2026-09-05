CREATE TRIGGER trg_b_update_messages_time
BEFORE UPDATE OF section_id ON messages
FOR EACH ROW
WHEN (OLD.section_id IS DISTINCT FROM NEW.section_id) 
EXECUTE FUNCTION set_updated_at();