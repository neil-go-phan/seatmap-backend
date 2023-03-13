
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

CREATE UNIQUE INDEX index_roles_role_name on roles (role_name);
CREATE UNIQUE INDEX index_users_username on users (username);

INSERT INTO roles(role_name) VALUES('CEO');
INSERT INTO roles(role_name) VALUES('CTO');
INSERT INTO roles(role_name) VALUES('HR');
INSERT INTO roles(role_name) VALUES('Golang developer');
INSERT INTO roles(role_name) VALUES('Ruby developer');
INSERT INTO roles(role_name) VALUES('Nodejs developer');
INSERT INTO roles(role_name) VALUES('Mobile developer');
INSERT INTO roles(role_name) VALUES('React developer');
INSERT INTO roles(role_name) VALUES('Angular developer');
INSERT INTO roles(role_name) VALUES('Admin');
INSERT INTO roles(role_name) VALUES('Staff');

INSERT INTO users(username, password, salt, full_name, role, created_at, updated_at) VALUES('admingoldenowl', '$argon2id$v=19$m=65536,t=3,p=2$ABq2zu0OCaM$C/n98UGEgOMhSt1FCcpxqJkvt5EOh/UBZQ/veR0gKE8', 'ABq2zu0OCaM', 'Phan Tran Khanh Hung', 'Admin', current_timestamp, current_timestamp);
INSERT INTO users(username, password, salt, full_name, role, created_at, updated_at) VALUES('hungintern', '$argon2id$v=19$m=65536,t=3,p=2$ABq2zu0OCaM$C/n98UGEgOMhSt1FCcpxqJkvt5EOh/UBZQ/veR0gKE8', 'ABq2zu0OCaM', 'Phan Tran Khanh Hung', 'Golang developer', current_timestamp, current_timestamp);
INSERT INTO users(username, password, salt, full_name, role, created_at, updated_at) VALUES('hungintern1', '$argon2id$v=19$m=65536,t=3,p=2$ABq2zu0OCaM$C/n98UGEgOMhSt1FCcpxqJkvt5EOh/UBZQ/veR0gKE8', 'ABq2zu0OCaM', 'Phan Tran Khanh Hung', 'Nodejs developer', current_timestamp, current_timestamp);
INSERT INTO users(username, password, salt, full_name, role, created_at, updated_at) VALUES('hungintern2', '$argon2id$v=19$m=65536,t=3,p=2$ABq2zu0OCaM$C/n98UGEgOMhSt1FCcpxqJkvt5EOh/UBZQ/veR0gKE8', 'ABq2zu0OCaM', 'Phan Tran Khanh Hung', 'Nodejs developer', current_timestamp, current_timestamp);