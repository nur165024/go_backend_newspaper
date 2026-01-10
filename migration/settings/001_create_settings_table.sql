-- migration/settings/001_create_settings_table.sql
CREATE TABLE IF NOT EXISTS settings (
  id SERIAL PRIMARY KEY,
  key VARCHAR(100) UNIQUE NOT NULL,
  value TEXT,
  description TEXT,
  data_type VARCHAR(20) DEFAULT 'string',
  category VARCHAR(50) DEFAULT 'general',
  is_public BOOLEAN DEFAULT false, -- Can be accessed by frontend
  is_editable BOOLEAN DEFAULT true, -- Can be modified via admin panel
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Check constraints
  CONSTRAINT chk_setting_data_type CHECK (data_type IN ('string', 'integer', 'boolean', 'json', 'url', 'email')),
  CONSTRAINT chk_setting_category CHECK (category IN ('general', 'email', 'social', 'seo', 'security', 'payment', 'notification'))
);

CREATE INDEX IF NOT EXISTS idx_settings_category ON settings(category);
CREATE INDEX IF NOT EXISTS idx_settings_key ON settings(key);

-- Insert default settings
INSERT INTO settings (key, value, description, data_type, category, is_public) VALUES 
('site_name', 'News Portal', 'Website name', 'string', 'general', true),
('site_description', 'Latest news and updates', 'Website description', 'string', 'seo', true),
('admin_email', 'admin@example.com', 'Administrator email', 'email', 'general', false),
('posts_per_page', '10', 'Number of posts per page', 'integer', 'general', true),
('enable_comments', 'true', 'Enable comments on news', 'boolean', 'general', true),
('maintenance_mode', 'false', 'Enable maintenance mode', 'boolean', 'general', false)
ON CONFLICT (key) DO NOTHING;
