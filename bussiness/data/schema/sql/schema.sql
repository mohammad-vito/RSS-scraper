create table if not exists users
(
    id            bigserial
    primary key,
    name          text,
    email         text,
    roles         text[],
    password_hash bytea,
    date_created  timestamp with time zone,
    date_updated  timestamp with time zone
);

alter table users
    owner to postgres;

create table if not exists feeds
(
    id          bigserial
    primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    link        text,
    title       text,
    description text
);

alter table feeds
    owner to postgres;

create index if not exists idx_feeds_deleted_at
    on feeds (deleted_at);

create table if not exists posts
(
    id           bigserial
    primary key,
    created_at   timestamp with time zone,
    updated_at   timestamp with time zone,
    deleted_at   timestamp with time zone,
    link         text,
    title        text,
    author       text,
    content      text,
    feed_id      bigint
    constraint fk_feeds_posts
    references feeds,
    published_at timestamp not null
);

alter table posts
    owner to postgres;

create index if not exists idx_posts_deleted_at
    on posts (deleted_at);