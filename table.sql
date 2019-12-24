CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint(20) unsigned NOT NULL COMMENT '分类id',
  `content` longtext NOT NULL COMMENT '文章内容',
  `title` varchar(1024) NOT NULL COMMENT '文章标题',
  `comment_count` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '评论次数',
  `status` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `summary` varchar(256) NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `view_count` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '阅读数',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_count` (`comment_count`) USING BTREE COMMENT '评论数索引',
  KEY `idx_category_id` (`category_id`) USING BTREE COMMENT '分类id索引'
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL COMMENT '分类名字',
  `category_no` int(10) unsigned NOT NULL COMMENT '分类排序',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


CREATE TABLE `comment` (
  `id` bigint(20) unsigned NOT NULL COMMENT '评论id',
  `content` text NOT NULL COMMENT '评论内容',
  `username` varchar(64) NOT NULL COMMENT '评论作者',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) unsigned NOT NULL COMMENT '评论状态: 0, 删除；1， 正常',
  `article_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


CREATE TABLE `leave` (
  `id` bigint(20) NOT NULL,
  `username` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

