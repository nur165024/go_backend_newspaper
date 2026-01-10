-- migration/categories/002_create_categories_table.sql
-- Add new filed to categories table
ALTER TABLE categories 
ADD COLUMN sort_order INTEGER DEFAULT 0,
ADD COLUMN image_url VARCHAR(255),
ADD COLUMN meta_title VARCHAR(255),
ADD COLUMN meta_description TEXT,