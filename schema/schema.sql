CREATE TABLE `words` (
    `trap_id` VARCHAR(32) NOT NULL,
    -- Keep registered words binary-distinct; notification matching applies app-side normalization.
    `word` VARCHAR(50) NOT NULL COLLATE utf8mb4_bin,
    `register_time` DATETIME DEFAULT NOW(),
    `bot_notification` BOOLEAN,
    `me_notification` BOOLEAN,
    PRIMARY KEY (`trap_id`, `word`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- TODO:FOREIGN KEY for trap_id
CREATE TABLE `users` (
    `trap_id` VARCHAR(32) NOT NULL PRIMARY KEY,
    `traq_uuid` VARCHAR(36) NOT NULL,
    `is_bot` BOOLEAN
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE TABLE IF NOT EXISTS `pollinginfo` (
  `key` int NOT NULL,
  `lastpollingtime` datetime NOT NULL,
  PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
