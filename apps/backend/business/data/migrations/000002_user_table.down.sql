ALTER TABLE users
DROP COLUMN IF EXISTS preference_id;

DROP TABLE IF EXISTS users_preferences;

DROP TABLE IF EXISTS users_access_controls;

DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS system_appearance_type;

DROP TYPE IF EXISTS user_account_status_type;