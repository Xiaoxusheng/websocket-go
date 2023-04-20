-- auto-generated definition
create table user
(
    username      varchar(45)   not null,
    indently      varchar(36)   not null
        primary key,
    password      varchar(45)   not null,
    email         varchar(45)   not null,
    register_time timestamp     not null,
    use_status    int default 0 not null comment '0 表示用户正常 1表示禁用',
    account       varchar(10)   not null,
    headpicture   varchar(45)   not null,
    constraint user_account_uindex
        unique (account),
    constraint user_email_uindex
        unique (email),
    constraint user_idently_uindex
        unique (indently),
    constraint user_register_time_uindex
        unique (register_time),
    constraint user_username_uindex
        unique (username)
) comment '用户表';

