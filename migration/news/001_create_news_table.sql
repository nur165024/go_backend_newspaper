-- migration/news/001_create_news_table.sql

CREATE TABLE IF NOT EXISTS "news" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "slug" VARCHAR(255) UNIQUE NOT NULL,
  "content" TEXT,
  "excerpt" TEXT,
  "featured_image" VARCHAR(255),
  "images" JSONB,
  "language" VARCHAR(10) DEFAULT 'bn',
  "is_premium" BOOLEAN DEFAULT false,
  "status" VARCHAR(20) DEFAULT 'draft',
  "is_featured" BOOLEAN DEFAULT false,
  "published_at" TIMESTAMP,
  "scheduled_at" TIMESTAMP, -- When to auto-publish
  "category_id" INT,
  "user_id" INT NOT NULL,
  -- SEO fields
  "meta_title" VARCHAR(255),
  "meta_description" TEXT,
  "meta_keywords" VARCHAR(500),
  -- Additional fields
  "reading_time" INT,
  "source" VARCHAR(255),
  "priority" INT DEFAULT 0,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Constraints
  CONSTRAINT fk_news_category FOREIGN KEY (category_id) REFERENCES categories(id),
  CONSTRAINT fk_news_user FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT chk_status CHECK (status IN ('draft', 'published', 'archived', 'scheduled'))
);

-- Indexes for better performance
CREATE INDEX idx_news_category_id ON news(category_id);
CREATE INDEX idx_news_user_id ON news(user_id);
CREATE INDEX idx_news_slug ON news(slug);
