create schema if not exists main;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists users(
    id uuid primary key,
    username text not null unique,
    password text not null,
    created_at timestamptz default now(),
    disabled_at timestamptz
);

create or replace function main.create_user(username text, password text)
    returns uuid
as $$

declare
    id uuid;
    hashed_password text;

begin
    id = uuid_generate_v4();
    hashed_password = md5(password);

    insert into users(id, username, password)
        values(id, username, hashed_password);

    return id;
end;

$$
language plpgsql;
