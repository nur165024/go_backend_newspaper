CREATE TABLE seo_metadata (
  id SERIAL PRIMARY KEY,
  entity_type VARCHAR(20) NOT NULL, -- 'news', 'category'
  entity_id INT NOT NULL,
  meta_title VARCHAR(255),
  meta_description TEXT,
  meta_keywords VARCHAR(500),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);