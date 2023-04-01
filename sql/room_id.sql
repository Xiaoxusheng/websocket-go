-- auto-generated definition
create table room_id
(
    roomidently varchar(16) not null comment '房间唯一标识',
    useridently varchar(36) not null comment '用户唯一标识',
    room_type   varchar(10) not null comment '房间类型',
    creaet_time int         not null comment '创建时间'
        primary key,
    create_uesr varchar(36) not null comment '创建者',
    info        varchar(45) not null,
    constraint room_id_creaet_time_uindex
        unique (creaet_time)
)
    comment '创建房间表';


