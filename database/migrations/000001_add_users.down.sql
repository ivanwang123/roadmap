DROP TRIGGER IF EXISTS revise_user_updated_at ON users;
DROP FUNCTION IF EXISTS revise_updated_at CASCADE;
DROP TABLE IF EXISTS users CASCADE;