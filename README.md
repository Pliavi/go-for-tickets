# Go For Tickets

## Description

A ticketing buying api that I'm using to learn Go.

## Tasks

### GEN 1

- [ ] Simple buying
  - [x] Pre-fill concert with limited number of tickets
  - [x] Show concert data
    - [x] Simple name filter
  - [ ] Create user by email(magic login)
  - [ ] Resolve buying ticket(fake)
    - [ ] Deduct tickets bought from concert
    - [ ] Add ticket to user's ticket list

### GEN 1+

- [ ] Better routing
  - [ ] Using Chi as router
    - [ ] Use correct methods for resource controllers
    - [ ] Use and take values from url wildcards  
           eg. /concerts/:id -> /concerts/1234567890
          .... /concerts/{name:[a-z-]+} -> /concerts/the-opera

### GEN 2

- [ ] Create concert
- [ ] More advanced concert search filters
- [ ] Create ticket
  - [ ] Add transfer ticket

### GEN 3

- [ ] Checkin ticket
  - [ ] Add QRCode to the Ticket
  - [ ] Validate the ticket
- [ ] Add real payment process

### GEN 4

- [ ] Add queue to high load concerts

### GEN 5

- [ ] Add more user authentication methods
- [ ] Add Pix payment

### GEN 6

- [ ] Add Criptocurrency payment
