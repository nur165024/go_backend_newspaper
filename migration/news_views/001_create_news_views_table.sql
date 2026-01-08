-- migration/news_views/001_create_news_views_table.sql
CREATE TABLE IF NOT EXISTS news_views (
  id SERIAL PRIMARY KEY,
  news_id INT NOT NULL,
  user_id INT, -- Nullable for anonymous users
  ip_address VARCHAR(45), -- IPv6 support (39 chars) + buffer
  user_agent TEXT, -- Browser/device info
  viewed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_view_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE,
  CONSTRAINT fk_view_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
  
  -- Unique constraint to prevent duplicate views (same user/IP within timeframe)
  CONSTRAINT unique_user_news_view UNIQUE (news_id, user_id)
);

-- Indexes for performance
CREATE INDEX idx_views_news_id ON news_views(news_id);
CREATE INDEX idx_views_user_id ON news_views(user_id);
