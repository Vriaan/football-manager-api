DROP DATABASE IF EXISTS footballmanager_test;
CREATE DATABASE `footballmanager_test`;
use footballmanager_test;

CREATE TABLE `footballers` (
    `id` INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(50) COLLATE utf8_unicode_ci NOT NULL,
    `last_name` VARCHAR(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `teams` (
    `id` INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
    `number1` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 1',
    `number2` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 2',
    `number3` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 3',
    `number4` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 4',
    `number5` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 5',
    `number6` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 6',
    `number7` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 7',
    `number8` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 8',
    `number9` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 9',
    `number10` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 10',
    `number11` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 11',
    `number12` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 12 (substitute player)',
    `number13` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 13 (substitute player)',
    `number14` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Footballer ID being at football position number 14 (substitute player)',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `managers` (
    `id` INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `team_id` INT UNSIGNED NOT NULL,
    `first_name` VARCHAR(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
    `last_name` VARCHAR(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
    `email` VARCHAR(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
    `password` VARCHAR(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data Injection
INSERT INTO `footballers` (`id`, `first_name`, `last_name`)
VALUES
(1, 'Homer', 'Simpson'),
(2, 'Bart', 'Simpson'),
(3, 'Lisa', 'Simpson'),
(4, 'Carl', 'Carlson'),
(5, 'Lenny', 'Leonard'),
(6, 'Moe', 'Szyslak'),
(7, 'Patty', 'Bouvier'),
(8, 'Selma', 'Bouvier'),
(9, 'Edna', 'Krapabelle'),
(10, 'Seymour', 'Skinner'),
(11, 'Waylon', 'Smithers'),
(12, 'Charles Montgomery', 'Burns'),
(13, 'Thimoty', 'Lovejoy'),
(14, 'Apu', 'Nahasapeemapetilon');

INSERT INTO `teams` (`id`, `name`, `number1`, `number2`, `number3`, `number4`, `number5`, `number6`, `number7`, `number8`, `number9`, `number10`, `number11`, `number12`, `number13`,`number14`)
VALUES
(1, 'Springfield Spirit', 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14),
(2, 'Quahog Demolisher',  15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28);

-- For this test I let passwords there crytal clear.
INSERT INTO `managers` (`id`, `team_id`, `first_name`, `last_name`, `email`, `password`)
VALUES
(1, 1,  'Marge', 'Simpson', 'marge.simpson@gmail.com', 'ILoveMyFamily'),
(2, 2, 'Peter', 'Griffin', 'peter.griffin@gmail.com', 'Peter Griffin');

-- INSERT INTO `footballer_team` (`footballer_id`, `team_id`)
-- VALUES,
-- (1, 1),
-- (1, 2),
-- (1, 3),
-- (1, 4),
-- (1, 5),
-- (1, 6),
-- (1, 7),
-- (1, 8),
-- (1, 9),
-- (1, 10),
-- (1, 11),
-- (1, 12),
-- (1, 13),
-- (1, 14),
-- (1, 15),
-- (2, 2),
