-- MySQL dump 10.17  Distrib 10.3.22-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: go_blog
-- ------------------------------------------------------
-- Server version	10.3.22-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父id',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `parent_id` (`parent_id`),
  KEY `sort` (`sort`)
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COMMENT='分类表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'编程语言',0,999,'2018-11-06 02:00:29','2018-11-06 02:00:32',NULL),(2,'GO',1,0,'2018-11-06 02:00:45','2020-02-14 12:41:27',NULL),(3,'gin',2,0,'2018-11-06 02:00:45','2020-02-14 12:43:00',NULL),(4,'iris',2,0,'2018-11-06 02:00:45','2020-02-14 12:43:13',NULL),(5,'echo',2,0,'2018-11-06 02:00:45','2020-02-14 12:43:27',NULL),(6,'PHP',1,0,'2018-11-06 02:00:45','2020-02-14 12:43:39',NULL),(7,'JavaScript',1,0,'2018-11-06 02:00:45','2020-02-14 12:46:00',NULL),(8,'南美洲',1,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:48:18'),(9,'非洲',0,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:47:35'),(10,'大洋洲',1,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:47:59'),(11,'CakePHP',6,0,'2018-11-06 02:00:45','2020-02-14 12:44:30',NULL),(12,'ThinkPHP',6,0,'2018-11-06 02:00:45','2020-02-14 12:44:17',NULL),(13,'Laravel',6,0,'2018-11-06 02:00:45','2020-02-14 12:44:03',NULL),(14,'React',7,0,'2018-11-06 02:00:45','2020-02-14 12:46:14',NULL),(15,'Vue',7,0,'2018-11-06 02:00:45','2020-02-14 12:46:22',NULL),(16,'AngularJS',7,0,'2018-11-06 02:00:45','2020-02-14 12:47:14',NULL),(17,'巴西',8,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:48:11'),(18,'阿根廷',8,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:48:07'),(19,'秘鲁',8,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:48:14'),(20,'埃及',9,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:47:32'),(21,'南非',9,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:47:27'),(22,'肯尼亚',9,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:47:23'),(23,'澳大利亚',10,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:47:47'),(24,'新西兰',10,0,'2018-11-06 02:00:45','2018-11-06 02:00:45','2020-02-14 12:47:40');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '发表用户ID',
  `article_id` bigint(20) NOT NULL COMMENT '评论博文ID',
  `comment_like_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `comment_unlike_count` bigint(20) DEFAULT 0,
  `comment_content` text NOT NULL COMMENT '评论内容',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `article_id` (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (5,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:02:38','2020-02-21 09:02:38',NULL),(6,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:03:07','2020-02-21 09:03:07',NULL),(7,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:03:45','2020-02-21 09:03:45',NULL),(8,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:04:55','2020-02-21 09:04:55',NULL),(9,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:09:32','2020-02-21 09:09:32',NULL),(10,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:10:32','2020-02-21 09:10:32',NULL),(11,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:11:37','2020-02-21 09:11:37',NULL),(12,0,7,0,0,'山东分公司答复大法师打发\n阿萨德法师打发','2020-02-21 09:11:37','2020-02-21 09:11:37',NULL),(13,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:24:14','2020-02-21 09:24:14',NULL),(14,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:28:01','2020-02-21 09:28:01',NULL),(15,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:29:36','2020-02-21 09:29:36',NULL),(16,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:33:10','2020-02-21 09:33:10',NULL),(17,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:36:23','2020-02-21 09:36:23',NULL),(18,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:37:38','2020-02-21 09:37:38',NULL),(19,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:40:15','2020-02-21 09:40:15',NULL),(20,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:43:04','2020-02-21 09:43:04',NULL),(21,0,7,0,0,'山东分公司答复大法师打发\r\n阿萨德法师打发','2020-02-21 09:43:56','2020-02-21 09:43:56',NULL),(22,0,6,0,0,'VUESTC','2020-02-21 09:45:28','2020-02-21 09:45:28',NULL),(23,0,4,0,0,'火腿好吃','2020-02-21 09:48:48','2020-02-21 09:48:48',NULL),(24,0,7,2,2,'xvzxcvz','2020-02-27 12:45:00','2020-02-27 12:45:00',NULL),(25,0,5,0,0,'最佳组件','2020-02-28 03:01:46','2020-02-28 03:01:46',NULL),(26,0,7,23,2,'是打发斯蒂芬阿萨德法师打发','2020-02-28 04:56:45','2020-02-28 04:56:45',NULL);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menus`
--

DROP TABLE IF EXISTS `menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT 'Menus名称',
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父id',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `url` varchar(100) NOT NULL DEFAULT '' COMMENT 'URL',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `icon` varchar(50) DEFAULT '' COMMENT 'icon',
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `parent_id` (`parent_id`),
  KEY `sort` (`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=1005 DEFAULT CHARSET=utf8mb4 COMMENT='Menus表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (1,'系统概况',999,1,'/backend/system/main','2018-11-07 07:36:20','2018-11-07 07:36:20',NULL,'fa fa-dashboard'),(2,'管理员',999,2,'/backend/administrators','2018-11-07 07:36:20','2020-02-17 14:32:06',NULL,'fa fa-users'),(27,'文章',999,3,'#','2020-02-17 10:47:42','2020-02-17 14:32:54',NULL,'fa fa-laptop'),(999,'菜单栏',0,1,'/backend','2018-11-07 07:36:20','2020-02-17 11:26:17',NULL,'fa fa-dashboard'),(1000,'分类管理',27,1,'/backend/categorys','2020-02-17 11:32:35','2020-02-17 11:36:07',NULL,'fa fa-reorder '),(1001,'标签管理',27,2,'/backend/tags','2020-02-17 11:37:52','2020-02-17 11:37:52',NULL,'fa  fa-tags'),(1002,'文章管理',27,3,'/backend/news','2020-02-17 11:40:55','2020-02-17 11:40:55',NULL,'fa fa-circle-o'),(1003,'菜单管理',999,5,'/backend/menus','2020-02-17 11:50:18','2020-02-17 11:50:18',NULL,'fa fa-list'),(1004,'评论',999,7,'/backend/comments','2020-03-04 08:24:49','2020-03-04 08:24:49',NULL,'fa fa-comment');
/*!40000 ALTER TABLE `menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `news`
--

DROP TABLE IF EXISTS `news`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` varchar(100) NOT NULL,
  `title` varchar(250) NOT NULL DEFAULT '' COMMENT '分类名称',
  `descript` varchar(500) NOT NULL DEFAULT '' COMMENT '父id',
  `content` text NOT NULL,
  `tags` varchar(100) DEFAULT NULL,
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `tags_id` varchar(100) NOT NULL DEFAULT '' COMMENT '标签',
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `sort` (`sort`),
  KEY `category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='内容表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `news`
--

LOCK TABLES `news` WRITE;
/*!40000 ALTER TABLE `news` DISABLE KEYS */;
INSERT INTO `news` VALUES (1,'2,3','测试1','测试1','测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1',NULL,1,'2018-11-07 02:05:02','2018-11-07 02:05:04',NULL,'1'),(2,'3','213','wd','<p><strong>dwdwdwd</strong></p>\r\n','',1,'2018-11-07 07:35:33','2018-11-07 07:35:33',NULL,'1'),(3,'3','324234','efef','<p>ef</p>\r\n','',1,'2018-11-07 07:36:20','2018-11-07 07:36:20',NULL,'1'),(4,'2,4,6,20,22','234324','我的','<p>二次沟</p>\r\n\r\n<table border=\"1\" cellpadding=\"1\" cellspacing=\"1\" style=\"width:500px\">\r\n<tbody>\r\n<tr>\r\n<td>\r\n<h2>多吃点</h2>\r\n</td>\r\n<td>&nbsp;</td>\r\n</tr>\r\n<tr>\r\n<td>&nbsp;</td>\r\n<td>&nbsp;</td>\r\n</tr>\r\n<tr>\r\n<td>&nbsp;</td>\r\n<td>&nbsp;</td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n\r\n<p>&nbsp;</p>\r\n','',1,'2018-11-07 07:42:02','2018-11-07 08:36:23',NULL,'1'),(5,'7','React 初学','React JS','<p>组件</p>\r\n',NULL,1,'2020-02-14 12:50:25','2020-02-14 12:50:25',NULL,'1'),(6,'15','Vue','Vue 框架','<p>test</p>\r\n',NULL,1,'2020-02-15 03:08:44','2020-02-16 04:41:32',NULL,'2'),(7,'7','JavaScript Alter','test','<p>&lt;script&gt;alter(&quot;hehe&quot;)&lt;/script&gt;</p>\r\n',NULL,1,'2020-02-15 03:20:02','2020-02-15 03:20:02',NULL,'1');
/*!40000 ALTER TABLE `news` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permission_role`
--

DROP TABLE IF EXISTS `permission_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `permission_role` (
  `permission_id` int(10) unsigned NOT NULL,
  `role_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`permission_id`,`role_id`),
  KEY `permission_role_role_id_foreign` (`role_id`),
  CONSTRAINT `permission_role_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `permission_role_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permission_role`
--

LOCK TABLES `permission_role` WRITE;
/*!40000 ALTER TABLE `permission_role` DISABLE KEYS */;
/*!40000 ALTER TABLE `permission_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `display_name` varchar(255) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `permissions_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='permissions表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_user`
--

DROP TABLE IF EXISTS `role_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_user` (
  `user_id` int(10) unsigned NOT NULL,
  `role_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `role_user_role_id_foreign` (`role_id`),
  CONSTRAINT `role_user_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `role_user_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_user`
--

LOCK TABLES `role_user` WRITE;
/*!40000 ALTER TABLE `role_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `role_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `display_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `roles_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名称',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sort` (`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='标签表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (1,'React',0,'2019-11-06 02:00:29','2019-11-06 02:00:32',NULL),(2,'vue',2,'2020-02-16 04:41:16','2020-02-16 04:41:16',NULL);
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `account` varchar(20) NOT NULL,
  `password` char(32) NOT NULL,
  `descript` varchar(255) DEFAULT '',
  `nickname` char(100) DEFAULT '',
  `email` varchar(100) DEFAULT '',
  `headico` varchar(200) DEFAULT '',
  `online` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `username` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2018-10-22 06:03:48','2020-03-18 04:27:45',NULL,'admin','21232f297a57a5a743894a0e4a801fc3','一个golang iris学习者','灯火阑珊','7146275@qq.com','/uploads/headico/aba25897d629b893a51ff4a7aac073b0_bg2016091802.jpg',1),(2,'2018-11-01 02:53:53','2018-11-01 06:15:55',NULL,'cuijun','3b7fb9742017f12726bcebcd69fb7470','Go Web Iris中文网致力于，在中国国内推广Go语言','众里寻他','10000@qq.cm','/uploads/headico/4862_head_1 (4).png',1),(3,'2018-11-01 06:16:42','2020-03-10 08:32:43',NULL,'test','e10adc3949ba59abbe56e057f20f883e','testtesttesttest','test1231','test03@qq.com','/uploads/headico/8314_head_1 (3).png',1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-18 17:13:01
