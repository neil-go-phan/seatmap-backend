
CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  role_name varchar(20) UNIQUE
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(16) UNIQUE,
  password VARCHAR(255),
  salt varchar(50),
  full_name VARCHAR(50),
  role varchar(20),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT fk_role
    FOREIGN KEY(role) 
	  REFERENCES roles(role_name)
);

INSERT INTO roles(role_name) VALUES('CEO');
INSERT INTO roles(role_name) VALUES('CTO');
INSERT INTO roles(role_name) VALUES('HR');
INSERT INTO roles(role_name) VALUES('Golang developer');
INSERT INTO roles(role_name) VALUES('Ruby developer');
INSERT INTO roles(role_name) VALUES('Nodejs developer');
INSERT INTO roles(role_name) VALUES('Admin');
INSERT INTO roles(role_name) VALUES('Staff');

select * from roles;
select * from users;

DROP TABLE users;
DROP TABLE roles;

UPDATE users
SET role = 'Admin' 
WHERE username = 'admin1234';