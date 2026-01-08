-- migration/news_comments/0001_create_news_comments_table.sql
CREATE TABLE IF NOT EXISTS news_comments (
  "id" SERIAL PRIMARY KEY,
  "news_id" INT NOT NULL,
  "user_id" INT NOT NULL,
  "comment_text" TEXT NOT NULL,
  "status" VARCHAR(20) DEFAULT 'approved',
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_comment_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE,
  CONSTRAINT fk_comment_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  
  -- Check constraint
  CONSTRAINT chk_comment_status CHECK (status IN ('pending', 'approved', 'rejected', 'spam'))
);

-- Indexes for performance
CREATE INDEX idx_comments_news_id ON news_comments(news_id);
CREATE INDEX idx_comments_user_id ON news_comments(user_id);
