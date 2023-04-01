-- auto-generated definition
create table message
(
    idently          varchar(36) not null,
    message_id       int         not null,
    room_idently     varchar(16) not null,
    messagesend_time int         not null,
    message          varchar(45) not null,
    id               int auto_increment
        primary key,
    constraint message_id_uindex
        unique (id),
    constraint message_message_id_uindex
        unique (message_id)
)
    comment '消息表';
