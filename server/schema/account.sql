CREATE TABLE `user`
(
    `user_id`      bigint NOT NULL,
    `role`         int          NOT NULL COMMENT 'role use in CMS',
    `user_name`    varchar(255) NOT NULL,
    `password`     varchar(255) NOT NULL,
    `full_name`    varchar(255) DEFAULT NULL,
    `email`        varchar(255) DEFAULT NULL,
    `gender`       int          DEFAULT NULL,
    `avatar`       varchar(255) DEFAULT NULL,
    `phone_number` varchar(255) DEFAULT NULL,
    -- `reference_id` varchar(255) DEFAULT NULL COMMENT 'id 3rd',
    -- `reference_type` int DEFAULT NULL COMMENT '1: google, 2: apple, 3: zalo, 4: facebook',
    `created_at`   bigint NOT NULL,
    `updated_at`   bigint    NOT NULL,
    `created_by`   bigint DEFAULT NULL,
    `updated_by`   bigint DEFAULT NULL,
    PRIMARY KEY (user_id),
    constraint user_user_name_uindex
        unique (user_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;