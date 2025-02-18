CREATE TABLE IF NOT EXISTS daily_questions(
  id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  original_id varchar(200) NOT NULL UNIQUE,
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

CREATE INDEX idx_questions_difficulty_level ON daily_questions(difficulty_level);
CREATE INDEX idx_questions_received_at ON daily_questions(received_at);
CREATE INDEX idx_questions_company_name ON daily_questions(company_name);
CREATE INDEX idx_questions_original_id ON daily_questions(original_id);