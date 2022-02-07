CREATE TABLE IF NOT EXISTS `inventory`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `registrationId` INT NOT NULL,
    `serialNumber` INT NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `category` CHAR(1) NOT NULL,
    `status` CHAR(1) NOT NULL,
    `description` VARCHAR(256) NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`registrationId`) REFERENCES assetRecord(`id`)
);