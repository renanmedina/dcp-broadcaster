CREATE TABLE IF NOT EXISTS daily_questions_solutions_languages(
  language_name varchar(100) NOT NULL PRIMARY KEY,
  file_extension varchar(30) NOT NULL,
  enabled boolean NOT NULL DEFAULT true,
  created_at TIMESTAMP NOT NULL default NOW(),
  updated_at TIMESTAMP NOT NULL default NOW(),
  deleted_at TIMESTAMP NULL
)