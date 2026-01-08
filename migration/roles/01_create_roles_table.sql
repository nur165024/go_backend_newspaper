-- migration/roles/001_create_roles_table.sql
CREATE TABLE IF NOT EXISTS roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL,
  description TEXT,
  is_active BOOLEAN DEFAULT true,
  priority INT DEFAULT 0, -- Role hierarchy (higher = more access)
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Check constraint
  CONSTRAINT chk_role_priority CHECK (priority >= 0 AND priority <= 100)
);

-- Insert default system roles
INSERT INTO roles (name, description, priority) VALUES 
('super_admin', 'Super Administrator with full access', 100),
('admin', 'Administrator with management access', 80),
('editor', 'Content Editor with publishing rights', 60),
('author', 'Content Author with limited publishing', 40),
('user', 'Regular User with basic access', 20);
