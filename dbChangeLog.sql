CREATE TABLE `test` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
  `username` varchar(50) CHARACTER SET utf8mb4  NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4  NOT NULL COMMENT '密码',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被删除（0）未删除，（1）已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='测试表';
CREATE TABLE `video_class` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(100) DEFAULT '' COMMENT '分类名称',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '最后修改时间',
  `deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否被删除(0:否 1:是)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='视频分类表';

CREATE TABLE `video_download_url` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '下载地址',
  `hash_url` varchar(255) NOT NULL DEFAULT '' COMMENT '下载地址hash',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '最后修改时间',
  `deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否被删除(0:否 1:是)',
  PRIMARY KEY (`id`),
  KEY `idx_url_del` (`hash_url`,`deleted`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='去除地址记录表';

CREATE TABLE `video_list` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `type_id` int unsigned NOT NULL DEFAULT '0' COMMENT '关联视频分类id',
  `video_name` varchar(200) NOT NULL DEFAULT '' COMMENT '视频名称',
  `image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '视频封面地址',
  `video_url` varchar(255) NOT NULL DEFAULT '' COMMENT '视频播放地址',
  `page_view_num` int unsigned NOT NULL DEFAULT '0' COMMENT '浏览量',
  `md5_file` varchar(200) NOT NULL DEFAULT '' COMMENT '文件指纹',
  `org_download_url` varchar(255) NOT NULL DEFAULT '' COMMENT '原始下载地址',
  `download_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '下载状态：0-未下载，1-下载中，2下载完毕',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '最后修改时间',
  `deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否被删除(0:否 1:是)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频列表';

CREATE TABLE `video_setting` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `key` varchar(255) NOT NULL DEFAULT '' COMMENT '设置项key',
  `value` varchar(255) NOT NULL DEFAULT '' COMMENT '设置项value',
  `notes` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '最后修改时间',
  `deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否被删除(0:否 1:是)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='全局环境变量设置表';