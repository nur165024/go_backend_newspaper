-- migration/task_histories/001_create_task_histories_table.sql
CREATE TABLE IF NOT EXISTS "task_histories" (
  "id" SERIAL PRIMARY KEY,
  "task_id" INT NOT NULL,
  "user_id" INT NOT NULL,
  "action" VARCHAR(50) NOT NULL,
  "old_value" TEXT,
  "new_value" TEXT,
  "description" TEXT,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_history_task FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
  CONSTRAINT fk_history_user FOREIGN KEY (user_id) REFERENCES users(id),
  
  -- Check constraint
  CONSTRAINT chk_history_action CHECK (action IN ('created', 'updated', 'assigned', 'status_changed', 'completed', 'commented'))
);

CREATE INDEX idx_task_histories_task_id ON task_histories(task_id);
CREATE INDEX idx_task_histories_user_id ON task_histories(user_id);
