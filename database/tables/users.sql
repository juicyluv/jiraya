create table if not exists main.users(
    id uuid primary key,
    username text not null unique,
    email text not null unique,
    password text not null,
    created_at timestamptz default now(),
    disabled_at timestamptz
);