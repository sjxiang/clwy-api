

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
    `sex` smallint NOT NULL DEFAULT '9' COMMENT '性别, 默认为9, 0 为男性, 1 为女性, 9 为不选择',
    `company` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '公司.学校名',
    `intro` text COLLATE utf8mb4_unicode_ci COMMENT '简介',
    `role` TINYINT NOT NULL DEFAULT '0' COMMENT '用户组, 默认为0, 0 表示普通用户, 100 表示管理员',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`),
    UNIQUE KEY `idx_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `users`
    (`id`, `username`, `nickname`, `password`, `email`, `avatar`, `sex`, `company`, `intro`, `role`, `created_at`, `updated_at`)
VALUES
    (6, 'admin', '管理员', '$2a$10$KbT1.D6PkJY6A/KgfwuB6ucInSVLsArJyRSGVYOzBo3k0fJ1WaE42', 't.uif@qq.com', 'default.jpeg', 1, '扬州301组合', '啪啪啪小王子', 100, UTC_TIMESTAMP(), NULL),
    (7, 'guest', '游客', '$2a$10$KbT1.D6PkJY6A/KgfwuB6ucInSVLsArJyRSGVYOzBo3k0fJ1WaE42', 'tyr@cisco.cn', 'default.jpeg', 0, '成都九眼桥', 'save the people', 0, UTC_TIMESTAMP(), NULL);



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
    KEY `idx_user_id` (`user_id`),
    UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `courses`
    (`id`, `category_id`, `user_id`, `name`, `image`, `recommended`, `introductory`, `content`, `likes_count`, `chapters_count`, `created_at`, `updated_at`)    
VALUES
    (1, 10, 6, 'Electron+Vue3+AI+云存储--实战跨平台桌面应用', 'default.jpeg', 1, 1, 
    '无论前端、后端，还是应届生，只要具备前端基础，想系统掌握Electron，及整合开发中疑难问题解决方案的，并希望将理论转化实践的，这门课程是你的不二之选。课程融合Electron、Vue3、AI及云存储，以项目为驱动，从零打造跨平台、智能化、高安全性及扩展性的桌面应用。全面应用并高效掌握Electron，及与主流技术整合的疑难问题解决。无论晋升涨薪、技能拓展，还是面试求职，都能让你在激烈竞争中脱颖而出！', 
    9, 60, UTC_TIMESTAMP(), NULL),
    (2, 11, 6, 'AI助手Copilot辅助Go+Flutter打造全栈式在线教育系统', 'default.jpeg', 1, 1, 
    '无论是做后端、还是前端，缺乏大型项目实战经验都是技术进阶和职业晋升一大拦路虎。为了帮助大家突破技术与职业瓶颈，课程采用了高性能热门Go语言、跨平台利器Flutter及强大的PostgreSQL数据库，手把手带你打造一款大型功能全面的全栈在线教育系统（涵盖20+核心功能实现、30+复杂页面设计），同时教你借助AI，10倍+提升开发与学习效能。无论是毕设、求职、晋升、转型还是拓宽技术视野，都能让你受益匪浅。', 
    10, 88, UTC_TIMESTAMP(), NULL);



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

INSERT INTO `chapters`
    (`id`, `course_id`, `title`, `content`, `video`, `rank`)
VALUES
    (1, 1, '第1章 Electron桌面应用实战--课程导学', '1-1 课程导学', 'default.mp4', 1),
    (2, 1, '第2章 开启项目前的准备工作', '2-1 项目需求分析', 'default.mp4', 2),
    (3, 1, '第2章 开启项目前的准备工作', '2-2 桌面端跨平台开发框架介绍', 'default.mp4', 3),
    (4, 1, '第2章 开启项目前的准备工作', '2-3 技术栈选型介绍', 'default.mp4', 4),
    (5, 2, '第1章 在线教育系统--课程导学 ', '1-1 在线教学系统--课程导学', 'default.mp4', 1),
    (6, 2, '第2章 项目介绍与开发环境配置 ', '2-1 在线教育系统项目效果演示', 'default.mp4', 2),
    (7, 2, '第2章 项目介绍与开发环境配置 ', '2-2 在线教育系统技术选型', 'default.mp4', 3),
    (8, 2, '第2章 项目介绍与开发环境配置 ', '2-3 在线教育系统项目代码预览', 'default.mp4', 4),
    (9, 2, '第2章 项目介绍与开发环境配置 ', '2-4 人工智能编程助手--GitHub Copilot 配置', 'default.mp4', 5),
    (10, 2, '第2章 项目介绍与开发环境配置 ', '2-5 Flutter 本地开发环境配置', 'default.mp4', 6),
    (11, 2, '第2章 项目介绍与开发环境配置 ', '2-6 Go 本地开发环境配置', 'default.mp4', 7),
    (12, 2, '第2章 项目介绍与开发环境配置 ', '2-7 项目介绍与开发环境配置总结', 'default.mp4', 8);



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
