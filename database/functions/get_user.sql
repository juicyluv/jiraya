create or replace function main.get_user(
    _id uuid,
    OUT username text,
    OUT email text,
    OUT created_at timestamptz,
    OUT disabled_at timestamptz,
    OUT error jsonb default null::jsonb)

    returns record
    stable
    strict
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

$$;