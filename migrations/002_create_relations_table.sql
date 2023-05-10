CREATE TABLE `relation` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`url` TEXT NOT NULL COLLATE 'utf8mb4_general_ci',
	`code` VARCHAR(6) NOT NULL COLLATE 'utf8mb4_general_ci',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `code` (`code`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;