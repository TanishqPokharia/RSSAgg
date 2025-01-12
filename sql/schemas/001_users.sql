-- up migration (creation)

-- +goose Up

create table users (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null ,
    name text not null
);


-- down migration (deletion)

-- +goose Down

drop table users;