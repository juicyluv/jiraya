create or replace function main.get_user_by_password(
    _username text,
    _password text,
    OUT id uuid,
    OUT username text,
    OUT created_at timestamptz,
    OUT error jsonb)
    returns record
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
        error := jsonb_build_object('error', 'not found');
        return;
    end if;

exception
    when others then

        error := jsonb_build_object('error', 'internal');

end

$$ language plpgsql stable security definer;