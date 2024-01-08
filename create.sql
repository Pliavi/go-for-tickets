CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS concerts;
DROP TABLE IF EXISTS concert_queues;
DROP TABLE IF EXISTS customers;

CREATE TABLE concerts (
  id SERIAL PRIMARY KEY,
  name TEXT,
  booking_size INTEGER
);

CREATE TABLE customers (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4()
);

CREATE TABLE concert_queues (
  id SERIAL PRIMARY KEY,
  purchase_deadline TIMESTAMP, -- NULL means in queue, non-NULL means booking deadline
  concert_id INTEGER NOT NULL,
  customer_id UUID NOT NULL,
  FOREIGN KEY (concert_id) REFERENCES concerts(id) ON DELETE CASCADE,
  FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);
