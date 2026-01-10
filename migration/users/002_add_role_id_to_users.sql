-- migrate/users/002_create_votes_table.sql

-- Add role_id column to users table
ALTER TABLE users ADD COLUMN role_id INT;

ALTER TABLE users ADD CONSTRAINT fk_user_role FOREIGN KEY (role_id) REFERENCES roles(id);

-- Optional: Add comment
COMMENT ON COLUMN users.role_id IS 'Foreign key reference to roles table';
