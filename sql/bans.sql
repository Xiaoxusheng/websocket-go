-- auto-generated definition
create table bans
(
    id   int auto_increment
        primary key,
    ip   varchar(45) not null comment '封杀ip',
    time timestamp   not null comment '封杀时间',
    constraint bans_id_uindex
        unique (id),
    constraint bans_ip_uindex
        unique (ip)
)
    comment 'ip黑名单';

