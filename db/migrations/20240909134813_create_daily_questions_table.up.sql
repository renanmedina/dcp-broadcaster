CREATE TABLE IF NOT EXISTS daily_questions(
  question_id SERIAL NOT NULL PRIMARY KEY,
  received_at timestamp NOT NULL,
  title VARCHAR(500) NOT NULL,
  question_email_body TEXT NOT NULL,
  question_text TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL DEFAULT NULL
);