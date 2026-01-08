-- migration/categories/003_create_categories_table.sql
-- Remove parent_id field from categories table
ALTER TABLE categories DROP COLUMN IF EXISTS parent_id;

-- Drop related index if exists
DROP INDEX IF EXISTS idx_categories_parent_id;
