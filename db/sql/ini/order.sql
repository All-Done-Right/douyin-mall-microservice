CREATE TABLE `order`
(
     `id` bigint unsigned AUTO_INCREMENT,
     `created_at` datetime(3) NULL,
     `updated_at` datetime(3) NULL,
     `deleted_at` datetime(3) NULL,
     `order_id` varchar(100),
     `user_id` int(11),
     `user_currency` varchar(10),
     `email` longtext,
     `street_address` longtext,
     `city` longtext,
     `state` longtext,
     `country` longtext,
     `zip_code` int,
     PRIMARY KEY (`id`),
     INDEX `idx_order_deleted_at` (`deleted_at`),
     UNIQUE INDEX `idx_order_order_id` (`order_id`)
);
CREATE TABLE `order_item`
(`id` bigint unsigned AUTO_INCREMENT,
 `created_at` datetime(3) NULL,
 `updated_at` datetime(3) NULL,
 `deleted_at` datetime(3) NULL,
 `order_id` varchar(100),
 `product_id` int(11),
 `quantity` int(11),
 `cost` decimal(10,2),
 PRIMARY KEY (`id`),
 INDEX `idx_order_item_deleted_at` (`deleted_at`),
 CONSTRAINT `fk_order_order_items` FOREIGN KEY (`order_id`) REFERENCES `order`(`order_id`))

