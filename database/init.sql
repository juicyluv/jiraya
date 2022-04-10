create schema if not exists main;

create table if not exists users(
    id uuid primary key,
    title text
);

create or replace function main.create_user(username varchar)
    returns varchar
as $$

begin

    return 'hello world';

end;

$$
language plpgsql;
