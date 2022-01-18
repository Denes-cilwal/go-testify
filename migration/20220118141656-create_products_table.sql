
-- +migrate Up
CREATE TABLE IF NOT EXISTS `products` (
  `id` VARCHAR(28) NOT NULL, 
  `name` VARCHAR(20) NOT NULL,
  `description` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
-- +migrate Down

DROP TABLE `products`;