CREATE TABLE IF NOT EXISTS events(
  event_id SERIAL NOT NULL PRIMARY KEY,
  event_name TEXT NOT NULL,
  object_id VARCHAR(150) NOT NULL, 
  object_type VARCHAR(300) NOT NULL,
  event_data jsonb NULL,
  author TEXT NOT NULL DEFAULT 'system-dcp-broadcaster',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_events_event_name ON events(event_name);
CREATE INDEX idx_events_object_id ON events(object_id);
CREATE INDEX idx_events_created_at ON events(created_at);
CREATE INDEX idx_events_updated_at ON events(updated_at);
