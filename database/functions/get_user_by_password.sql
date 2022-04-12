create or replace function main.get_user_by_password(
    _username text,
    _password text,
    OUT id uuid,
    OUT email text,
    OUT created_at timestamptz,
    OUT disabled_at timestamptz,
    OUT error jsonb default null::jsonb)
    returns record
as
$$

begin
    select
        u.id,
        u.email,
        u.created_at,
        u.disabled_at
    into
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

    if disabled_at is not null then
        error := jsonb_build_object('error', 'disabled');
        return;
    end if;

exception
    when others then
        error := jsonb_build_object('error', 'internal');

end

$$ language plpgsql stable security definer;