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