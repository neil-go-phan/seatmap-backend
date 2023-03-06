CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(16),
  password VARCHAR(255),
salt varchar(50),
  full_name VARCHAR(50),
  role VARCHAR(20),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);


-- INSERT INTO users(username, password, full_name, role, create_at, )