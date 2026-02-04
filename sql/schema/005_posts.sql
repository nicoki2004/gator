-- +goose Up
CREATE TABLE posts (
  id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  title TEXT, 
  url TEXT UNIQUE NOT NULL, 
  description TEXT,
  published_at TIMESTAMP NOT NULL DEFAULT NOW(),
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE posts;
