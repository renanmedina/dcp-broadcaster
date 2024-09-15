CREATE TABLE IF NOT EXISTS users(
  id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  username varchar(60) NOT NULL,
  name varchar(300) NOT NULL,
  phone_number varchar(18) NULL DEFAULT NULL,
  subscribed smallint NOT NULL DEFAULT 1,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL DEFAULT NULL
);

CREATE INDEX idx_users_phone_number ON users(phone_number);