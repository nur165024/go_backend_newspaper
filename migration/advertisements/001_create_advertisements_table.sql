-- migration/advertisements/001_create_advertisements_table.sql
CREATE TABLE IF NOT EXISTS "advertisements" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "image_url" VARCHAR(500) NOT NULL,
  "link" VARCHAR(500) NOT NULL,
  "position" VARCHAR(50) NOT NULL,
  "is_active" BOOLEAN DEFAULT true,
  "start_date" TIMESTAMP NOT NULL,
  "end_date" TIMESTAMP NOT NULL,
  "priority" INT DEFAULT 1,
  "click_count" INT DEFAULT 0,
  "impression_count" INT DEFAULT 0,
  "target_audience" VARCHAR(100), -- 'all', 'premium', 'registered'
  "created_by" INT NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_ad_creator FOREIGN KEY (created_by) REFERENCES users(id),
  
  -- Check constraints
  CONSTRAINT chk_ad_position CHECK (position IN ('header', 'sidebar', 'footer', 'banner', 'popup', 'inline')),
  CONSTRAINT chk_ad_target CHECK (target_audience IN ('all', 'premium', 'registered')),
  CONSTRAINT chk_ad_dates CHECK (end_date > start_date),
  CONSTRAINT chk_ad_priority CHECK (priority >= 1 AND priority <= 10)
);
