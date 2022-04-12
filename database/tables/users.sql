create table if not exists users(
    id uuid primary key,
    username text not null unique,
    password text not null,
    created_at timestamptz default now(),
    disabled_at timestamptz
);