-- auto-generated definition
create table video
(
    id             int auto_increment comment '作品的id'
        primary key,
    title          varchar(50)                         not null comment '作品标题',
    author_id      int                                 not null comment '作者的id',
    play_url       varchar(255)                        not null comment '视频资源的url',
    cover_url      varchar(255)                        not null comment '封面的url',
    favorite_count int       default 0                 not null comment '点赞数量',
    comment_count  int       default 0                 not null comment '评论数量',
    create_time    timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time    timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at     timestamp                           null comment '逻辑删除',
    kind           int                                 null comment '分类 '
)
    charset = utf8mb4;

create index author_id
    on video (author_id)
    comment '针对视频作者设置索引';

create index id
    on video (id, author_id)
    comment '针对id和作者id设置联合索引';

create index title
    on video (title)
    comment '针对视频标题添加索引';

