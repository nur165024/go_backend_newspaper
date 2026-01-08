-- migration/news_likes/001_create_news_likes_table.sql
CREATE TABLE IF NOT EXISTS news_likes (
  id SERIAL PRIMARY KEY,
  news_id INT NOT NULL,
  user_id INT NOT NULL,
  like_type VARCHAR(20) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_like_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE,
  CONSTRAINT fk_like_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  
  -- Check constraint for like types
  CONSTRAINT chk_like_type CHECK (like_type IN ('like', 'dislike', 'love', 'angry', 'sad')),
  
  -- Unique constraint to prevent duplicate likes from same user
  CONSTRAINT unique_user_news_like UNIQUE (news_id, user_id)
);

-- Indexes for performance
CREATE INDEX idx_likes_news_id ON news_likes(news_id);
CREATE INDEX idx_likes_user_id ON news_likes(user_id);
