create schema if not exists main;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

create table if not exists users(
    id uuid primary key,
    username text not null unique,
    password text not null,
    created_at timestamptz default now(),
    disabled_at timestamptz
);

create or replace function main.create_user(_username text, _password text)
    returns uuid
    language plpgsql
as
$$

declare
    id uuid;
    hashed_password text;

begin
    id = gen_random_uuid();
    hashed_password = crypt(_password, gen_salt('bf', 8));

    insert into users(id, username, password)
        values(id, _username, hashed_password);

    return id;
end;

$$;

create or replace function main.get_user(_id uuid, _username text default null::text, OUT username text, OUT password text, OUT error jsonb) returns record
    stable
    strict
    language plpgsql
as
$$
begin

    select
        u.username,
        u.password
    into
        username,
        password
    from users u
    where u.id = _id or u.username = _username;

    if not found then
        error := jsonb_build_object('code', 1, 'detail', json_build_object('user_id', 5));
        return;
    end if;

    error := jsonb_build_object('code', 0);

    exception
        when others then
            error := jsonb_build_object('code', -1);

end;
$$;

create or replace function main.get_user_by_password(_username text, _password text, OUT id uuid, OUT username text, OUT created_at timestamptz, OUT error jsonb) returns record

as
$$
begin
    select
        u.id,
        u.username,
        u.created_at
    into
        id,
        username,
        created_at
    from users u
    where
        u.username = lower(_username)
        and u.password = crypt(_password, u.password)
        and (u.disabled_at is null);

    if not found then
        error := jsonb_build_object('code', 1, 'detail', jsonb_build_object('user_id', 5));
        return;
    end if;

exception
    when others then

        error := jsonb_build_object('code', -1);

end
$$ language plpgsql stable security definer;