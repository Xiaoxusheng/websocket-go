-- auto-generated definition
create table ip
(
    id          int auto_increment
        primary key,
    ip          varchar(16)  not null comment 'ip',
    time        int          null comment '访问时间',
    useindently varchar(200) not null,
    constraint IP_id_uindex
        unique (id)
)
    comment '访问者IP记录表';

