create or replace function main.create_user(
    _username text,
    _password text,
    _email text,
    OUT id uuid,
    OUT error jsonb default null::jsonb)

    returns record
    language plpgsql
as
$$

declare
    hashed_password text;
begin
    id = gen_random_uuid();

    hashed_password = crypt(_password, gen_salt('bf', 8));

    insert into users(id, username, email, password)
    values(id, lower(_username), _email, hashed_password);

exception
    when others then
        id = null::uuid;
        error := jsonb_build_object('error', 'internal');
end;

$$;