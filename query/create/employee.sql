CREATE TABLE IF NOT EXISTS `employee`(
    `id` INT NOT NULL,
    `employeeNumber` CHAR(18) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `gender` CHAR(1) NOT NULL,
    `address1` VARCHAR(64) NOT NULL,
    `address2` VARCHAR(64),
    `division` CHAR(1),
    `position` VARCHAR(16),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`id`) REFERENCES user(`id`)
)