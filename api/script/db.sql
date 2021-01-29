SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
Use video_server;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`login_name` varchar(64) DEFAULT NULL,
	`pwd`		text NOT NULL,
  	PRIMARY KEY (`id`),
	UNIQUE KEY (`login_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for video_info
-- ----------------------------
DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info` (
	`id` varchar(64) NOT NULL,
	`author_id` int(10) DEFAULT NULL,
	`name`	text DEFAULT NULL,
	`display_ctime` text DEFAULT NULL,
	`create_time` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for comments;
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
	`id` varchar(64) NOT NULL,
	`video_id` varchar(64),
	`author_id` int(10) DEFAULT NULL,
	`content` text,
	`time` timestamp DEFAULT CURRENT_TIMESTAMP,
  	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sessions;
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
	`session_id` varchar(200) NOT NULL,
	`TTL` tinytext NOT NULL,
	`login_name` text,
  	PRIMARY KEY (`session_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for video_del_rec;
-- ----------------------------
DROP TABLE IF EXISTS `video_del_rec`;
CREATE TABLE `video_del_rec` (
	`video_id` varchar(64) NOT NULL,
  	PRIMARY KEY (`video_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;














