BEGIN;

CREATE TABLE IF NOT EXISTS checkpoints (
  id SERIAL PRIMARY KEY,
  title VARCHAR(256) NOT NULL,
  instructions TEXT NOT NULL,
  links TEXT[] NOT NULL,
  isCompleted BOOLEAN,
  roadmap INT NOT NULL REFERENCES roadmaps(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER revise_checkpoints_updated_at BEFORE UPDATE ON checkpoints FOR EACH ROW EXECUTE PROCEDURE revise_updated_at();

COMMIT;