 CREATE DATABASE `ximu_flashcard` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_bin */;

 use `ximu_flashcard`;
 
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



CREATE TABLE `word` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `word` varchar(64) DEFAULT '' COMMENT 'word',
  `content` varchar(8192) NOT NULL,
  `status` tinyint(4) DEFAULT '0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_word` (`word`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

INSERT INTO `card_template` VALUES 
(1,'问答题模板','<div style=\"position:relative; width:100%; height:100%;\">\n    <div style=\"position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);\">\n        {{Q}}\n    </div>\n</div>','<div style=\"position:relative; width:100%; height:100%;\">\n    <div style=\"position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);\">\n        {{A}}\n    </div>\n</div>','{\n   \"Q\": \"问题是什么\",\n    \"A\": \"答案是什么\"\n}','Q',1,0,'2020-06-15 19:29:21'),
(2,'选择题模板','<div style=\"position:relative; width:100%; height:100%;\">\n    <div style=\"position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);\">\n        <div>{{Q}}</div>\n        <div v-for=\"opt in Options\">{{opt}}</div>\n    </div>\n</div>','<div style=\"position:relative; width:100%; height:100%;\">\n    <div style=\"position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);\">\n        {{A}}\n    </div>\n</div>','{\n    \"Q\": \"世界上有几类人？\",\n    \"Options\": [\n        \"A 男人\",\n        \"B 女人\"\n    ],\n    \"A\": \"A, B\"\n}','Q',1,0,'2020-06-16 09:17:39'),
(3,'单词模板','<div style=\"position:relative; width:100%; height:100%;\"> \n <audio id=\"pa\" autoplay style=\"display:none\"> \n  <source :src=\"Detail.wp\" type=\"audio/mpeg\"></source> \n </audio> \n <div style=\"position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);\">\n   <div><h3>{{Word}}</h3></div>\n  <div>\n   [{{Detail.p}}] \n   <button style=\"border:0\" onclick=\"document.getElementById(\'pa\').play()\">⏯</button> \n  </div> \n </div> \n</div>','<div style=\"position:relative; width:100%; height:100%;\"> \n <audio id=\"pb\" autoplay style=\"display:none\"> \n  <source :src=\"Detail.ep\" type=\"audio/mpeg\"></source> \n </audio> \n <div style=\"position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);\"> \n  <div style=\"line-height:20px\">\n    {{Detail.d}} \n  </div> \n  <div style=\"line-height:20px; padding-top:20px\">\n    {{Detail.e}} \n    <button style=\"border:0\" onclick=\"document.getElementById(\'pb\').play()\">⏯</button> \n  </div> \n </div> \n</div>','{\n    \"Word\": \"legacy\",\n    \"Detail\": {\n        \"w\":\"legacy\",\n        \"wp\":\"https://www.collinsdictionary.com/zh/sounds/hwd_sounds/31300.mp3\",\n        \"p\":\"legəsi\",\n        \"d\":\"A legacy is money or property which someone leaves to you when they die.\",\n        \"e\":\"You could make a real difference to someone\'s life by leaving them a generous legacy.\",\n        \"ep\":\"https://www.collinsdictionary.com/zh/sounds/hwd_sounds/en_gb_exa_you_could_make_a_real_difference_to_someone_s_life_by_leaving_them_a_generous_legacy.mp3\"\n    }\n}','Word',1,0,'2020-06-23 15:57:57');