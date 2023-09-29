CREATE TABLE balances (
  id   varchar(255) NOT NULL PRIMARY KEY,
  account_id varchar(255) NOT NULL,
  amount int NOT NULL,
  last_update  date NOT NULL
);