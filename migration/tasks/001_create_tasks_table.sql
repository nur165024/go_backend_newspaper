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