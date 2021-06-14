BEGIN;

CREATE TABLE IF NOT EXISTS roadmaps (
  id SERIAL PRIMARY KEY,
  title VARCHAR(256) NOT NULL,
  description TEXT NOT NULL,
  creator_id INT NOT NULL REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER revise_roadmap_updated_at BEFORE UPDATE ON roadmaps FOR EACH ROW EXECUTE PROCEDURE revise_updated_at();

CREATE TABLE IF NOT EXISTS roadmap_follower (
  user_id INT NOT NULL REFERENCES users(id),
  roadmap_id INT NOT NULL REFERENCES roadmaps(id),
  PRIMARY KEY (user_id, roadmap_id)
);

COMMIT;