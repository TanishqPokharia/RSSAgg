-- +goose Up

alter table users add constraint unique_name unique (name);

-- +goose Down

alter table users drop constraint unique_name;