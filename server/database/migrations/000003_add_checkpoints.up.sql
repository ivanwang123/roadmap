BEGIN;

CREATE TABLE IF NOT EXISTS checkpoints (
  id SERIAL PRIMARY KEY,
  title VARCHAR(256) NOT NULL,
  instructions TEXT NOT NULL,
  links TEXT[] NOT NULL DEFAULT '{}',
  roadmap_id INT NOT NULL REFERENCES roadmaps(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER revise_checkpoints_updated_at BEFORE UPDATE ON checkpoints FOR EACH ROW EXECUTE PROCEDURE revise_updated_at();

CREATE TYPE status_type AS ENUM ('complete', 'incomplete', 'skip');

CREATE TABLE IF NOT EXISTS checkpoint_status (
  user_id INT NOT NULL REFERENCES users(id),
  checkpoint_id INT NOT NULL REFERENCES checkpoints(id),
  roadmap_id INT NOT NULL REFERENCES roadmaps(id),
  PRIMARY KEY (user_id, checkpoint_id),
  status status_type NOT NULL DEFAULT 'incomplete'
);

COMMIT;