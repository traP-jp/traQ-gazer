CREATE TABLE `words` (
    `user_name` VARCHAR(32) NOT NULL,
    `word` VARCHAR(50) NOT NULL,
    `register_time` DATETIME DEFAULT NOW(),
    `bot_notification` BOOLEAN,
    `me_notification` BOOLEAN,
    PRIMARY KEY (`user_name`, `word`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- TODO:FOREIGN KEY for user_name
CREATE TABLE `users` (
    `user_name` VARCHAR(32) NOT NULL PRIMARY KEY,
    `traq_uuid` VARCHAR(36) NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
