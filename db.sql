create table ad
(
    id          bigserial    not null primary key,
    title       varchar(200) not null,
    description varchar(1000),
    price       bigint       not null,
    created_at  timestamp with time zone default now()
);

create table photo
(
    id    bigserial    not null primary key,
    ad_id bigint       not null,
    url   varchar(500) not null,
    main  boolean      not null default false
);
