CREATE TABLE IF NOT EXISTS news (
  id CHAR(27) PRIMARY KEY,
  title TEXT,
  description TEXT,
  h1 TEXT,
  text TEXT,
  user_id CHAR(27) NOT NULL,
  published BOOLEAN DEFAULT false,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);