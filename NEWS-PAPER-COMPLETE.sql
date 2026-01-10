-- migration/advertisements/001_create_advertisements_table.sql
CREATE TABLE IF NOT EXISTS advertisements (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  image_url VARCHAR(500) NOT NULL,
  link VARCHAR(500) NOT NULL,
  position VARCHAR(50) NOT NULL,
  is_active BOOLEAN DEFAULT true,
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL,
  priority INT DEFAULT 1,
  click_count INT DEFAULT 0,
  impression_count INT DEFAULT 0,
  target_audience VARCHAR(100), -- 'all', 'premium', 'registered'
  created_by INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_ad_creator FOREIGN KEY (created_by) REFERENCES users(id),
  
  -- Check constraints
  CONSTRAINT chk_ad_position CHECK (position IN ('header', 'sidebar', 'footer', 'banner', 'popup', 'inline')),
  CONSTRAINT chk_ad_target CHECK (target_audience IN ('all', 'premium', 'registered')),
  CONSTRAINT chk_ad_dates CHECK (end_date > start_date),
  CONSTRAINT chk_ad_priority CHECK (priority >= 1 AND priority <= 10)
);
-- --------------------------------------------------------
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

-- -----------------------------------------------------------------------

-- migration/categories/001_create_categories_table.sql
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sort_order INTEGER DEFAULT 0,
    image_url VARCHAR(255),
    meta_title VARCHAR(255),
    meta_description TEXT,
    meta_keywords TEXT,
);

--------------------------------------------------------------------------
-- migration/categories/004_create_category_relationships.sql
 
-- Create category relationships table for parent-child mapping
CREATE TABLE IF NOT EXISTS category_relationships (
    id SERIAL PRIMARY KEY,  -- PostgreSQL standard
    parent_id INT NOT NULL,
    child_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraints
    CONSTRAINT fk_parent_category FOREIGN KEY (parent_id) REFERENCES categories(id) ON DELETE CASCADE,
    CONSTRAINT fk_child_category FOREIGN KEY (child_id) REFERENCES categories(id) ON DELETE CASCADE,
    
    -- Unique constraint to prevent duplicate relationships
    CONSTRAINT unique_parent_child UNIQUE (parent_id, child_id)
);
----------------------------------------------------------------------

-- migration/news/001_create_news_table.sql

CREATE TABLE IF NOT EXISTS news (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  slug VARCHAR(255) UNIQUE NOT NULL,
  content TEXT,
  excerpt TEXT,
  featured_image VARCHAR(255),
  images JSONB,
  language VARCHAR(10) DEFAULT 'bn',
  is_premium BOOLEAN DEFAULT false,
  status VARCHAR(20) DEFAULT 'draft',
  is_featured BOOLEAN DEFAULT false,
  published_at TIMESTAMP,
  scheduled_at TIMESTAMP, -- When to auto-publish
  category_id INT,
  user_id INT NOT NULL,
  -- SEO fields
  meta_title VARCHAR(255),
  meta_description TEXT,
  meta_keywords VARCHAR(500),
  -- Additional fields
  reading_time INT,
  source VARCHAR(255),
  priority INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Constraints
  CONSTRAINT fk_news_category FOREIGN KEY (category_id) REFERENCES categories(id),
  CONSTRAINT fk_news_user FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT chk_status CHECK (status IN ('draft', 'published', 'archived', 'scheduled'))
);

-- Indexes for better performance
CREATE INDEX idx_news_category_id ON news(category_id);
CREATE INDEX idx_news_user_id ON news(user_id);
CREATE INDEX idx_news_slug ON news(slug);
------------------------------------------------------------------------

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

-- ------------------------------------------------------------------------

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

-- -------------------------------------------------

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

-- -------------------------------------------------

-- migration/newsletter_subscriptions/001_create_newsletter_subscriptions_table.sql
CREATE TABLE IF NOT EXISTS newsletter_subscriptions (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  user_id INT,
  is_active BOOLEAN DEFAULT true,
  subscribed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  unsubscribed_at TIMESTAMP,
  verification_token VARCHAR(255),
  is_verified BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  CONSTRAINT fk_subscription_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE INDEX idx_subscriptions_email ON newsletter_subscriptions(email);
CREATE INDEX idx_subscriptions_user_id ON newsletter_subscriptions(user_id);

-- --------------------------------------------------------

-- migration/newsletters/001_create_newsletters_table.sql
CREATE TABLE IF NOT EXISTS newsletters (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  subject VARCHAR(255) NOT NULL,
  sent_to VARCHAR(20) DEFAULT 'all', -- 'all', 'subscribers', 'segment'
  sent_at TIMESTAMP,
  status VARCHAR(20) DEFAULT 'draft',
  created_by INT NOT NULL,
  total_sent INT DEFAULT 0,
  total_opened INT DEFAULT 0,
  total_clicked INT DEFAULT 0,
  open_rate DECIMAL(5,2) DEFAULT 0.00,
  click_rate DECIMAL(5,2) DEFAULT 0.00,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  CONSTRAINT fk_newsletter_creator FOREIGN KEY (created_by) REFERENCES users(id),
  CONSTRAINT chk_newsletter_status CHECK (status IN ('draft', 'scheduled', 'sent', 'failed')),
  CONSTRAINT chk_sent_to CHECK (sent_to IN ('all', 'subscribers', 'segment'))
);

-- --------------------------------------------------------

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

-- --------------------------------------------------------

-- migration/permissions/001_create_permissions_table.sql
CREATE TABLE IF NOT EXISTS permissions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) UNIQUE NOT NULL,
  description TEXT,
  resource VARCHAR(50) NOT NULL,
  action VARCHAR(50) NOT NULL,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Check constraints
  CONSTRAINT chk_permission_resource CHECK (resource IN ('users', 'news', 'categories', 'comments', 'polls', 'tasks', 'roles', 'system')),
  CONSTRAINT chk_permission_action CHECK (action IN ('create', 'read', 'update', 'delete', 'publish', 'moderate', 'manage')),
  
  -- Unique constraint for resource-action combination
  CONSTRAINT unique_resource_action UNIQUE (resource, action)
);

-- ------------------------------------------------------------------------

-- migration/poll_options/001_create_poll_options_table.sql
CREATE TABLE IF NOT EXISTS poll_options (
  id SERIAL PRIMARY KEY,
  poll_id INT NOT NULL,
  option_text VARCHAR(255) NOT NULL,
  option_image VARCHAR(255),
  sort_order INT DEFAULT 0,
  vote_count INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  CONSTRAINT fk_option_poll FOREIGN KEY (poll_id) REFERENCES polls(id) ON DELETE CASCADE
);

-- ---------------------------------------------------------------------------

-- migration/polls/001_create_polls_table.sql
CREATE TABLE IF NOT EXISTS polls (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  question TEXT NOT NULL,
  poll_type VARCHAR(20) NOT NULL,
  status VARCHAR(20) DEFAULT 'draft',
  is_featured BOOLEAN DEFAULT false,
  allow_anonymous BOOLEAN DEFAULT true,
  max_votes_per_user INT DEFAULT 1,
  start_date TIMESTAMP,
  end_date TIMESTAMP,
  created_by INT NOT NULL,
  total_votes INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  CONSTRAINT fk_poll_creator FOREIGN KEY (created_by) REFERENCES users(id),
  CONSTRAINT chk_poll_type CHECK (poll_type IN ('single_choice', 'multiple_choice', 'rating')),
  CONSTRAINT chk_poll_status CHECK (status IN ('draft', 'active', 'closed'))
);

-- ------------------------------------------------------------------------

-- migration/role_permissions/001_create_role_permissions_table.sql
CREATE TABLE IF NOT EXISTS role_permissions (
  id SERIAL PRIMARY KEY,
  role_id INT NOT NULL,
  permission_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_role_permission_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
  CONSTRAINT fk_role_permission_permission FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
  
  -- Unique constraint to prevent duplicate assignments
  CONSTRAINT unique_role_permission UNIQUE (role_id, permission_id)
);

CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);

-- ------------------------------------------------------------------------

-- migration/roles/001_create_roles_table.sql
CREATE TABLE IF NOT EXISTS roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL,
  description TEXT,
  is_active BOOLEAN DEFAULT true,
  priority INT DEFAULT 0, -- Role hierarchy (higher = more access)
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Check constraint
  CONSTRAINT chk_role_priority CHECK (priority >= 0 AND priority <= 100)
);

-- Insert default system roles
INSERT INTO roles (name, description, priority) VALUES 
('super_admin', 'Super Administrator with full access', 100),
('admin', 'Administrator with management access', 80),
('editor', 'Content Editor with publishing rights', 60),
('author', 'Content Author with limited publishing', 40),
('user', 'Regular User with basic access', 20);

-- ------------------------------------------------------------------------

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

-- Insert default settings
INSERT INTO settings (key, value, description, data_type, category, is_public) VALUES 
('site_name', 'News Portal', 'Website name', 'string', 'general', true),
('site_description', 'Latest news and updates', 'Website description', 'string', 'seo', true),
('admin_email', 'admin@example.com', 'Administrator email', 'email', 'general', false),
('posts_per_page', '10', 'Number of posts per page', 'integer', 'general', true),
('enable_comments', 'true', 'Enable comments on news', 'boolean', 'general', true),
('maintenance_mode', 'false', 'Enable maintenance mode', 'boolean', 'general', false);

-- -----------------------------------------------------------------------------

-- migration/tags/001_create_tags_table.sql
CREATE TABLE IF NOT EXISTS tags (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  slug VARCHAR(100) UNIQUE NOT NULL,
  description TEXT,
  color VARCHAR(7), -- Hex color code (#FF5733)
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-----------------------------------------------------------------------------

-- migration/news_tags/001_create_news_tags_table.sql
CREATE TABLE IF NOT EXISTS news_tags (
  id SERIAL PRIMARY KEY,
  news_id INT NOT NULL,
  tag_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  
  CONSTRAINT fk_news_tag_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE,
  CONSTRAINT fk_news_tag_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
  
  CONSTRAINT unique_news_tag UNIQUE (news_id, tag_id)
);

-------------------------------------------------------------------------

-- migration/task_comments/001_create_task_comments_table.sql
CREATE TABLE IF NOT EXISTS task_comments (
  id SERIAL PRIMARY KEY,
  task_id INT NOT NULL,
  user_id INT NOT NULL,
  comment_text TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_comment_task FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
  CONSTRAINT fk_comment_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_task_comments_task_id ON task_comments(task_id);
CREATE INDEX idx_task_comments_user_id ON task_comments(user_id);

--------------------------------------------------------------------------------

-- migration/task_histories/001_create_task_histories_table.sql
CREATE TABLE IF NOT EXISTS task_histories (
  id SERIAL PRIMARY KEY,
  task_id INT NOT NULL,
  user_id INT NOT NULL,
  action VARCHAR(50) NOT NULL,
  old_value TEXT,
  new_value TEXT,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_history_task FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
  CONSTRAINT fk_history_user FOREIGN KEY (user_id) REFERENCES users(id),
  
  -- Check constraint
  CONSTRAINT chk_history_action CHECK (action IN ('created', 'updated', 'assigned', 'status_changed', 'completed', 'commented'))
);

CREATE INDEX idx_task_histories_task_id ON task_histories(task_id);
CREATE INDEX idx_task_histories_user_id ON task_histories(user_id);

--------------------------------------------------------------------------------

-- migration/tasks/001_create_tasks_table.sql

CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  task_type VARCHAR(50) NOT NULL,
  priority VARCHAR(20) NOT NULL,
  status VARCHAR(20) DEFAULT 'pending',
  created_by INT NOT NULL,
  assigned_to INT,
  due_date TIMESTAMP,
  started_at TIMESTAMP,
  completed_at TIMESTAMP,
  progress_percentage INT DEFAULT 0,
  estimated_hours DECIMAL(5,2),
  actual_hours DECIMAL(5,2),
  notes TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_task_creator FOREIGN KEY (created_by) REFERENCES users(id),
  CONSTRAINT fk_task_assignee FOREIGN KEY (assigned_to) REFERENCES users(id),
  
  -- Check constraints
  CONSTRAINT chk_task_type CHECK (task_type IN ('bug', 'feature', 'improvement', 'research', 'documentation')),
  CONSTRAINT chk_task_priority CHECK (priority IN ('low', 'medium', 'high', 'urgent')),
  CONSTRAINT chk_task_status CHECK (status IN ('pending', 'in_progress', 'completed', 'cancelled', 'on_hold')),
  CONSTRAINT chk_progress_range CHECK (progress_percentage >= 0 AND progress_percentage <= 100)
);

CREATE INDEX idx_tasks_assigned_to ON tasks(assigned_to);
CREATE INDEX idx_tasks_created_by ON tasks(created_by);

----------------------------------------------------------------------------------

-- migration/trending_news/001_create_trending_news_table.sql
CREATE TABLE IF NOT EXISTS trending_news (
  id SERIAL PRIMARY KEY,
  news_id INT NOT NULL,
  rank INT NOT NULL,
  score DECIMAL(10,2) NOT NULL,
  date DATE NOT NULL,
  views_count INT DEFAULT 0,
  likes_count INT DEFAULT 0,
  comments_count INT DEFAULT 0,
  shares_count INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
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

------------------------------------------------------------------------------

-- migrations/users/001_create_users_table.sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    user_name VARCHAR(50) UNIQUE,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    designation VARCHAR(100),
    bio TEXT,
    role_id INT,
    profile_picture VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    is_verified BOOLEAN DEFAULT FALSE,
    verification_token VARCHAR(255),
    reset_password_token VARCHAR(255),
    reset_password_expires TIMESTAMP,
    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-----------------------------------------------------------------------------

-- migrate/votes/001_create_votes_table.sql
CREATE TABLE IF NOT EXISTS votes (
  id SERIAL PRIMARY KEY,
  poll_id INT NOT NULL,
  option_id INT NOT NULL,
  user_id INT,
  ip_address VARCHAR(45),
  user_agent TEXT,
  rating_value INT,
  voted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  CONSTRAINT fk_vote_poll FOREIGN KEY (poll_id) REFERENCES polls(id) ON DELETE CASCADE,
  CONSTRAINT fk_vote_option FOREIGN KEY (option_id) REFERENCES poll_options(id) ON DELETE CASCADE,
  CONSTRAINT fk_vote_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);


