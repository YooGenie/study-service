CREATE SCHEMA IF NOT EXISTS `study` DEFAULT CHARACTER SET utf8;
USE `study`;

create table menu
(
    id          bigint auto_increment
        primary key,
    name        varchar(20)  not null,
    price       int          not null,
    description varchar(200) null,
    created_at  datetime     not null,
    created_by  varchar(20)  not null,
    updated_at  datetime     not null,
    updated_by  varchar(50)  not null
);