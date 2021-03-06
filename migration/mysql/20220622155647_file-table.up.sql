
CREATE TABLE `file` (
  `id` VARCHAR(128) NOT NULL,
  `name` VARCHAR(4096) NOT NULL,
  `path` TEXT NOT NULL,
  `mimetype` VARCHAR(256) NOT NULL,
  `extension` VARCHAR(128) NOT NULL,
  `size` BIGINT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) 
DEFAULT CHARACTER SET utf8
COLLATE utf8_unicode_ci
ENGINE = InnoDB;
