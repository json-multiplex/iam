create table access_keys (
  account_id uuid not null references accounts(id),
  user_id varchar not null,
  id varchar not null,
  create_time timestamptz not null,
  update_time timestamptz not null,
  delete_time timestamptz,
  secret_hash varchar not null,

  primary key (account_id, user_id, id),
  foreign key (account_id, user_id) references users(account_id, id)
);
