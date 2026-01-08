-- migration/notifications/001_create_notifications_table.sql
CREATE TABLE IF NOT EXISTS notifications (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  type VARCHAR(50) NOT NULL,
  title VARCHAR(255) NOT NULL,
  message TEXT NOT NULL,
  link VARCHAR(500),
  is_read BOOLEAN DEFAULT false,
  read_at TIMESTAMP,
  expires_at TIMESTAMP,
  priority VARCHAR(20) DEFAULT 'normal',
  data JSONB, -- Additional metadata
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_notification_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  
  -- Check constraints
  CONSTRAINT chk_notification_type CHECK (type IN ('news', 'comment', 'like', 'breaking', 'newsletter', 'system')),
  CONSTRAINT chk_notification_priority CHECK (priority IN ('low', 'normal', 'high', 'urgent'))
);

-- Indexes for performance
CREATE INDEX idx_notifications_user_id ON notifications(user_id);
