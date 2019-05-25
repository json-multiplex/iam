create table saml_users (
  account_id uuid not null references accounts(id),
  identity_provider_id varchar not null,
  id varchar not null,
  create_time timestamptz not null,
  update_time timestamptz not null,

  primary key (account_id, identity_provider_id, id),
  foreign key (account_id, identity_provider_id) references identity_providers(account_id, id)
);
