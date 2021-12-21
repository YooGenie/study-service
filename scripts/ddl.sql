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


create table store
(
    no int auto_increment
        primary key,
    id varchar(20) not null,
    password varchar(100) not null,
    business_registration_number VARCHAR(10) not null comment '사업자번호',
    mobile VARCHAR(20) not null comment '핸드폰번호',
    created JSON not null,
    updated JSON not null,
    deleted_at datetime null,
    UNIQUE INDEX `uq_store_business_registration_number` (`business_registration_number`),
    UNIQUE INDEX `uq_store_id` (`id`)
);

create table store_info
(
    id int auto_increment,
    store_id varchar(20) not null,
    business_registration_number VARCHAR(10) not null comment '사업자번호',
    Business_report_no varchar(20) not null comment '영업신고증고유번호',
    representative varchar(20) not null comment '대표자명',
    store_name varchar(50) not null comment '영업소명칭',
    post_no varchar(5) not null comment '우편번호',
    address varchar(50) not null comment '주소',
    address_detail varchar(50) not null,
    business_type varchar(50) not null comment '영업의종류',
    created json not null,
    updated json null,
    deleted_at datetime null,
    constraint store_info_pk
        primary key (id),
    UNIQUE INDEX `uq_store_info_store_id` (`store_id`)
);