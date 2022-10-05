CREATE DATABASE IF NOT EXISTS eps_info;
USE eps_info;
CREATE TABLE IF NOT EXISTS user_info
(
  `id`         int(11) NOT NULL AUTO_INCREMENT,
  `last_name`       text,
  `first_name`      text,
  `mail`      text,
  PRIMARY KEY (`id`)
);
INSERT INTO user_info (last_name,first_name,mail) VALUES ("yamada","hanako","hanako@mail.com");
INSERT INTO user_info (last_name,first_name,mail) VALUES ("yamamoto","tairo","tairo@mail.com");