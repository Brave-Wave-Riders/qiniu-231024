create table fans
(
    id      bigint auto_increment
        primary key,
    user_id bigint not null comment '用户id',
    fan_id  bigint not null comment '粉丝的id'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

create index idx_fan
    on fans (fan_id);

create index idx_user
    on fans (user_id);