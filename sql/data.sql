/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : gorm

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2018-11-07 17:06:10
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `admin`
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
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
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `username` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES ('1', '2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'admin', '21232f297a57a5a743894a0e4a801fc3', '一个golang iris学习者', '灯火阑珊', '7146275@qq.com', '/uploads/headico/706_head_1 (5).png');
INSERT INTO `admin` VALUES ('2', '2018-11-01 10:53:53', '2018-11-01 14:15:55', null, 'cuijun', '3b7fb9742017f12726bcebcd69fb7470', 'Go Web Iris中文网致力于，在中国国内推广Go语言', '众里寻他', '10000@qq.cm', '/uploads/headico/4862_head_1 (4).png');
INSERT INTO `admin` VALUES ('3', '2018-11-01 14:16:42', '2018-11-05 11:27:47', null, 'test', '098f6bcd4621d373cade4e832627b4f6', 'testtesttesttest', 'test1231', 'test03@qq.com', '/uploads/headico/8314_head_1 (3).png');

-- ----------------------------
-- Table structure for `category`
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父id',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `parent_id` (`parent_id`),
  KEY `sort` (`sort`)
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COMMENT='分类表';

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES ('1', '国家分类', '0', '999', '2018-11-06 10:00:29', '2018-11-06 10:00:32', null);
INSERT INTO `category` VALUES ('2', '亚洲', '1', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:48', null);
INSERT INTO `category` VALUES ('3', '中国', '2', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('4', '韩国', '2', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('5', '日本', '2', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('6', '北美洲', '1', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('7', '欧洲', '1', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('8', '南美洲', '1', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('9', '非洲', '0', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('10', '大洋洲', '1', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('11', '美国', '6', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('12', '加拿大', '6', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('13', '墨西哥', '6', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('14', '英国', '7', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('15', '法国', '7', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('16', '德国', '7', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('17', '巴西', '8', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('18', '阿根廷', '8', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('19', '秘鲁', '8', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('20', '埃及', '9', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('21', '南非', '9', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('22', '肯尼亚', '9', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('23', '澳大利亚', '10', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);
INSERT INTO `category` VALUES ('24', '新西兰', '10', '0', '2018-11-06 10:00:45', '2018-11-06 10:00:45', null);

-- ----------------------------
-- Table structure for `news`
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` varchar(100) NOT NULL,
  `title` varchar(250) NOT NULL DEFAULT '' COMMENT '分类名称',
  `descript` varchar(500) NOT NULL DEFAULT '' COMMENT '父id',
  `content` text NOT NULL,
  `tags` varchar(100) DEFAULT NULL,
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `sort` (`sort`),
  KEY `category_id` (`category_id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='内容表';

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news` VALUES ('1', '2,3', '测试1', '测试1', '测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1', null, '1', '2018-11-07 10:05:02', '2018-11-07 10:05:04', null);
INSERT INTO `news` VALUES ('2', '3', '213', 'wd', '<p><strong>dwdwdwd</strong></p>\r\n', '', '1', '2018-11-07 15:35:33', '2018-11-07 15:35:33', null);
INSERT INTO `news` VALUES ('3', '3', '324234', 'efef', '<p>ef</p>\r\n', '', '1', '2018-11-07 15:36:20', '2018-11-07 15:36:20', null);
INSERT INTO `news` VALUES ('4', '2,4,6,20,22', '234324', '我的', '<p>二次沟</p>\r\n\r\n<table border=\"1\" cellpadding=\"1\" cellspacing=\"1\" style=\"width:500px\">\r\n	<tbody>\r\n		<tr>\r\n			<td>\r\n			<h2>多吃点</h2>\r\n			</td>\r\n			<td>&nbsp;</td>\r\n		</tr>\r\n		<tr>\r\n			<td>&nbsp;</td>\r\n			<td>&nbsp;</td>\r\n		</tr>\r\n		<tr>\r\n			<td>&nbsp;</td>\r\n			<td>&nbsp;</td>\r\n		</tr>\r\n	</tbody>\r\n</table>\r\n\r\n<p>&nbsp;</p>\r\n', '', '1', '2018-11-07 15:42:02', '2018-11-07 16:36:23', null);