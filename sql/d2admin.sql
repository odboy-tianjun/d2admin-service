/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.235.111
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : 192.168.235.111:3308
 Source Schema         : d2admin

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 04/02/2024 21:15:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for system_api
-- ----------------------------
DROP TABLE IF EXISTS `system_api`;
CREATE TABLE `system_api`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `router_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `router_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `router_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `router_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `router_status` int NOT NULL,
  `api_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `api_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `api_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `api_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `api_status` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `index_npm`(`router_name`, `router_path`, `router_method`) USING BTREE,
  INDEX `idx_system_router_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_system_api_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of system_api
-- ----------------------------
INSERT INTO `system_api` VALUES (5, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'pageUser', '/pageUser', 'POST', '分页查询用户', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (6, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'createUser', '/createUser', 'POST', '新建用户', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (7, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'deleteUser', '/deleteUser', 'POST', '通过id删除用户', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (8, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyUserEmail', '/modifyUserEmail', 'POST', '修改用户邮箱', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (9, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyUserPassword', '/modifyUserPassword', 'POST', '修改用户密码', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (10, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyUserPhone', '/modifyUserAvator', 'POST', '修改用户联系方式', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (11, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'pageRole', '/pageRole', 'POST', '分页查询角色', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (12, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'createRole', '/createRole', 'POST', '新建角色', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (13, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'deleteRole', '/deleteRole', 'POST', '删除角色', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (14, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyRoleName', '/modifyRoleName', 'POST', '修改角色名称', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (15, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'menus', '/menus', 'POST', '获取所有菜单和接口路由', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (16, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'createMenu', '/createMenu', 'POST', '新增菜单', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (17, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'deleteMenu', '/deleteMenu', 'POST', '删除菜单', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (18, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyMenu', '/modifyMenu', 'POST', '修改菜单', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (19, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'pageApi', '/pageApi', 'POST', '分页查询接口路由', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (20, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'registerApi', '/registerApi', 'POST', '注册api', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (21, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'unsubscribeApi', '/unsubscribeApi', 'POST', '注销api', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (22, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'modifyApi', '/modifyApi', 'POST', '修改api', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (23, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'bindUserRole', '/bindUserRole', 'POST', '用户绑定角色', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (24, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'bindRoleMenu', '/bindRoleMenu', 'POST', '角色绑定菜单', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (25, '2024-02-04 17:08:19', '2024-02-04 17:08:19', NULL, 'bindApiMenu', '/bindApiMenu', 'POST', 'Api接口绑定菜单', 1, '', '', '', '', 0);
INSERT INTO `system_api` VALUES (26, '2024-02-04 17:44:51', '2024-02-04 17:44:53', NULL, 'kickOut', '/kickOut', 'POST', '在线踢人', 1, '', '', '', '', 0);

-- ----------------------------
-- Table structure for system_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_menu`;
CREATE TABLE `system_menu`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `menu_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `menu_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `menu_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `router_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `router_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `router_auth` int NOT NULL,
  `router_hidden` int NOT NULL,
  `router_component_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `menu_parent_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_system_menu_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of system_menu
-- ----------------------------
INSERT INTO `system_menu` VALUES (1, '2024-02-04 20:01:59', '2024-02-04 20:02:02', NULL, '顶层页面', 'folder-o', '', '', '', 0, 0, '', NULL);
INSERT INTO `system_menu` VALUES (2, '2024-02-04 20:43:25', '2024-02-04 20:43:29', NULL, '页面1', 'folder-o', '/page1', 'page1', 'page1', 1, 0, 'demo/page1', NULL);
INSERT INTO `system_menu` VALUES (3, '2024-02-04 20:43:25', '2024-02-04 20:43:29', NULL, '页面2', 'folder-o', '/page2', 'page2', 'page2', 1, 0, 'demo/page2', 1);
INSERT INTO `system_menu` VALUES (4, '2024-02-04 21:02:44', '2024-02-04 21:02:47', NULL, '页面3', 'folder-o', '/page3', 'page3', 'page3', 1, 0, 'demo/page3', 1);

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_system_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of system_user
-- ----------------------------
INSERT INTO `system_user` VALUES (1, 'admin', '123456', '2024-02-04 13:17:00', '2024-02-04 13:17:03', NULL, 'admin', 'admin');

SET FOREIGN_KEY_CHECKS = 1;
