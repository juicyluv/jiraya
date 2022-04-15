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

create table if not exists main.projects(
    id uuid primary key,
    title text not null,
    created_at timestamptz not null default now(),
    closed_at timestamptz,
    description text,
    icon_url text,
    creator_id uuid not null references main.users(id)
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
    if _contact_name = '' then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'contact name is empty'));
        return;
    end if;

    if _contact = '' then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'contact is empty'));
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
----------------------------------------------------------------------
create or replace function main.create_project(
    _title text,
    _description text,
    _creator_id uuid,
    _icon_url text = null,
    OUT id uuid,
    OUT error jsonb)

    returns record

    language plpgsql
    strict
    security definer
as
$$

begin
    if _title = '' then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'title is empty'));
        return;
    end if;

    if _description = '' then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'description is empty'));
        return;
    end if;

    if not exists (select 1
                   from main.users u
                   where u.id = _creator_id)
    then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'user not found'));
        return;
    end if;

    id := gen_random_uuid();

    insert into main.projects(id, title, description, creator_id, icon_url)
    values (id, _title, _description, _creator_id, _icon_url);

    error := jsonb_build_object('code', 0);

exception
    when others then
        id := null;
        error := jsonb_build_object('code', -1);

end;

$$;
----------------------------------------------------------------------
create or replace function main.get_project(
    _project_id uuid,
    OUT title text,
    OUT description text,
    OUT created_at timestamptz,
    OUT closed_at timestamptz,
    OUT icon_url text,
    OUT creator_id uuid,
    OUT error jsonb)

    returns record
    stable
    language plpgsql
as
$$
begin
    select
        p.title, p.description, p.created_at, p.closed_at, p.icon_url, p.creator_id
    into
        title, description, created_at, closed_at, icon_url, creator_id
    from main.projects p
    where p.id = _project_id;

    if not found then
        error := jsonb_build_object('code', 1, 'details', jsonb_build_object('msg', 'project not found'));
    end if;

    error := jsonb_build_object('code', 0);

exception
    when others then
        error := jsonb_build_object('code', -1);
end;
$$;
----------------------------------------------------------------------
create or replace function main.get_projects(
    _user_id uuid default null::uuid,
    _closed bool default null::bool,
    _count int = 60,
    _sort_field text = 'created_at',
    _sort_order int = 1,
    _page int = 1,
    OUT project_id uuid,
    OUT title text,
    OUT description text,
    OUT created_at timestamptz,
    OUT closed_at timestamptz,
    OUT icon_url text,
    OUT creator_id uuid)

    returns setof record
    stable
    language plpgsql
as
$$
declare
    _sqlstr text;
begin
    _sqlstr := 'SELECT
                    p.id,
                    p.title,
                    p.description,
                    p.created_at,
                    p.closed_at,
                    p.icon_url,
                    p.creator_id
                 FROM main.projects p
                 WHERE ' ||

                case when _user_id is not null then ' p.creator_id = $1 ' else '' end ||
                case when _closed is not null then ' AND p.closed_at is null ' else '' end ||
                ' ORDER BY ' || quote_ident(trim(lower(_sort_field))) ||
                case when _sort_order = 1 then 'ASC' else 'DESC NULLS LAST ' end ||
                case when _count = 0 then '' else 'LIMIT $2 OFFSET $2*($3 - 1)' end;

    return query execute _sqlstr
    using _user_id, _count, _page;

exception
    when others then
        return query with t as (values(null::uuid,
                                       null::text,
                                       null::text,
                                       null::timestamptz,
                                       null::timestamptz,
                                       null::text,
                                       null::uuid))
                     select *
                     from t
                     where 1 = 2;
end;
$$;
----------------------------------------------------------------------
select error from main.create_user(_username := 'test', _email := 'test@test.com', _password := 'qwerty');