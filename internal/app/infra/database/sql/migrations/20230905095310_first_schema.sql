
-- +goose Up
CREATE TABLE IF NOT EXISTS product(
  id uuid PRIMARY KEY,
  name varchar(255) NOT NULL,
  price float NOT NULL,
  fornecedor varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  brand varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL
);


