-- migration/permissions/001_create_permissions_table.sql
CREATE TABLE IF NOT EXISTS permissions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) UNIQUE NOT NULL,
  description TEXT,
  resource VARCHAR(50) NOT NULL,
  action VARCHAR(50) NOT NULL,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Check constraints
  CONSTRAINT chk_permission_resource CHECK (resource IN ('users', 'news', 'categories', 'comments', 'polls', 'tasks', 'roles', 'system')),
  CONSTRAINT chk_permission_action CHECK (action IN ('create', 'read', 'update', 'delete', 'publish', 'moderate', 'manage')),
  
  -- Unique constraint for resource-action combination
  CONSTRAINT unique_resource_action UNIQUE (resource, action)
);
