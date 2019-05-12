create table accounts (
  id uuid not null primary key,
  create_time timestamptz not null,
  update_time timestamptz not null,
  delete_time timestamptz
);
