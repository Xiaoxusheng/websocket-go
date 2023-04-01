-- auto-generated definition
create table user_room
(
    useridently   varchar(36) not null,
    roomidently   varchar(16) not null,
    create_time   int         not null,
    update_time   int         not null,
    room_type     varchar(16) not null,
    id            int auto_increment
        primary key,
    friendidently varchar(36) not null,
    constraint user_room_id_uindex
        unique (id)
)
    comment '用户房间关联表';

