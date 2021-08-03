DROP TRIGGER IF EXISTS revise_checkpoints_updated_at ON checkpoints;
DROP TABLE IF EXISTS checkpoints CASCADE;
DROP TABLE IF EXISTS checkpoint_status CASCADE;
DROP TYPE IF EXISTS status_type CASCADE;