-- migration/breaking_news/001_create_breaking_news_table.sql
CREATE TABLE IF NOT EXISTS breaking_news (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  news_id INT NOT NULL,
  priority INT DEFAULT 1, -- 1=Low, 2=Medium, 3=High, 4=Critical
  alert_sent BOOLEAN DEFAULT false,
  is_active BOOLEAN DEFAULT true,
  expires_at TIMESTAMP NOT NULL,
  created_by INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_breaking_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE,
  CONSTRAINT fk_breaking_creator FOREIGN KEY (created_by) REFERENCES users(id),
  
  -- Check constraints
  CONSTRAINT chk_breaking_priority CHECK (priority >= 1 AND priority <= 4),
  CONSTRAINT chk_expires_future CHECK (expires_at > created_at),
  
  -- Unique constraint - one breaking news per news article
  CONSTRAINT unique_breaking_news UNIQUE (news_id)
);

-- Indexes for performance
CREATE INDEX idx_breaking_news_id ON breaking_news(news_id);
