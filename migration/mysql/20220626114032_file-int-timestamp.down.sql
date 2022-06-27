
ALTER TABLE `file`
  CHANGE `created_at` `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CHANGE `updated_at` `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CHANGE `deleted_at` `deleted_at` TIMESTAMP NULL DEFAULT NULL;

