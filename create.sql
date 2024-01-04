drop table if exists concert;
drop table if exists concert_queue;
drop table if exists customer;

create table concert (
  id integer primary key auto_increment,
  name text,
  booking_size integer,
)

create table customer (
  id text primary key, -- UUID
)

create table concert_queue (
  id integer primary key auto_increment,
  -- if null, means customer is in queue
  -- if deadline, means customer is booking
  purchase_deadline datetime, 

  concert_id integer not null,
  customer_id text not null,

  foreign key(concert_id) references concert(id)
  foreign key(customer_id) references customer(id)
)



