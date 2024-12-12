CREATE TABLE IF NOT EXISTS daily_questions_solutions(
  id SERIAL NOT NULL PRIMARY KEY,
  daily_question_id UUID NOT NULL REFERENCES daily_questions(id) ON DELETE CASCADE,
  programming_language varchar(40) NOT NULL,
  solution_code TEXT NOT NULL,
  solver_name varchar(100) NOT NULL DEFAULT 'ollama',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP NULL DEFAULT NULL
);