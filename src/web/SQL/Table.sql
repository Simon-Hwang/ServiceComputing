CREATE DATABASE `test`;

CREATE TABLE `test`.`Article` (
  `Article_id` INT NOT NULL,
  `Article_name` VARCHAR(100) NOT NULL,
  `Article_content` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`Article_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE `test`.`User` (
  `username` VARCHAR(100) NOT NULL,
  `password` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`username`))
  ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE `test`.`Comment` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Content` VARCHAR(100) NOT NULL,
  `Comment_date` VARCHAR(100) NOT NULL,
  `Comment_publisher` VARCHAR(100) NOT NULL,
  `Article_id` INT NOT NULL,
  PRIMARY KEY (`id`))
  ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;