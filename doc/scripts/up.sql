

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



-- 创建 categories 表


    -- id: Mapped[int] = mapped_column(Integer, primary_key=True)           
    -- name: Mapped[str] = mapped_column(String(255), unique=True)                       
    -- rank: Mapped[int] = mapped_column(SmallInteger, server_default='1')  
    
    
    -- def to_dict(self):
    --     return {
    --         '编号': self.id,
    --         '分类名称': self.name,
    --         '排序': self.rank,
    --     }



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



