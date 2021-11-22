SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`email` varchar(50) UNIQUE NOT NULL,
	`password` varchar(50) NOT NULL,
	`salt` varchar(50) NOT NULL,
	`fb_id` varchar(50),
	`gg_id` varchar(50),
	`first_name` varchar(50),
	`last_name` varchar(50),
	`phone` varchar(20),
	`status` smallint unsigned NOT NULL DEFAULT '1',
	`role` ENUM ('user', 'admin') DEFAULT "user",
	`avatar` json,
	`birthday` date,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `settings`;

CREATE TABLE `settings` (
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`user_id` int,
	`language_code` ENUM ('en', 'vi') DEFAULT "en",
	`available_bars` varchar(50) DEFAULT "20",
	`available_plates` varchar(50) DEFAULT "20,15,10,5",
	`weight_unit` ENUM ('kg', 'ibs') DEFAULT "kg",
	`size_unit` ENUM ('cm', 'in') DEFAULT "cm",
	`theme` ENUM ('dark', 'light') DEFAULT "dark",
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	KEY `user_id` (`user_id`)
	USING BTREE
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `measurements`;

CREATE TABLE `measurements` (
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`user_id` int NOT NULL,
	`body_part` ENUM (
		'neck',
		'chest',
		'left_biceps',
		'right_biceps',
		'left_forearms',
		'right_forearms',
		'waist',
		'hips',
		'glutes',
		'left_thigh',
		'right_thigh',
		'left_calf',
		'right_calf'
	) NOT NULL,
	`value` int NOT NULL,
	`unit` varchar(255) NOT NULL,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	KEY `user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `exercises`;

CREATE TABLE `exercises` (
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`name` varchar(255) UNIQUE NOT NULL,
	`created_by` int NOT NULL,
	`category` ENUM (
		'barbell',
		'dumbbell',
		'machine/other',
		'weighted bodyweight',
		'assisted body',
		'reps only',
		'cardio exercise',
		'duration'
	) NOT NULL,
	`body_parts` varchar(255) NOT NULL COMMENT 'arms, core',
	`mechanics` ENUM ('isolation', 'compound'),
	`force` ENUM ('push', 'pull'),
	`rest_time` int DEFAULT 60000 COMMENT 'miliseconds',
	`instructions` varchar(255),
	`image` json COMMENT 'image object includes image size',
	`video` json COMMENT 'youtube link or any video link to explain the moves',
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `workouts`;

CREATE TABLE `workouts` (
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`name` varchar(50) UNIQUE NOT NULL,
	`note` varchar(255),
	`created_by` int NOT NULL,
	`template_id` int,
	`practice_date` timestamp DEFAULT CURRENT_TIMESTAMP,
	`duration` int NOT NULL,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `templates`;

CREATE TABLE `templates` (
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`name` varchar(50) NOT NULL,
	`note` varchar(255),
	`template` json NOT NULL,
	`created_by` int NOT NULL,
	`collection_id` int DEFAULT NULL,
	`is_archived` boolean DEFAULT FALSE,
	`is_favorite` boolean DEFAULT FALSE,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	KEY `name` (`name`) USING BTREE,
	KEY `collection_id` (`collection_id`) USING BTREE
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `records`;

CREATE TABLE `records` (
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`workout_id` int NOT NULL,
	`exercise_id` int NOT NULL,
	`set` varchar(255) NOT NULL,
	`reps` int NOT NULL,
	`order` int unsigned DEFAULT 1,
	`rest_time` int DEFAULT 0,
	`rpe` ENUM ('5', '6', '7', '8',	'9', '10') DEFAULT NULL,
	`weights` float4 unsigned DEFAULT 0,
	`unit` ENUM ('kg', 'ibs') DEFAULT "kg",
	`time` int,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	KEY `workout_id` (`workout_id`) USING BTREE,
	KEY `exercise_id` (`exercise_id`) USING BTREE
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `collections`;

CREATE TABLE `collections` (
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`parent_id` int DEFAULT NULL,
	`is_favorite` boolean DEFAULT FALSE,
	`is_archived` boolean DEFAULT FALSE,
	`created_by` int NOT NULL,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	KEY `parent_id` (`parent_id`) USING BTREE,
	KEY `name` (`name`) USING BTREE
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `password_reset_tokens`;

CREATE TABLE `password_reset_tokens` (
	`user_id` int NOT NULL,
	`token` varchar(128) UNIQUE NOT NULL,
	`token_expiry` int NOT NULL,
	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (user_id, token),
	KEY `user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `user_device_tokens`;

CREATE TABLE `user_device_tokens` (
	`id` int unsigned NOT NULL AUTO_INCREMENT,
	`user_id` int NOT NULL,
	`is_production` tinyint (1) DEFAULT '0',
	`os` enum ('ios', 'android', 'web') DEFAULT 'ios',
	`token` varchar(255) DEFAULT NULL,
	`status` smallint unsigned NOT NULL DEFAULT '1',
	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `user_id` (`user_id`) USING BTREE,
	KEY `os` (`os`) USING BTREE
) ENGINE = InnoDB;

ALTER TABLE `settings`
	ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `measurements`
	ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `exercises`
	ADD FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

ALTER TABLE `workouts`
	ADD FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

ALTER TABLE `workouts`
	ADD FOREIGN KEY (`template_id`) REFERENCES `templates` (`id`);

ALTER TABLE `templates`
	ADD FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

ALTER TABLE `templates`
	ADD FOREIGN KEY (`collection_id`) REFERENCES `collections` (`id`);

ALTER TABLE `records`
	ADD FOREIGN KEY (`workout_id`) REFERENCES `workouts` (`id`);

ALTER TABLE `records`
	ADD FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`);

ALTER TABLE `collections`
	ADD FOREIGN KEY (`id`) REFERENCES `collections` (`parent_id`);

ALTER TABLE `collections`
	ADD FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

ALTER TABLE `password_reset_tokens`
	ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_device_tokens`
	ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

SET FOREIGN_KEY_CHECKS = 1;
