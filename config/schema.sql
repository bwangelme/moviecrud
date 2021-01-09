create table movie
(
    id         int unsigned not null auto_increment primary key,
    title      varchar(64)  not null,
    pubdate    datetime     not null,
    country    varchar(64)  not null,

    created_at timestamp default current_timestamp,
    update_at  timestamp default current_timestamp on update current_timestamp
);