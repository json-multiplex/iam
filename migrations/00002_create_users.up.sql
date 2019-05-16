create table users (
  account_id uuid not null references accounts(id),
  id varchar not null,
  create_time timestamptz not null,
  update_time timestamptz not null,
  delete_time timestamptz,
  is_root boolean not null,
  password_hash varchar,

  primary key (account_id, id)
);
