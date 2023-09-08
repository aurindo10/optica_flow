

-- +goose Up
CREATE TABLE IF NOT EXISTS fornecedor (
  id uuid,
  name varchar(255) NOT NULL,
  telefone varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  adress varchar(255) NOT NULL,
  company_id varchar(255) NOT NULL,
  who_created_id varchar(255) NOT NULL,
  who_updated_id varchar(255) NOT NULL,
  cnpj varchar(255) NOT NULL,
   PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS product(
  id uuid PRIMARY KEY,
  name varchar(255) NOT NULL,
  price float NOT NULL,
  fornecedor_id uuid REFERENCES fornecedor(id) ON DELETE SET NULL,
  CONSTRAINT fk_fornecedor FOREIGN KEY (fornecedor_id) REFERENCES fornecedor(id),
  description varchar(255) NOT NULL,
  brand varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL,
  bar_code varchar(255) NOT NULL,
  quantity int NOT NULL,
  company_id varchar(255) NOT NULL,
  who_created_id varchar(255) NOT NULL,
  who_updated_id varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS client (
  id uuid PRIMARY KEY,
  full_name varchar(255) NOT NULL,
  telefone varchar(255) NOT NULL,
  cpf varchar(255),
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL,
  email varchar(255) NOT NULL,
  birth_date timestamp NOT NULL,
  adress varchar(255),
  gender varchar(255),
  city varchar(255),
  seller_id varchar(255),
  company_id varchar(255) NOT NULL,
  who_created_id varchar(255) NOT NULL,
  who_updated_id varchar(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS orders (
  id   uuid PRIMARY KEY,
  product_name varchar(255) NOT NULL,
  quantity int NOT NULL,
  order_date timestamp NOT NULL DEFAULT current_timestamp,
  who_created_id varchar(255) NOT NULL,
  who_updated_id varchar(255) NOT NULL,
  client_id uuid REFERENCES client(id) ON DELETE CASCADE,
  CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES client(id),
  company_id varchar(255) NOT NULL,
  fase varchar(255) NOT NULL
    
);

CREATE TABLE IF NOT EXISTS product_order (
  id uuid PRIMARY KEY,
  amout int NOT NULL,
  product_id uuid REFERENCES product(id) ON DELETE CASCADE,
  order_id uuid REFERENCES orders(id) ON DELETE CASCADE,
  CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product(id),
  CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id) 
);

CREATE TABLE IF NOT EXISTS points(
  id uuid PRIMARY KEY,
  name varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  active boolean NOT NULL,
  ammount float NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL,
  valid_date timestamp NOT NULL,
  company_id varchar(255) NOT NULL,
  order_id uuid REFERENCES orders(id) ON DELETE CASCADE,
  CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS trade_product(
  id uuid PRIMARY KEY,
  name varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL,
  company_id varchar(255) NOT NULL,
  point_ammount int NOT NULL,
  image_url varchar(255),
  who_created_id varchar(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS commission(
  id uuid PRIMARY KEY,
  name varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL,
  company_id varchar(255) NOT NULL,
  who_created_id varchar(255) NOT NULL,
  who_updated_id varchar(255) NOT NULL,
  order_id uuid REFERENCES orders(id) ON DELETE CASCADE,
  CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id),
  value float NOT NULL
);


CREATE TABLE IF NOT EXISTS comission_values (
  id uuid PRIMARY KEY,
  percentage float NOT NULL,
  company_id varchar(255) NOT NULL,
  who_created_id varchar(255) NOT NULL,
  who_updated_id varchar(255) NOT NULL,
  type varchar(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS cash_flow_entries (
  id uuid PRIMARY KEY,
  date timestamp NOT NULL DEFAULT current_timestamp,
  type varchar(255) NOT NULL, -- 'REVENUE' ou 'EXPENSE'
  amount float NOT NULL,
  description varchar(255) NOT NULL,
  company_id varchar(255) NOT NULL,
  order_id uuid REFERENCES orders(id) ON DELETE CASCADE, -- para relacionar uma entrada de receita a um pedido
  CONSTRAINT fk_order_entries FOREIGN KEY (order_id) REFERENCES orders(id),
  who_created_id varchar(255) NOT NULL,
  who_updated_id varchar(255) NOT NULL

);


CREATE TABLE IF NOT EXISTS cash_flow_balance (
  id uuid PRIMARY KEY,
  date date NOT NULL, 
  company_id varchar(255) NOT NULL,
  who_created_id varchar(255) NOT NULL,
  who_updated_id varchar(255) NOT NULL,
  comission_id uuid REFERENCES commission(id) ON DELETE CASCADE,
  CONSTRAINT fk_comission FOREIGN KEY (comission_id) REFERENCES commission(id),
  due_date timestamp NOT NULL,
  paid_date timestamp NOT NULL,
  paid boolean NOT NULL,
  value float NOT NULL,
  type varchar(255) NOT NULL,
  description varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS goose_db_version (
  id serial primary key,
  version_id bigint,
  is_applied boolean,
  tstamp timestamp
);