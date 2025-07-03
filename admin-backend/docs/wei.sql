/*
 Navicat Premium CustomClaims Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80400
 Source Host           : 127.0.0.1:3306
 Source Schema         : wei

 Target Server Type    : MySQL
 Target Server Version : 80400
 File Encoding         : 65001

 Date: 23/06/2024 22:43:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_api
-- ----------------------------
DROP TABLE IF EXISTS `admin_api`;
CREATE TABLE `admin_api`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '接口ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '接口路由',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '接口唯一名称',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '接口名称',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '接口描述',
  `is_enabled` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '接口状态：1：正常，0：禁用',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `uk_key`(`key` ASC) USING BTREE,
  UNIQUE INDEX `uk_path`(`path` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '接口权限关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_api
-- ----------------------------
INSERT INTO `admin_api` VALUES (1, '/admin/user/list', 'adminUser::list', '账号列表', '账号列表', 0, '2024-06-23 13:53:39', '2024-06-23 13:57:51');
INSERT INTO `admin_api` VALUES (2, '/admin/user/add', 'adminUser::add', '账号创建', '账号创建', 1, '2022-08-07 11:01:21', '2022-08-07 11:01:21');
INSERT INTO `admin_api` VALUES (3, '/admin/user/detail', 'adminUser::detail', '账号详情', '账号详情', 1, '2022-08-07 10:59:05', '2022-08-07 10:59:05');
INSERT INTO `admin_api` VALUES (4, '/admin/user/edit', 'adminUser::edit', '账号编辑', '账号编辑', 1, '2022-08-07 11:24:04', '2022-08-07 11:24:04');
INSERT INTO `admin_api` VALUES (5, '/admin/user/enable', 'adminUser::enable', '账号启用禁用', '账号启用禁用', 1, '2022-08-07 11:31:53', '2022-08-07 11:31:53');
INSERT INTO `admin_api` VALUES (6, '/admin/user/delete', 'adminUser::delete', '账号删除', '账号删除', 1, '2022-08-07 11:31:58', '2022-08-07 11:31:58');
INSERT INTO `admin_api` VALUES (7, '/admin/user/bindRoles', 'adminUser::bindRoles', '账号绑定角色', '账号绑定角色', 1, '2022-08-07 11:37:20', '2022-08-07 11:37:20');
INSERT INTO `admin_api` VALUES (8, '/admin/role/list', 'adminRole::list', '角色列表', '角色列表', 1, '2022-08-07 11:48:30', '2022-08-07 11:48:30');
INSERT INTO `admin_api` VALUES (9, '/admin/role/add', 'adminRole::add', '角色创建', '', 1, '2022-08-07 11:48:55', '2022-08-07 11:48:56');
INSERT INTO `admin_api` VALUES (10, '/admin/role/detail', 'adminRole::detail', '角色详情', '', 1, '2022-08-07 11:50:29', '2022-08-07 11:50:29');
INSERT INTO `admin_api` VALUES (11, '/admin/role/edit', 'adminRole::edit', '角色编辑', '', 1, '2022-08-07 12:14:07', '2022-08-07 12:14:07');
INSERT INTO `admin_api` VALUES (12, '/admin/role/enable', 'adminRole::enable', '角色禁用启用', '', 1, '2022-08-07 12:14:38', '2022-08-07 12:14:38');
INSERT INTO `admin_api` VALUES (13, '/admin/role/delete', 'adminRole::delete', '角色删除', '', 1, '2022-08-07 12:15:21', '2022-08-07 12:15:21');
INSERT INTO `admin_api` VALUES (14, '/admin/role/bindPermissions', 'adminRole::bindPermissions', '角色绑定权限', '', 1, '2022-08-07 12:16:03', '2022-08-07 12:16:03');
INSERT INTO `admin_api` VALUES (15, '/admin/role/permissions', 'adminRole::permissions', '角色权限列表', '', 1, '2022-08-07 12:16:56', '2022-08-07 12:16:56');
INSERT INTO `admin_api` VALUES (16, '/admin/menu/tree', 'adminMenu::tree', '菜单树', '', 1, '2022-08-07 12:27:17', '2022-08-07 12:27:17');
INSERT INTO `admin_api` VALUES (17, '/admin/menu/list', 'adminMenu::list', '菜单列表', '', 1, '2022-08-07 12:27:40', '2022-08-07 12:27:40');
INSERT INTO `admin_api` VALUES (18, '/admin/menu/add', 'adminMenu::add', '菜单创建', '', 1, '2022-08-07 12:28:02', '2022-08-07 12:28:02');
INSERT INTO `admin_api` VALUES (19, '/admin/menu/detail', 'adminMenu::detail', '菜单详情', '', 1, '2022-08-07 12:28:17', '2022-08-07 12:28:17');
INSERT INTO `admin_api` VALUES (20, '/admin/menu/edit', 'adminMenu::edit', '菜单编辑', '', 1, '2022-08-07 12:28:38', '2022-08-07 12:28:38');
INSERT INTO `admin_api` VALUES (21, '/admin/menu/enable', 'adminMenu::enable', '菜单禁用启用', '', 1, '2022-08-07 12:28:52', '2022-08-07 12:28:52');
INSERT INTO `admin_api` VALUES (22, '/admin/menu/delete', 'adminMenu::delete', '菜单删除', '', 1, '2022-08-07 12:29:06', '2022-08-07 12:29:06');
INSERT INTO `admin_api` VALUES (23, '/admin/menu/permissions', 'adminMenu::permissions', '菜单权限', '', 1, '2022-08-07 14:18:02', '2022-08-07 14:18:03');
INSERT INTO `admin_api` VALUES (24, '/admin/menu/pages', 'adminMenu::pages', '菜单页面列表', '', 1, '2022-08-07 12:30:16', '2022-08-07 12:30:16');
INSERT INTO `admin_api` VALUES (25, '/admin/permission/list', 'adminPermission::list', '权限列表', '', 1, '2022-08-07 12:31:00', '2022-08-07 12:31:00');
INSERT INTO `admin_api` VALUES (26, '/admin/permission/add', 'adminPermission::add', '权限创建', '', 1, '2022-08-07 12:31:15', '2022-08-07 12:31:15');
INSERT INTO `admin_api` VALUES (27, '/admin/permission/detail', 'adminPermission::detail', '权限详情', '', 1, '2022-08-07 12:31:29', '2022-08-07 12:31:29');
INSERT INTO `admin_api` VALUES (28, '/admin/permission/edit', 'adminPermission::edit', '权限编辑', '', 1, '2022-08-07 12:36:02', '2022-08-07 12:36:02');
INSERT INTO `admin_api` VALUES (29, '/admin/permission/enable', 'adminPermission::enable', '权限禁用启用', '', 1, '2022-08-07 12:36:18', '2022-08-07 12:36:18');
INSERT INTO `admin_api` VALUES (30, '/admin/permission/delete', 'adminPermission::delete', '权限删除', '', 1, '2022-08-07 12:36:28', '2022-08-07 12:36:28');
INSERT INTO `admin_api` VALUES (31, '/admin/permission/addMenuPermissions', 'adminPermission::addMenuPermissions', '权限指定菜单批量创建权限', '给指定的菜单创建查看，编辑，删除权限', 1, '2022-08-07 12:37:22', '2022-08-07 12:37:22');
INSERT INTO `admin_api` VALUES (32, '/admin/permission/bindApis', 'adminPermission::bindApis', '权限绑定接口', '', 1, '2022-08-07 12:37:40', '2022-08-07 12:37:40');
INSERT INTO `admin_api` VALUES (33, '/admin/api/list', 'adminApi::list', '接口列表', '', 1, '2022-08-07 12:38:24', '2022-08-07 12:38:24');
INSERT INTO `admin_api` VALUES (34, '/admin/api/add', 'adminApi::add', '接口创建', '', 1, '2022-08-07 12:38:44', '2022-08-07 12:38:44');
INSERT INTO `admin_api` VALUES (35, '/admin/api/detail', 'adminApi::detail', '接口详情', '', 1, '2022-08-07 12:38:59', '2022-08-07 12:38:59');
INSERT INTO `admin_api` VALUES (36, '/admin/api/edit', 'adminApi::edit', '接口编辑', '', 1, '2022-08-07 12:39:10', '2022-08-07 12:39:10');
INSERT INTO `admin_api` VALUES (37, '/admin/api/enable', 'adminApi::enable', '接口禁用启用', '', 1, '2022-08-07 12:39:23', '2022-08-07 12:39:23');
INSERT INTO `admin_api` VALUES (38, '/admin/api/delete', 'adminApi::delete', '接口删除', '', 1, '2022-08-07 12:39:34', '2022-08-07 12:39:34');
INSERT INTO `admin_api` VALUES (39, '/admin/api/all', 'adminApi::all', '接口全部', '全部有效接口列表', 1, '2022-08-07 12:40:07', '2022-08-07 12:40:07');
INSERT INTO `admin_api` VALUES (40, '/admin/role/all', 'adminRole::all', '角色全部', '全部有效的角色列表', 1, '2022-08-07 16:26:45', '2022-08-07 16:26:45');
INSERT INTO `admin_api` VALUES (41, '/admin/menu/mode', 'adminMenu::mode', '菜单页面权限列表', '', 1, '2022-08-08 03:18:10', '2022-08-08 03:18:10');

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '路径',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限中文名称',
  `key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '菜单的唯一键名',
  `describe` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '路径图标',
  `sort` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序值',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '重定向路径',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '组件名称',
  `is_hide_in_menu` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否隐藏：0显示，1隐藏',
  `is_hide_children_in_menu` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否在children中隐藏：1隐藏，0显示',
  `is_enabled` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '1：启用，0禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `uk_key`(`key` ASC) USING BTREE,
  UNIQUE INDEX `uk_path`(`path` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
INSERT INTO `admin_menu` VALUES (1, 0, '/admin', '系统设置', 'Admin', '系统设置', '', 0, '/', '', 0, 0, 1, '2022-06-26 22:28:24', '2022-06-26 22:28:24');
INSERT INTO `admin_menu` VALUES (2, 1, '/admin/user', '账号管理', 'AdminUser', '账号列表', '', 0, '/', '', 0, 1, 1, '2022-08-06 15:09:09', '2022-08-11 00:57:19');
INSERT INTO `admin_menu` VALUES (3, 1, '/admin/role', '角色管理', 'AdminRole', '角色列表', '', 0, '/', '', 0, 1, 1, '2022-08-06 15:08:22', '2022-08-10 22:45:34');
INSERT INTO `admin_menu` VALUES (4, 1, '/admin/menu', '菜单管理', 'AdminMenu', '菜单列表', '', 0, '/', '', 0, 1, 1, '2022-06-26 22:29:08', '2022-08-08 02:41:15');
INSERT INTO `admin_menu` VALUES (5, 1, '/admin/permission', '权限管理', 'AdminPermission', '权限管理', '', 0, '/', '', 0, 1, 1, '2022-08-06 15:10:20', '2022-08-11 00:57:49');
INSERT INTO `admin_menu` VALUES (6, 1, '/admin/api', '接口管理', 'AdminApi', '接口管理', '', 0, '/', '', 0, 0, 1, '2022-08-06 15:11:37', '2022-08-06 15:11:40');
INSERT INTO `admin_menu` VALUES (7, 4, '/admin/menu/add', '添加菜单', 'AdminMenuAdd', '添加菜单', '', 0, '/', '', 1, 1, 1, '2022-08-06 15:12:39', '2022-08-11 00:57:29');
INSERT INTO `admin_menu` VALUES (8, 4, '/admin/menu/edit', '编辑菜单', 'AdminMenuEdit', '编辑菜单', '', 0, '/', '', 1, 1, 1, '2022-08-06 15:13:40', '2022-08-11 00:57:33');

-- ----------------------------
-- Table structure for admin_permission
-- ----------------------------
DROP TABLE IF EXISTS `admin_permission`;
CREATE TABLE `admin_permission`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `menu_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属菜单ID',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限唯一标识名称',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限显示名称',
  `type` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL DEFAULT 'view' COMMENT '权限的操作类型\r\nview：查看（只读）\r\nedit：编辑（读写）\r\ndelete：删除（彻底删除）',
  `describe` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限描述',
  `is_enabled` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否启用：1启用，0禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_key`(`key` ASC) USING BTREE,
  UNIQUE INDEX `uk_permission`(`menu_id` ASC, `type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_permission
-- ----------------------------
INSERT INTO `admin_permission` VALUES (1, 2, 'AdminUserView', '账号查看', 'view', '账号管理查看', 1, '2022-08-07 10:42:31', '2022-08-07 10:42:31');
INSERT INTO `admin_permission` VALUES (2, 2, 'AdminUserEdit', '账号编辑', 'edit', '账号管理编辑', 1, '2022-08-06 20:20:58', '2022-08-06 20:20:58');
INSERT INTO `admin_permission` VALUES (3, 2, 'AdminUserDelete', '账号删除', 'delete', '账号管理删除', 1, '2022-08-06 20:20:58', '2022-08-06 20:20:58');
INSERT INTO `admin_permission` VALUES (4, 3, 'AdminRoleView', '角色查看', 'view', '角色管理查看', 1, '2022-08-06 20:21:52', '2022-08-06 20:21:52');
INSERT INTO `admin_permission` VALUES (5, 3, 'AdminRoleEdit', '角色编辑', 'edit', '角色管理编辑', 1, '2022-08-06 20:21:52', '2022-08-06 20:21:52');
INSERT INTO `admin_permission` VALUES (6, 3, 'AdminRoleDelete', '角色删除', 'delete', '角色管理删除', 1, '2022-08-06 20:21:52', '2022-08-06 20:21:52');
INSERT INTO `admin_permission` VALUES (7, 4, 'AdminMenuView', '菜单查看', 'view', '菜单管理查看', 1, '2022-08-06 20:22:07', '2022-08-06 20:22:07');
INSERT INTO `admin_permission` VALUES (8, 4, 'AdminMenuEdit', '菜单编辑', 'edit', '菜单管理编辑', 1, '2022-08-06 20:22:07', '2022-08-06 20:22:07');
INSERT INTO `admin_permission` VALUES (9, 4, 'AdminMenuDelete', '菜单删除', 'delete', '菜单管理删除', 1, '2022-08-06 20:22:07', '2022-08-06 20:22:07');
INSERT INTO `admin_permission` VALUES (10, 5, 'AdminPermissionView', '权限查看', 'view', '权限管理查看', 1, '2022-08-06 20:22:17', '2022-08-06 20:22:17');
INSERT INTO `admin_permission` VALUES (11, 5, 'AdminPermissionEdit', '权限编辑', 'edit', '权限管理编辑', 1, '2022-08-06 20:22:17', '2022-08-06 20:22:17');
INSERT INTO `admin_permission` VALUES (12, 5, 'AdminPermissionDelete', '权限删除', 'delete', '权限管理删除', 1, '2022-08-06 20:22:17', '2022-08-06 20:22:17');
INSERT INTO `admin_permission` VALUES (13, 6, 'AdminApiView', '接口查看', 'view', '接口管理查看', 1, '2022-08-06 20:22:27', '2022-08-06 20:22:27');
INSERT INTO `admin_permission` VALUES (14, 6, 'AdminApiEdit', '接口编辑', 'edit', '接口管理编辑', 1, '2022-08-06 20:22:27', '2022-08-06 20:22:27');
INSERT INTO `admin_permission` VALUES (15, 6, 'AdminApiDelete', '接口删除', 'delete', '接口管理删除', 1, '2022-08-06 20:22:27', '2022-08-06 20:22:27');

-- ----------------------------
-- Table structure for admin_permission_api
-- ----------------------------
DROP TABLE IF EXISTS `admin_permission_api`;
CREATE TABLE `admin_permission_api`  (
  `permission_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限ID',
  `api_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '接口ID',
  UNIQUE INDEX `uk_permission_api`(`permission_id` ASC, `api_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '接口权限关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_permission_api
-- ----------------------------
INSERT INTO `admin_permission_api` VALUES (1, 1);
INSERT INTO `admin_permission_api` VALUES (1, 3);
INSERT INTO `admin_permission_api` VALUES (1, 40);
INSERT INTO `admin_permission_api` VALUES (2, 2);
INSERT INTO `admin_permission_api` VALUES (2, 4);
INSERT INTO `admin_permission_api` VALUES (2, 5);
INSERT INTO `admin_permission_api` VALUES (2, 7);
INSERT INTO `admin_permission_api` VALUES (3, 6);
INSERT INTO `admin_permission_api` VALUES (4, 8);
INSERT INTO `admin_permission_api` VALUES (4, 10);
INSERT INTO `admin_permission_api` VALUES (4, 15);
INSERT INTO `admin_permission_api` VALUES (4, 40);
INSERT INTO `admin_permission_api` VALUES (4, 41);
INSERT INTO `admin_permission_api` VALUES (5, 9);
INSERT INTO `admin_permission_api` VALUES (5, 11);
INSERT INTO `admin_permission_api` VALUES (5, 12);
INSERT INTO `admin_permission_api` VALUES (5, 14);
INSERT INTO `admin_permission_api` VALUES (6, 13);
INSERT INTO `admin_permission_api` VALUES (7, 16);
INSERT INTO `admin_permission_api` VALUES (7, 17);
INSERT INTO `admin_permission_api` VALUES (7, 19);
INSERT INTO `admin_permission_api` VALUES (7, 23);
INSERT INTO `admin_permission_api` VALUES (7, 24);
INSERT INTO `admin_permission_api` VALUES (7, 41);
INSERT INTO `admin_permission_api` VALUES (8, 18);
INSERT INTO `admin_permission_api` VALUES (8, 20);
INSERT INTO `admin_permission_api` VALUES (8, 21);
INSERT INTO `admin_permission_api` VALUES (8, 31);
INSERT INTO `admin_permission_api` VALUES (9, 22);
INSERT INTO `admin_permission_api` VALUES (10, 25);
INSERT INTO `admin_permission_api` VALUES (10, 27);
INSERT INTO `admin_permission_api` VALUES (11, 26);
INSERT INTO `admin_permission_api` VALUES (11, 28);
INSERT INTO `admin_permission_api` VALUES (11, 29);
INSERT INTO `admin_permission_api` VALUES (11, 32);
INSERT INTO `admin_permission_api` VALUES (12, 30);
INSERT INTO `admin_permission_api` VALUES (13, 33);
INSERT INTO `admin_permission_api` VALUES (13, 35);
INSERT INTO `admin_permission_api` VALUES (13, 39);
INSERT INTO `admin_permission_api` VALUES (14, 34);
INSERT INTO `admin_permission_api` VALUES (14, 36);
INSERT INTO `admin_permission_api` VALUES (14, 37);
INSERT INTO `admin_permission_api` VALUES (15, 38);

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '角色名称',
  `describe` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL DEFAULT '' COMMENT '角色描述',
  `modify_admin_id` int UNSIGNED NOT NULL COMMENT '修改人',
  `create_admin_id` int UNSIGNED NOT NULL COMMENT '创建人',
  `is_enabled` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '1：启用，0：禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '管理员角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, '管理员', '拥有全部权限', 9, 1, 1, '2022-08-11 02:09:58', '2022-08-11 02:09:59');

-- ----------------------------
-- Table structure for admin_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_permission`;
CREATE TABLE `admin_role_permission`  (
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `permission_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限ID',
  UNIQUE INDEX `uk_role_permission`(`role_id` ASC, `permission_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '角色权限关系表（包含菜单和权限）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role_permission
-- ----------------------------
INSERT INTO `admin_role_permission` VALUES (1, 1);
INSERT INTO `admin_role_permission` VALUES (1, 2);
INSERT INTO `admin_role_permission` VALUES (1, 3);
INSERT INTO `admin_role_permission` VALUES (1, 4);
INSERT INTO `admin_role_permission` VALUES (1, 5);
INSERT INTO `admin_role_permission` VALUES (1, 6);
INSERT INTO `admin_role_permission` VALUES (1, 7);
INSERT INTO `admin_role_permission` VALUES (1, 8);

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '管理员账号',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '管理昵称姓名',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '登录密码',
  `avatar` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户头像',
  `login_total` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录次数',
  `last_login_ip` json NOT NULL COMMENT '上次登录IP',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '上次登录时间',
  `is_enabled` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '账户状态：1正常，0：禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', '骑着八戒游天河', 'ddd@q1.com', '$2a$10$8DN3n4k4C7H3vrYoQ5AUA.5KrOJQYcUCaq5X7J94JmHRH.XeiEm.m', '', 39, '[\"127.0.0.1\", \"127.0.0.1\"]', '2024-06-23 11:31:49', 1, '2024-06-23 03:31:49', '2024-06-23 11:31:49');
INSERT INTO `admin_user` VALUES (5, 'test00001', '测试00001', '', '$2a$10$xEnugTiRvgBY1n21Mg7g7uCnzBP7aA9G0vzUv.jAnTgF2tM3JBSsC', '', 0, '[]', NULL, 1, '2023-12-04 19:31:25', '2023-12-04 19:31:25');
INSERT INTO `admin_user` VALUES (7, 'test00002', '测试00002', '', '$2a$10$MHuAUxjZAG.8bITT12hEZu6qFb9C8izBM7NSe/FPB4Q3Jth29NnBW', '', 0, '[]', NULL, 1, '2023-12-04 19:31:26', '2023-12-04 19:31:26');
INSERT INTO `admin_user` VALUES (8, 'test00003', '测试00003', '', '$2a$10$l2YYmOBMX0WX3a27NGFYyeGwHAbi6Jozu0k.P/YDjnkYKs94Bi75u', '', 0, '[]', NULL, 1, '2023-12-04 19:31:27', '2023-12-04 19:31:27');
INSERT INTO `admin_user` VALUES (9, 'test00004', '测试00004', 'test0004@qq.com', '$2a$10$Z/WRvsc9apFLcD5DwWjvleZAaaPja99IeOuTXFFzmDZ0tCn8FVLjm', '', 19, '[]', '2022-08-11 02:17:29', 1, '2023-12-04 19:31:30', '2023-12-04 19:31:30');

-- ----------------------------
-- Table structure for admin_user_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_role`;
CREATE TABLE `admin_user_role`  (
  `admin_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  UNIQUE INDEX `uk_admin_id_role_id`(`admin_id` ASC, `role_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '管理员-游戏-角色关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user_role
-- ----------------------------
INSERT INTO `admin_user_role` VALUES (9, 1);

SET FOREIGN_KEY_CHECKS = 1;
