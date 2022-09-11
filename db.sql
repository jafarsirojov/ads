create table ad
(
    id              bigserial    not null primary key,
    title           varchar(200) not null,
    description     varchar(1000),
    price           bigint       not null,
    links_to_photos text[]       not null,
    created_at      timestamp with time zone default now()
);
