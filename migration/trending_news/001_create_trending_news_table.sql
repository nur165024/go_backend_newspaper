-- migration/trending_news/001_create_trending_news_table.sql
CREATE TABLE IF NOT EXISTS "trending_news" (
  "id" SERIAL PRIMARY KEY,
  "news_id" INT NOT NULL,
  "rank" INT NOT NULL,
  "score" DECIMAL(10,2) NOT NULL,
  "date" DATE NOT NULL,
  "views_count" INT DEFAULT 0,
  "likes_count" INT DEFAULT 0,
  "comments_count" INT DEFAULT 0,
  "shares_count" INT DEFAULT 0,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_trending_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE,
  
  -- Check constraints
  CONSTRAINT chk_trending_rank CHECK (rank > 0),
  CONSTRAINT chk_trending_score CHECK (score >= 0),
  
  -- Unique constraint - one rank per date
  CONSTRAINT unique_rank_per_date UNIQUE (rank, date),
  CONSTRAINT unique_news_per_date UNIQUE (news_id, date)
);

-- Indexes for performance
CREATE INDEX idx_trending_news_id ON trending_news(news_id);
