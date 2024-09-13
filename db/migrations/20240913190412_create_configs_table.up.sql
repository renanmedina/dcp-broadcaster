CREATE TABLE IF NOT EXISTS configs(
  config_key varchar(100) NOT NULL UNIQUE,
  config_value jsonb NULL DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP NULL
);