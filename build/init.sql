create table users (
    id uuid primary key,
    name text not null,
    password text
);

insert into users values ('f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b', 'Jerry', 'text1234')