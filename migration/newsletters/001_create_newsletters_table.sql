-- migration/newsletters/001_create_newsletters_table.sql
CREATE TABLE IF NOT EXISTS "newsletters" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "content" TEXT NOT NULL,
  "subject" VARCHAR(255) NOT NULL,
  "sent_to" VARCHAR(20) DEFAULT 'all', -- 'all', 'subscribers', 'segment'
  "sent_at" TIMESTAMP,
  "status" VARCHAR(20) DEFAULT 'draft',
  "created_by" INT NOT NULL,
  "total_sent" INT DEFAULT 0,
  "total_opened" INT DEFAULT 0,
  "total_clicked" INT DEFAULT 0,
  "open_rate" DECIMAL(5,2) DEFAULT 0.00,
  "click_rate" DECIMAL(5,2) DEFAULT 0.00,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  CONSTRAINT fk_newsletter_creator FOREIGN KEY (created_by) REFERENCES users(id),
  CONSTRAINT chk_newsletter_status CHECK (status IN ('draft', 'scheduled', 'sent', 'failed')),
  CONSTRAINT chk_sent_to CHECK (sent_to IN ('all', 'subscribers', 'segment'))
);
