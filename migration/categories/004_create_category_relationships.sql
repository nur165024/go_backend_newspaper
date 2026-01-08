-- migration/categories/004_create_category_relationships.sql
 
-- Create category relationships table for parent-child mapping
CREATE TABLE IF NOT EXISTS category_relationships (
    id SERIAL PRIMARY KEY,  -- PostgreSQL standard
    parent_id INT NOT NULL,
    child_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraints
    CONSTRAINT fk_parent_category FOREIGN KEY (parent_id) REFERENCES categories(id) ON DELETE CASCADE,
    CONSTRAINT fk_child_category FOREIGN KEY (child_id) REFERENCES categories(id) ON DELETE CASCADE,
    
    -- Unique constraint to prevent duplicate relationships
    CONSTRAINT unique_parent_child UNIQUE (parent_id, child_id)
);
