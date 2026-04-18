CREATE TABLE accounts (
  id UUID PRIMARY KEY,
  balance BIGINT
);

CREATE TABLE transactions (
  id UUID PRIMARY KEY,
  from_account UUID,
  to_account UUID,
  amount BIGINT,
  status TEXT
);
