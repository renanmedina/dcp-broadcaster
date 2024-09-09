CREATE TABLE IF NOT EXISTS users(
  user_id SERIAL NOT NULL PRIMARY KEY,
  username varchar(60) NOT NULL,
  phone_number varchar(18) NULL DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL DEFAULT NULL
);