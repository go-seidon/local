
CREATE TABLE `oauth_client` (
  `id` VARCHAR(128) NOT NULL,
  `name` VARCHAR(128) NOT NULL,
  `client_id` VARCHAR(256) NOT NULL,
  `client_secret` TEXT NOT NULL,
  `created_at` BIGINT(20) NOT NULL,
  `updated_at` BIGINT(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE unique_client_id(`client_id`)
) ENGINE = InnoDB;
