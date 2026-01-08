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
