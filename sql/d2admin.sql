/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80025 (8.0.25)
 Source Host           : localhost:3306
 Source Schema         : d2admin

 Target Server Type    : MySQL
 Target Server Version : 80025 (8.0.25)
 File Encoding         : 65001

 Date: 04/02/2024 20:16:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for system_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_menu`;
CREATE TABLE `system_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `menu_title` varchar(255) NOT NULL,
  `menu_icon` varchar(255) NOT NULL DEFAULT '',
  `menu_path` varchar(255) NOT NULL,
  `router_path` varchar(255) NOT NULL,
  `router_name` varchar(255) NOT NULL,
  `router_auth` int NOT NULL,
  `router_hidden` int NOT NULL,
  `router_component_path` varchar(255) NOT NULL,
  `menu_parent_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_system_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of system_menu
-- ----------------------------
BEGIN;
INSERT INTO `system_menu` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_title`, `menu_icon`, `menu_path`, `router_path`, `router_name`, `router_auth`, `router_hidden`, `router_component_path`, `menu_parent_id`) VALUES (1, '2024-02-04 20:01:59', '2024-02-04 20:02:02', NULL, '页面1', '', '/page1', 'page1', 'page1', 1, 0, 'demo/page1', NULL);
COMMIT;

-- ----------------------------
-- Table structure for system_router
-- ----------------------------
DROP TABLE IF EXISTS `system_router`;
CREATE TABLE `system_router` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `router_name` varchar(255) NOT NULL,
  `router_path` varchar(255) NOT NULL,
  `router_method` varchar(255) NOT NULL,
  `router_desc` varchar(255) NOT NULL,
  `router_status` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_npm` (`router_name`,`router_path`,`router_method`),
  KEY `idx_system_router_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of system_router
-- ----------------------------
BEGIN;
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (5, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'pageUser', '/pageUser', 'POST', '分页查询用户', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (6, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'createUser', '/createUser', 'POST', '新建用户', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (7, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'deleteUser', '/deleteUser', 'POST', '通过id删除用户', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (8, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyUserEmail', '/modifyUserEmail', 'POST', '修改用户邮箱', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (9, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyUserPassword', '/modifyUserPassword', 'POST', '修改用户密码', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (10, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyUserPhone', '/modifyUserAvator', 'POST', '修改用户联系方式', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (11, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'pageRole', '/pageRole', 'POST', '分页查询角色', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (12, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'createRole', '/createRole', 'POST', '新建角色', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (13, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'deleteRole', '/deleteRole', 'POST', '删除角色', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (14, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyRoleName', '/modifyRoleName', 'POST', '修改角色名称', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (15, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'menus', '/menus', 'POST', '获取所有菜单和接口路由', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (16, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'createMenu', '/createMenu', 'POST', '新增菜单', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (17, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'deleteMenu', '/deleteMenu', 'POST', '删除菜单', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (18, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyMenu', '/modifyMenu', 'POST', '修改菜单', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (19, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'pageApi', '/pageApi', 'POST', '分页查询接口路由', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (20, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'registerApi', '/registerApi', 'POST', '注册api', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (21, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'unsubscribeApi', '/unsubscribeApi', 'POST', '注销api', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (22, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyApi', '/modifyApi', 'POST', '修改api', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (23, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'bindUserRole', '/bindUserRole', 'POST', '用户绑定角色', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (24, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'bindRoleMenu', '/bindRoleMenu', 'POST', '角色绑定菜单', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (25, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'bindApiMenu', '/bindApiMenu', 'POST', 'Api接口绑定菜单', 1);
INSERT INTO `system_router` (`id`, `created_at`, `updated_at`, `deleted_at`, `router_name`, `router_path`, `router_method`, `router_desc`, `router_status`) VALUES (26, '2024-02-04 17:44:51', '2024-02-04 17:44:53', NULL, 'kickOut', '/kickOut', 'POST', '在线踢人', 1);
COMMIT;

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_system_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of system_user
-- ----------------------------
BEGIN;
INSERT INTO `system_user` (`id`, `username`, `password`, `created_at`, `updated_at`, `deleted_at`, `uuid`, `name`) VALUES (1, 'admin', '123456', '2024-02-04 13:17:00', '2024-02-04 13:17:03', NULL, 'admin', 'admin');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
