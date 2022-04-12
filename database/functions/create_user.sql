create or replace function main.create_user(
    _username text,
    _password text,
    OUT id uuid,
    OUT error jsonb)

    returns record
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

exception
    when others then
        error := jsonb_build_object('error', 'internal');
end;

$$;