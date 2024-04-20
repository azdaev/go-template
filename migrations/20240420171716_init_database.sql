-- +goose Up
-- +goose StatementBegin
create table if not exists author (
    id serial primary key,
    fullname varchar(255) not null,
    biography varchar(255) not null,
    created_at timestamp default current_timestamp
);

create table if not exists playlist (
    id serial primary key,
    title varchar(255) not null,
    description varchar(511)
);

create table if not exists audiolecture (
    id serial primary key,
    title varchar(255) not null,
    author_id integer,
    playlist_id integer,
    description text,
    telegram_file_id varchar(511),
    relative_path varchar(511) not null,
    created_at timestamp default current_timestamp,

    foreign key (author_id) references author(id) on delete cascade,
    foreign key (playlist_id) references playlist(id)
);

create table if not exists category (
    id serial primary key,
    title varchar(31) not null
);

create table if not exists audiolecture_category_link (
    id serial primary key,
    audiolecture_id integer not null,
    category_id integer not null,

    foreign key (audiolecture_id) references audiolecture(id) on delete cascade,
    foreign key (category_id) references category(id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists audiolecture_category_link;
drop table if exists category;
drop table if exists audiolecture;
drop table if exists author;
drop table if exists playlist;
-- +goose StatementEnd
