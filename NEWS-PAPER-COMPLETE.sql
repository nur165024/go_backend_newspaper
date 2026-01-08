CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "user_name" varchar UNIQUE,
  "email" varchar UNIQUE,
  "password" varchar,
  "designation" varchar,
  "bio" text,
  "profile_picture" varchar,
  "is_active" boolean DEFAULT true,
  "is_verified" boolean DEFAULT false,
  "verification_token" varchar,
  "reset_password_token" varchar,
  "reset_password_expires" timestamp,
  "last_login" timestamp,
  "role_id" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "slug" varchar UNIQUE,
  "description" text,
  "image_url" varchar,
  "sort_order" int,
  "is_active" boolean DEFAULT true,
  "meta_title" varchar,
  "meta_description" text,
  "meta_keywords" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "news" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "slug" varchar UNIQUE,
  "content" text,
  "excerpt" text,
  "featured_image" varchar,
  "images" json,
  "is_premium" boolean DEFAULT false,
  "status" varchar DEFAULT 'draft',
  "is_featured" boolean DEFAULT false,
  "published_at" timestamp,
  "category_id" int,
  "user_id" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "news_comments" (
  "id" SERIAL PRIMARY KEY,
  "news_id" int,
  "user_id" int,
  "comment_text" text,
  "status" varchar DEFAULT 'pending',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "news_likes" (
  "id" SERIAL PRIMARY KEY,
  "news_id" int,
  "user_id" int,
  "like_type" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "news_views" (
  "id" SERIAL PRIMARY KEY,
  "news_id" int,
  "user_id" int,
  "ip_address" varchar,
  "viewed_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "tags" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "slug" varchar UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "news_tags" (
  "id" SERIAL PRIMARY KEY,
  "news_id" int,
  "tag_id" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "polls" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "question" TEXT NOT NULL,
  "poll_type" VARCHAR(20) NOT NULL,
  "status" VARCHAR(20) DEFAULT 'draft',
  "is_featured" BOOLEAN DEFAULT false,
  "allow_anonymous" BOOLEAN DEFAULT true,
  "max_votes_per_user" INT DEFAULT 1,
  "start_date" TIMESTAMP,
  "end_date" TIMESTAMP,
  "created_by" INT NOT NULL,
  "total_votes" INT DEFAULT 0,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE "poll_options" (
  "id" SERIAL PRIMARY KEY,
  "poll_id" INT NOT NULL,
  "option_text" VARCHAR(255) NOT NULL,
  "option_image" VARCHAR(255),
  "sort_order" INT DEFAULT 0,
  "vote_count" INT DEFAULT 0,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE "votes" (
  "id" SERIAL PRIMARY KEY,
  "poll_id" INT NOT NULL,
  "option_id" INT NOT NULL,
  "user_id" INT,
  "ip_address" VARCHAR(45),
  "user_agent" TEXT,
  "rating_value" INT,
  "voted_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE "newsletters" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "content" text,
  "subject" varchar,
  "sent_to" varchar,
  "sent_at" timestamp,
  "status" varchar DEFAULT 'draft',
  "created_by" int,
  "open_rate" decimal,
  "click_rate" decimal,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "newsletter_subscriptions" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar,
  "user_id" int,
  "is_active" boolean DEFAULT true,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "breaking_news" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "news_id" int,
  "priority" int,
  "alert_send" boolean DEFAULT false,
  "expires_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "trending_news" (
  "id" SERIAL PRIMARY KEY,
  "news_id" int,
  "rank" int,
  "score" decimal,
  "date" date,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "notifications" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "type" varchar,
  "title" varchar,
  "message" text,
  "link" varchar,
  "is_read" boolean DEFAULT false,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "advertisements" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "image_url" varchar,
  "link" varchar,
  "position" varchar,
  "is_active" boolean DEFAULT true,
  "start_date" timestamp,
  "end_date" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "tasks" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "description" text,
  "task_type" varchar,
  "priority" varchar,
  "status" varchar DEFAULT 'pending',
  "created_by" int,
  "assigned_to" int,
  "due_date" timestamp,
  "started_at" timestamp,
  "completed_at" timestamp,
  "progress_percentage" int DEFAULT 0,
  "estimated_hours" decimal,
  "actual_hours" decimal,
  "notes" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "task_comments" (
  "id" SERIAL PRIMARY KEY,
  "task_id" int,
  "user_id" int,
  "comment_text" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "task_histories" (
  "id" SERIAL PRIMARY KEY,
  "task_id" int,
  "user_id" int,
  "action" varchar,
  "old_value" text,
  "new_value" text,
  "description" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "description" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "permissions" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "description" text,
  "resource" varchar,
  "action" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "role_permissions" (
  "id" SERIAL PRIMARY KEY,
  "role_id" int,
  "permission_id" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "settings" (
  "id" SERIAL PRIMARY KEY,
  "key" varchar UNIQUE,
  "value" text,
  "description" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Foreign Key Constraints
ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "news" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "news" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "news_comments" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");
ALTER TABLE "news_comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "news_likes" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");
ALTER TABLE "news_likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "news_views" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");
ALTER TABLE "news_views" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "news_tags" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");
ALTER TABLE "news_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");

ALTER TABLE "polls" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

ALTER TABLE "poll_options" ADD FOREIGN KEY ("poll_id") REFERENCES "polls" ("id");

ALTER TABLE "votes" ADD FOREIGN KEY ("poll_id") REFERENCES "polls" ("id");
ALTER TABLE "votes" ADD FOREIGN KEY ("option_id") REFERENCES "poll_options" ("id");
ALTER TABLE "votes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "poll_results" ADD FOREIGN KEY ("poll_id") REFERENCES "polls" ("id");
ALTER TABLE "poll_results" ADD FOREIGN KEY ("option_id") REFERENCES "poll_options" ("id");

ALTER TABLE "newsletters" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

ALTER TABLE "newsletter_subscriptions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "breaking_news" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");

ALTER TABLE "trending_news" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");

ALTER TABLE "notifications" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");
ALTER TABLE "tasks" ADD FOREIGN KEY ("assigned_to") REFERENCES "users" ("id");

ALTER TABLE "task_comments" ADD FOREIGN KEY ("task_id") REFERENCES "tasks" ("id");
ALTER TABLE "task_comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "task_histories" ADD FOREIGN KEY ("task_id") REFERENCES "tasks" ("id");
ALTER TABLE "task_histories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "role_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

-- Performance Indexes
CREATE INDEX idx_news_category ON news(category_id);
CREATE INDEX idx_news_user ON news(user_id);
CREATE INDEX idx_news_status ON news(status);
CREATE INDEX idx_news_published ON news(published_at);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role_id);
CREATE INDEX idx_news_slug ON news(slug);
CREATE INDEX idx_categories_slug ON categories(slug);
CREATE INDEX idx_tags_slug ON tags(slug);
CREATE INDEX idx_news_comments_news ON news_comments(news_id);
CREATE INDEX idx_news_likes_news ON news_likes(news_id);
CREATE INDEX idx_news_views_news ON news_views(news_id);
CREATE INDEX idx_polls_status ON polls(status);
CREATE INDEX idx_votes_poll ON votes(poll_id);
CREATE INDEX idx_notifications_user ON notifications(user_id);
CREATE INDEX idx_tasks_assigned ON tasks(assigned_to);
CREATE INDEX idx_tasks_status ON tasks(status);