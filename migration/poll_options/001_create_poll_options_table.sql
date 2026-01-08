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
