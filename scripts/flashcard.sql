 CREATE DATABASE `ximu_flashcard` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_bin */
 
 CREATE TABLE `auth_learner` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(45) NOT NULL DEFAULT '',
  `password` varchar(45) NOT NULL  DEFAULT '',
  `register_time` timestamp NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `status` tinyint(3) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE `know_point` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `pid` int(11) NOT NULL DEFAULT '0',
  `inde` int(11) NOT NULL DEFAULT '0',
  `name` varchar(200) NOT NULL,
  `status` tinyint(4) DEFAULT '0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`uid`, `name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE `card_template` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  `content_a` varchar(4096) DEFAULT '',
  `content_b` varchar(4096) DEFAULT '',
  `data_demo` varchar(4096) DEFAULT '',
  `card_name_key` varchar(128) DEFAULT '',
  `uid` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) DEFAULT '0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE `card` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  `kid` int(11) DEFAULT '0' COMMENT '知识点ID',
  `uid` int(11) NOT NULL DEFAULT '0',
  `tid` int(11) DEFAULT '0' COMMENT '模板ID',
  `data` varchar(4096) DEFAULT '' COMMENT 'JSON',
  `progress` tinyint(4) DEFAULT '0',
  `next_learn_time` datetime DEFAULT NULL,
  `status` tinyint(4) DEFAULT '0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


-- CREATE TABLE `card_progress` (
--   `id` int(11) NOT NULL AUTO_INCREMENT,
--   `uid` int(11) NOT NULL DEFAULT '0',
--   `cid` int(11) NOT NULL DEFAULT '0',
--   `rid` int(11) NOT NULL DEFAULT '0' COMMENT '记忆模板ID',
--   `progress` tinyint(4) DEFAULT '0',
--   `next_learn_time` datetime DEFAULT NULL,
--   `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
--   PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;