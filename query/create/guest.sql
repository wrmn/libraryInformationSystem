CREATE TABLE IF NOT EXISTS `guest`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(32) NOT NULL,
    `gender` CHAR(1) NOT NULL,
    `address` VARCHAR(64) NOT NULL,
    `profession` VARCHAR(16) NOT NULL,
    `institution` VARCHAR(32) NOT NULL,
    PRIMARY KEY(`id`)
);