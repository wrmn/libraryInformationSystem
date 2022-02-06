CREATE TABLE IF NOT EXISTS `book`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `registrationId` INT NOT NULL,
    `serialNumber` INT NOT NULL,
    `ddc` CHAR(3) NOT NULL,
    `ddcOrder` INT NOT NULL,
    `title` VARCHAR(64) NOT NULL,
    `author` VARCHAR(32) NOT NULL,
    `publisher` VARCHAR(32) NOT NULL,
    `availability` BOOLEAN NOT NULL DEFAULT TRUE,
    `price` INT NOT NULL,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    FOREIGN key(`registrationId`) REFERENCES assetRecord(`id`),
    FOREIGN key(`ddc`) REFERENCES ddc(`ddc`)
);