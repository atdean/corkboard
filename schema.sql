USE corkboard;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL UNIQUE,
    `email` varchar(255) NOT NULL UNIQUE,
    `password` char(64) NOT NULL, -- long enough for SHA256 hash
    `reset_token` char(64) NOT NULL DEFAULT "",
    `reset_expiration` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `account_status` varchar(30) NOT NULL DEFAULT "active",
    `profile_photo_name` varchar(255) DEFAULT "",
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(11) UNSIGNED NOT NULL, -- references users.id
    `body` text NOT NULL DEFAULT "",
    `photo_name` varchar(255) NOT NULL DEFAULT "",
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
