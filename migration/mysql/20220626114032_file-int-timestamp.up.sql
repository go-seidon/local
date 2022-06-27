
ALTER TABLE `file` 
  CHANGE `created_at` `created_at` BIGINT NOT NULL,
  CHANGE `updated_at` `updated_at` BIGINT NOT NULL,
  CHANGE `deleted_at` `deleted_at` BIGINT NULL DEFAULT NULL;
