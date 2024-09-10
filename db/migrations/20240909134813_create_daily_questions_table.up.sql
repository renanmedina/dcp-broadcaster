CREATE TABLE IF NOT EXISTS daily_questions(
  id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  original_id varchar(200) NULL,
  difficulty_level varchar(8) NOT NULL,
  received_at timestamp NOT NULL,
  title VARCHAR(500) NOT NULL,
  question_email_body TEXT NOT NULL,
  question_text TEXT NOT NULL,
  company_name varchar(220) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL DEFAULT NULL
);