

-- 临时设置, 中文
SET character_set_client = utf8mb4;
SET character_set_connection = utf8mb4;
SET character_set_database = utf8mb4;
SET character_set_results = utf8mb4;
SET collation_connection = utf8mb4_unicode_ci;
SET collation_database = utf8mb4_unicode_ci;
SET collation_server = utf8mb4_unicode_ci;


-- 创建数据库
CREATE DATABASE IF NOT EXISTS `clwy_api_development`;

-- 切换数据库
USE `clwy_api_development`;


-- 创建 notices 表

CREATE TABLE `notices` (
    `id` integer unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `content` text COLLATE utf8mb4_unicode_ci,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- 增加
INSERT INTO `notices` (`title`, `content`) VALUES ('文明', '待补充');

-- 批量增加
INSERT INTO `notices` 
    (`title`, `content`) 
VALUES 
    ('折扣牛社区零售', '马昕彤'), 
    ('送外卖选智迈', '无锡乐行'),
    ('大模型', 'deepseek'),
    ('折扣牛批发超市', '中大门');



-- 编辑
UPDATE `notices` 
SET `title` = 'tk教主语录', `content` = '要尽可能生活在离文明中心近的地方，纵然北上广有千般不是，奈何文明边缘地带有亿般不是。' 
WHERE `id` = 1;

-- 删除
DELETE FROM `notices` WHERE `id` = 3;

-- 查询
SELECT * FROM `notices` WHERE `id` = 1;

-- 分页 + 模糊搜索
SELECT `id`, `title`, `content`, `created_at`, `updated_at`
FROM `notices`
WHERE `title` LIKE '%折扣牛%'
ORDER BY updated_at DESC
LIMIT 0, 10;

        
-- 创建 users 表
CREATE TABLE `users` (
    `id` integer unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
    `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
    `nickname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
    `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
    `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮箱',
    `avatar` varchar(1024) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
    `sex` smallint NOT NULL DEFAULT '0' COMMENT '性别',
    `company` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '公司.学校名',
    `intro` text COLLATE utf8mb4_unicode_ci COMMENT '简介',
    `role` TINYINT NOT NULL DEFAULT '0' COMMENT '用户组',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`),
    UNIQUE KEY `idx_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- 创建 categories 表

CREATE TABLE `categories` (
    `id` integer unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
    `rank` integer unsigned NOT NULL DEFAULT '1' COMMENT '排序, 默认 1',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`)    
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `categories` 
    (`id`, `name`, `rank`) 
VALUES 
    (10, '前端开发', 1), 
    (11, '后端开发', 2), 
    (12, '数据库', 3), 
    (13, '服务器运维', 4), 
    (14, '测试', 5);


-- 创建 settings 表
CREATE TABLE `settings` (
    `id` integer unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目名称',
    `icp` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '备案号',
    `copyright` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '版权信息',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

        
INSERT INTO `settings` 
    (`id`, `name`, `icp`, `copyright`) 
VALUES 
    (1, '长乐未央课程网站', '苏ICP备123456789号', '© 2025 长乐未央课程网站 版权所有');



-- 创建 courses 表

CREATE TABLE `courses` (
    `id` integer unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
    `category_id` integer unsigned NOT NULL COMMENT '分类编号',
    `user_id` integer unsigned NOT NULL COMMENT '用户编号',
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '课程名称',
    `image` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '课程图片',
    `recommended` smallint NOT NULL DEFAULT '0' COMMENT '是否推荐, 默认0, 0 表示不推荐, 1 表示推荐',
    `introductory` smallint NOT NULL DEFAULT '0' COMMENT '是否为入门课程, 默认0, 0 表示不是, 1 表示是',
    `content` text COLLATE utf8mb4_unicode_ci COMMENT '课程内容',
    `likes_count` integer unsigned NOT NULL DEFAULT '0' COMMENT '课程的点赞数',
    `chapters_count` integer unsigned NOT NULL DEFAULT '0' COMMENT '课程的章节数量',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_category_id` (`category_id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- 创建 chapters 表
        
CREATE TABLE `chapters` (
    `id` integer unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
    `course_id` integer unsigned NOT NULL COMMENT '课程编号',
    `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '章节标题',
    `content` text COLLATE utf8mb4_unicode_ci COMMENT '章节内容',
    `video` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '章节视频',
    `rank` integer unsigned NOT NULL DEFAULT '1' COMMENT '排序',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_course_id` (`course_id`)
)  ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- 创建 likes 表
        
create table `likes` (
    `id` integer unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
    `course_id` integer unsigned NOT NULL COMMENT '课程编号',
    `user_id` integer unsigned NOT NULL COMMENT '用户编号',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_course_id` (`course_id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
