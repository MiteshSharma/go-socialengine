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

CREATE TABLE `channel` (
  `id` bigint(12) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `user_count` int(10) DEFAULT '1',
  `owner_id` bigint(12) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `NameIndex` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE `channel_detail` (
  `user_id` bigint(12) NOT NULL,
  `channel_id` bigint(12) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `ChannelIndex` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `channel_property` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `channel_id` bigint(12) NOT NULL,
  `user_id` bigint(12) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ChannelIndex` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `channel_share_object` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(12) DEFAULT NULL,
  `channel_id` bigint(12) DEFAULT NULL,
  `object_id` int(20) NOT NULL,
  `object_type` varchar(128) DEFAULT NULL,
  `message` text DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ChannelIndex` (`channel_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;