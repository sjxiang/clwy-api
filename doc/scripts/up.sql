

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

-- 列表, 分页 + 模糊搜索
SELECT `id`, `title`, `content`, `created_at`, `updated_at`
FROM `notices`
WHERE `title` LIKE '%折扣牛%'
ORDER BY updated_at DESC
LIMIT 0, 10;



-- 创建 users 表
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nickname` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(1024) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sex` smallint NOT NULL DEFAULT '0',
  `company` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `intro` text COLLATE utf8mb4_unicode_ci,
  `role` smallint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO `settings` (`name`, `icp`, `copyright`) VALUES ('bbs论坛', '苏ICP备123456789号', '© 2023 bbs论坛 版权所有');



