CREATE TABLE `user`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(255)  DEFAULT NULL,
    `email`      varchar(255)  DEFAULT NULL,
    `age`        int DEFAULT NULL,
    `birthday`   timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;