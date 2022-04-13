revoke all on database postgres, template0, template1 from public;

create user jiraya;

create database jiraya;

grant connect on database jiraya to jiraya;

\connect jiraya

revoke all on schema public from public;

create schema if not exists main authorization postgres;
grant all on schema main to postgres;
grant usage on schema main to jiraya;


----------------------------------------------------------------------

-------------------------------EXTENSIONS-----------------------------

----------------------------------------------------------------------

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

----------------------------------------------------------------------

-------------------------------TABLES---------------------------------

----------------------------------------------------------------------

create table if not exists users(
    id uuid primary key,
    username text not null unique,
    email text not null unique,
    password text not null,
    created_at timestamptz default now(),
    disabled_at timestamptz
);


----------------------------------------------------------------------

-------------------------------FUNCTIONS------------------------------

----------------------------------------------------------------------

create or replace function main.create_user(
    _username text,
    _password text,
    _email text,
    OUT id uuid,
    OUT error jsonb)

    returns record
as
$$

declare
    hashed_password text;
begin
    id = gen_random_uuid();

    hashed_password = crypt(_password, gen_salt('bf', 8));

    insert into users(id, username, email, password)
        values(id, lower(_username), _email, hashed_password);

    error := jsonb_build_object();

    exception
        when others then
            id = null::uuid;
            error := jsonb_build_object('error', 'internal');
end;

$$
language plpgsql;

----------------------------------------------------------------------

create or replace function main.get_user(
    _id uuid,
    OUT username text,
    OUT email text,
    OUT created_at timestamptz,
    OUT disabled_at timestamptz,
    OUT error jsonb)

    returns record
as
$$

begin

    select
        u.username,
        u.email,
        u.created_at,
        u.disabled_at
    into
        username,
        email,
        created_at,
        disabled_at
    from users u
    where u.id = _id;

    if not found then
        error := jsonb_build_object('error', 'not found');
        return;
    end if;

    exception
        when others then
            error := jsonb_build_object('error', 'internal');

end;

$$
language plpgsql;

----------------------------------------------------------------------

create or replace function main.get_user_by_password(
    _username text,
    _password text,
    OUT id uuid,
    OUT username text,
    OUT email text,
    OUT created_at timestamptz,
    OUT disabled_at timestamptz,
    OUT error jsonb)
returns record
as
$$

begin
    select
        u.id,
        u.username,
        u.email,
        u.created_at,
        u.disabled_at
    into
        id,
        username,
        email,
        created_at,
        disabled_at
    from users u
    where
        u.username = lower(_username)
        and u.password = crypt(_password, u.password);

    if not found then
        error := jsonb_build_object('error', 'not found');
        return;
    end if;

exception
    when others then
        error := jsonb_build_object('error', 'internal');

end

$$ language plpgsql stable security definer;

----------------------------------------------------------------------