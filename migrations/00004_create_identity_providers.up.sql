create table identity_providers (
  account_id uuid not null references accounts(id),
  id varchar not null,
  create_time timestamptz not null,
  update_time timestamptz not null,
  delete_time timestamptz,
  saml_metadata_url varchar not null,
  user_id_attribute varchar not null,

  primary key (account_id, id)
);
