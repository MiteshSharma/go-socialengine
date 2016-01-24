CREATE TABLE `user` (
  `user_id` int(20) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`email`),
  KEY `UserIdKey` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `object_detail` (
  `object_id` int(20) NOT NULL,
  `object_type` varchar(128) DEFAULT NULL,
  `like_count` int(10) DEFAULT 0,
  `comment_count` int(10) DEFAULT 0,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`object_id`,`object_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `object_like` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `object_id` int(20) NOT NULL,
  `object_type` varchar(128) DEFAULT NULL,
  `user_id` int(20) NOT NULL,
  `is_deleted` tinyint(1) DEFAULT 0,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `UserIdIndex` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `object_comment` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `object_id` int(20) NOT NULL,
  `object_type` varchar(128) DEFAULT NULL,
  `user_id` int(20) NOT NULL,
  `parent_id` int(20) DEFAULT NULL,
  `comment` varchar(255) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT 0,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `UserIdIndex` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;