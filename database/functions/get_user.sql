create or replace function main.get_user(
    _id uuid,
    OUT username text,
    OUT password text,
    OUT error jsonb)

    returns record
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
    where u.id = _id;

    if not found then
        error := jsonb_build_object('error', 'not found');
        return;
    end if;

    error := jsonb_build_object();

exception
    when others then
        error := jsonb_build_object('error', 'internal');

end;

$$;