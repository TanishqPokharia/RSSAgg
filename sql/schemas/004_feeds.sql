-- +goose Up

create table feeds (
    id uuid primary key default gen_random_uuid(),
    title text not null,
    url text ,
    user_id uuid not null references users(id) on delete cascade
);

-- +goose Down

drop table feeds;