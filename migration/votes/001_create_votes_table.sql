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
