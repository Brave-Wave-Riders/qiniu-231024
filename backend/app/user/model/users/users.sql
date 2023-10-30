CREATE TABLE `users`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3)  DEFAULT NULL,
    `updated_at` datetime(3)  DEFAULT NULL,
    `deleted_at` datetime(3)  DEFAULT NULL,
    `username`   varchar(20)         NOT NULL COMMENT '用户名',
    `avatar`     varchar(255) DEFAULT NULL COMMENT '头像',
    `email`      varchar(40)         NOT NULL COMMENT '邮箱',
    `password`   varchar(128)        NOT NULL COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_users_email` (`email`),
    KEY `idx_users_deleted_at` (`deleted_at`),
    KEY `idx_users_username` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8

