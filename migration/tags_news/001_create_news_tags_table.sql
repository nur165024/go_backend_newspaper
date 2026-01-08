-- migration/news_tags/001_create_news_tags_table.sql
CREATE TABLE IF NOT EXISTS news_tags (
  id SERIAL PRIMARY KEY,
  news_id INT NOT NULL,
  tag_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  
  CONSTRAINT fk_news_tag_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE,
  CONSTRAINT fk_news_tag_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
  
  CONSTRAINT unique_news_tag UNIQUE (news_id, tag_id)
);
