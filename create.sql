CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS concerts;
DROP TABLE IF EXISTS concert_queues;
DROP TABLE IF EXISTS customers;

CREATE TABLE concerts (
  id SERIAL PRIMARY KEY,
  name TEXT,
  total_tickets INTEGER,
  concurrent_customer_limit INTEGER
);

CREATE TABLE customers (
  -- id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE
);

CREATE TABLE concert_queues (
  id SERIAL PRIMARY KEY,
  purchase_deadline TIMESTAMP, -- NULL means in queue, non-NULL means booking deadline
  concert_id INTEGER NOT NULL,
  -- customer_id UUID NOT NULL,
  customer_id INTEGER NOT NULL,
  FOREIGN KEY (concert_id) REFERENCES concerts(id) ON DELETE CASCADE,
  FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

INSERT 
  INTO concerts (name, total_tickets, concurrent_customer_limit) 
  VALUES ('The Grand Ole Opry', 100, 10);