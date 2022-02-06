CREATE TABLE IF NOT EXISTS `visitor`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `userId` INT,
    `guestId` INT,
    `loginAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `method` CHAR(1) NOT NULL,
    `purpose` VARCHAR(32) NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`userId`) REFERENCES user(id),
    FOREIGN KEY(`guestId`) REFERENCES guest(id)
);