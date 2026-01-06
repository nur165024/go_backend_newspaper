-- Add new filed to categories table
ALTER TABLE categories 
ADD COLUMN sort_order INTEGER DEFAULT 0,
ADD COLUMN image_url VARCHAR(255),
ADD COLUMN meta_title VARCHAR(255),
ADD COLUMN meta_description TEXT,
ADD COLUMN meta_keywords TEXT;