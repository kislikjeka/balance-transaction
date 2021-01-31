create table if not exists balances
(
    id serial not null unique,
    account_id integer not null,
    balance float,
    created_at timestamp,
    updated_at timestamp
)