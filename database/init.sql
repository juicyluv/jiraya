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

create table if not exists main.users(
    id uuid primary key,
    username text not null unique check (username != ''),
    email text not null unique check (email != ''),
    password text not null check (password != ''),
    created_at timestamptz default now(),
    disabled_at timestamptz
);

create table if not exists main.user_contacts(
    id uuid primary key,
    user_id uuid references main.users(id),
    contact_name text not null,
    contact text not null check (contact != '')
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

    language plpgsql
    strict
    security definer
as
$$

declare
    hashed_password text;
    existsUsername text;
    existsEmail text;
begin
    select
        u.username,
        u.email
    into
        existsUsername,
        existsEmail
    from main.users u
    where u.username = _username or u.email = _email;

    if found then
        if existsUsername = _username then
            error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'username already taken'));
            return;
        end if;

        if existsEmail = _email then
            error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'email already taken'));
            return;
        end if;
    end if;

    id := gen_random_uuid();

    hashed_password := crypt(_password, gen_salt('bf', 8));

    insert into main.users(id, username, email, password)
        values(id, lower(_username), _email, hashed_password);

    error := jsonb_build_object('code', 0);

    exception
        when others then
            id := null;
            error := jsonb_build_object('code', -1);
end;

$$;

----------------------------------------------------------------------

create or replace function main.get_user(
    _id uuid,
    OUT username text,
    OUT email text,
    OUT created_at timestamptz,
    OUT disabled_at timestamptz,
    OUT error jsonb)

    returns record

    language plpgsql
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
    from main.users u
    where u.id = _id;

    if not found then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'user not found'));
        return;
    end if;

    error := jsonb_build_object('code', 0);

    exception
        when others then
            error := jsonb_build_object('code', -1);

end;

$$;

----------------------------------------------------------------------

create or replace function main.get_user_by_password(
    _login text,
    _password text,
    OUT id uuid,
    OUT username text,
    OUT email text,
    OUT created_at timestamptz,
    OUT disabled_at timestamptz,
    OUT error jsonb)

    returns record

    language plpgsql
    stable
    security definer
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
    from main.users u
    where
        (u.username = lower(_login) or u.email = lower(_login))
        and u.password = crypt(_password, u.password);

    if not found then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'user not found'));
        return;
    end if;

    error := jsonb_build_object('code', 0);

exception
    when others then
        error := jsonb_build_object('code', -1);

end;

$$;

----------------------------------------------------------------------

create or replace function main.create_user_contact(
    _user_id uuid,
    _contact_name text,
    _contact text,
    OUT id uuid,
    OUT error jsonb)

    returns record

    language plpgsql
    strict
    security definer
as
$$

begin
    if _contact_name = '' or _contact = '' then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'contact or contact name are empty'));
        return;
    end if;

    if not exists (select 1
                   from main.users u
                   where u.id = _user_id)
    then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'user not found'));
        return;
    end if;

    if exists(select 1 from main.user_contacts
              where contact_name = _contact_name)
    then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'contact already exists'));
        return;
    end if;

    id := gen_random_uuid();

    insert into main.user_contacts(id, user_id, contact_name, contact)
    values (id, _user_id, _contact_name, _contact);

    error := jsonb_build_object('code', 0);

exception
    when others then
        id := null;
        error := jsonb_build_object('code', -1);

end;

$$;

----------------------------------------------------------------------
create or replace function main.get_user_contact(
    _contact_id uuid,
    OUT user_id uuid,
    OUT contact_name text,
    OUT contact text,
    OUT error jsonb)

    returns record
    stable
    language plpgsql
as
$$
    begin
        select
               c.user_id, c.contact_name, c.contact
        into
            user_id, contact_name, contact
        from main.user_contacts c
        where c.id = _contact_id;

        error := jsonb_build_object('code', 0);

    exception
        when others then
            error := jsonb_build_object('code', -1);
    end;
$$;
----------------------------------------------------------------------
create or replace function main.get_user_contacts(
    _user_id uuid,
    OUT contact_id uuid,
    OUT contact_name text,
    OUT contact text)

    returns setof record
    stable
    language plpgsql
as
$$
begin
    return query select
        c.id, c.contact_name, c.contact
    from main.user_contacts c
    where c.user_id = _user_id;

exception
    when others then
        return query with t as (values(null::uuid,
                                       null::text,
                                       null::text))
        select *
        from t
        where 1 = 2;
end;
$$;
----------------------------------------------------------------------
create or replace function main.update_user_contact(
    _contact_id uuid,
    _contact_name text default null::text,
    _contact text default null::text)

    returns jsonb
    strict
    language plpgsql
as
$$
    declare
        _sqlstr text;
    begin
        if not exists(
            select 1 from main.user_contacts uc where uc.id = _contact_id
            )
        then
            return jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'contact not found'));
        end if;

        _sqlstr :=  case when _contact_name is null then '' else ' contact_name = $1,' end ||
                    case when _contact is null then '' else ' contact = $2,' end;

        if _sqlstr = '' then
            return jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'invalid arguments'));
        end if;

        execute _sqlstr;

        return jsonb_build_object('code', 0);

    exception
        when others then
            return jsonb_build_object('code', -1);
    end;
$$;
----------------------------------------------------------------------
create or replace function main.delete_user_contact(_contact_id uuid)
    returns jsonb
    strict
    language plpgsql
as
$$
begin
    if not exists(
            select 1 from main.user_contacts uc where uc.id = _contact_id
        )
    then
        return jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'contact not found'));
    end if;

    delete from main.user_contacts uc where uc.id = _contact_id;

    return jsonb_build_object('code', 0);

exception
    when others then
        return jsonb_build_object('code', -1);
end;
$$;
----------------------------------------------------------------------
select error from main.create_user(_username := 'test', _email := 'test@test.com', _password := 'qwerty');