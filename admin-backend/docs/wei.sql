/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : wei

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 05/09/2026 18:31:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_api
-- ----------------------------
DROP TABLE IF EXISTS `admin_api`;
CREATE TABLE `admin_api`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '接口ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '接口路由',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '接口唯一名称',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '接口名称',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '接口描述',
  `is_enabled` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '接口状态：1：正常，0：禁用',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `uk_key`(`key` ASC) USING BTREE,
  UNIQUE INDEX `uk_path`(`path` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 52 CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '接口权限关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_api
-- ----------------------------
INSERT INTO `admin_api` VALUES (1, '/admin/user/list', 'adminUser::list', '账号列表', '', 1, '2024-06-23 13:53:39', '2024-12-02 20:20:50');
INSERT INTO `admin_api` VALUES (2, '/admin/user/add', 'adminUser::add', '账号创建', '账号创建', 1, '2022-08-07 11:01:21', '2022-08-07 11:01:21');
INSERT INTO `admin_api` VALUES (3, '/admin/user/info', 'adminUser::detail', '账号详情', '账号详情', 1, '2022-08-07 10:59:05', '2026-09-04 10:55:33');
INSERT INTO `admin_api` VALUES (4, '/admin/user/edit', 'adminUser::edit', '账号编辑', '账号编辑', 1, '2022-08-07 11:24:04', '2022-08-07 11:24:04');
INSERT INTO `admin_api` VALUES (5, '/admin/user/enable', 'adminUser::enable', '账号启用禁用', '账号启用禁用', 1, '2022-08-07 11:31:53', '2022-08-07 11:31:53');
INSERT INTO `admin_api` VALUES (6, '/admin/user/delete', 'adminUser::delete', '账号删除', '账号删除', 1, '2022-08-07 11:31:58', '2022-08-07 11:31:58');
INSERT INTO `admin_api` VALUES (7, '/admin/user/bind_roles', 'adminUser::bindRoles', '账号绑定角色', '账号绑定角色', 1, '2022-08-07 11:37:20', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (8, '/admin/role/list', 'adminRole::list', '角色列表', '角色列表', 1, '2022-08-07 11:48:30', '2022-08-07 11:48:30');
INSERT INTO `admin_api` VALUES (9, '/admin/role/add', 'adminRole::add', '角色创建', '', 1, '2022-08-07 11:48:55', '2022-08-07 11:48:56');
INSERT INTO `admin_api` VALUES (10, '/admin/role/info', 'adminRole::detail', '角色详情', '', 1, '2022-08-07 11:50:29', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (11, '/admin/role/edit', 'adminRole::edit', '角色编辑', '', 1, '2022-08-07 12:14:07', '2022-08-07 12:14:07');
INSERT INTO `admin_api` VALUES (12, '/admin/role/enable', 'adminRole::enable', '角色禁用启用', '', 1, '2022-08-07 12:14:38', '2022-08-07 12:14:38');
INSERT INTO `admin_api` VALUES (13, '/admin/role/delete', 'adminRole::delete', '角色删除', '', 1, '2022-08-07 12:15:21', '2022-08-07 12:15:21');
INSERT INTO `admin_api` VALUES (14, '/admin/role/bind_permissions', 'adminRole::bindPermissions', '角色绑定权限', '', 1, '2022-08-07 12:16:03', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (15, '/admin/role/permissions', 'adminRole::permissions', '角色权限列表', '', 1, '2022-08-07 12:16:56', '2022-08-07 12:16:56');
INSERT INTO `admin_api` VALUES (16, '/admin/menu/tree', 'adminMenu::tree', '菜单树', '', 1, '2022-08-07 12:27:17', '2022-08-07 12:27:17');
INSERT INTO `admin_api` VALUES (17, '/admin/menu/list', 'adminMenu::list', '菜单列表', '', 1, '2022-08-07 12:27:40', '2022-08-07 12:27:40');
INSERT INTO `admin_api` VALUES (18, '/admin/menu/add', 'adminMenu::add', '菜单创建', '', 1, '2022-08-07 12:28:02', '2022-08-07 12:28:02');
INSERT INTO `admin_api` VALUES (19, '/admin/menu/info', 'adminMenu::detail', '菜单详情', '', 1, '2022-08-07 12:28:17', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (20, '/admin/menu/edit', 'adminMenu::edit', '菜单编辑', '', 1, '2022-08-07 12:28:38', '2022-08-07 12:28:38');
INSERT INTO `admin_api` VALUES (21, '/admin/menu/enable', 'adminMenu::enable', '菜单禁用启用', '', 1, '2022-08-07 12:28:52', '2022-08-07 12:28:52');
INSERT INTO `admin_api` VALUES (22, '/admin/menu/delete', 'adminMenu::delete', '菜单删除', '', 1, '2022-08-07 12:29:06', '2022-08-07 12:29:06');
INSERT INTO `admin_api` VALUES (23, '/admin/menu/permissions', 'adminMenu::permissions', '菜单权限', '', 1, '2022-08-07 14:18:02', '2022-08-07 14:18:03');
INSERT INTO `admin_api` VALUES (24, '/admin/menu/pages', 'adminMenu::pages', '菜单页面列表', '', 1, '2022-08-07 12:30:16', '2022-08-07 12:30:16');
INSERT INTO `admin_api` VALUES (25, '/admin/permission/list', 'adminPermission::list', '权限列表', '', 1, '2022-08-07 12:31:00', '2022-08-07 12:31:00');
INSERT INTO `admin_api` VALUES (26, '/admin/permission/add', 'adminPermission::add', '权限创建', '', 1, '2022-08-07 12:31:15', '2022-08-07 12:31:15');
INSERT INTO `admin_api` VALUES (27, '/admin/permission/info', 'adminPermission::detail', '权限详情', '', 1, '2022-08-07 12:31:29', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (28, '/admin/permission/edit', 'adminPermission::edit', '权限编辑', '', 1, '2022-08-07 12:36:02', '2022-08-07 12:36:02');
INSERT INTO `admin_api` VALUES (29, '/admin/permission/enable', 'adminPermission::enable', '权限禁用启用', '', 1, '2022-08-07 12:36:18', '2022-08-07 12:36:18');
INSERT INTO `admin_api` VALUES (30, '/admin/permission/delete', 'adminPermission::delete', '权限删除', '', 1, '2022-08-07 12:36:28', '2022-08-07 12:36:28');
INSERT INTO `admin_api` VALUES (31, '/admin/permission/add_menu_permissions', 'adminPermission::addMenuPermissions', '权限指定菜单批量创建权限', '给指定的菜单创建查看，编辑，删除权限', 1, '2022-08-07 12:37:22', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (32, '/admin/permission/bind_apis', 'adminPermission::bindApis', '权限绑定接口', '', 1, '2022-08-07 12:37:40', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (33, '/admin/api/list', 'adminApi::list', '接口列表', '', 1, '2022-08-07 12:38:24', '2022-08-07 12:38:24');
INSERT INTO `admin_api` VALUES (34, '/admin/api/add', 'adminApi::add', '接口创建', '', 1, '2022-08-07 12:38:44', '2022-08-07 12:38:44');
INSERT INTO `admin_api` VALUES (35, '/admin/api/info', 'adminApi::detail', '接口详情', '', 1, '2022-08-07 12:38:59', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (36, '/admin/api/edit', 'adminApi::edit', '接口编辑', '', 1, '2022-08-07 12:39:10', '2022-08-07 12:39:10');
INSERT INTO `admin_api` VALUES (37, '/admin/api/enable', 'adminApi::enable', '接口禁用启用', '', 1, '2022-08-07 12:39:23', '2022-08-07 12:39:23');
INSERT INTO `admin_api` VALUES (38, '/admin/api/delete', 'adminApi::delete', '接口删除', '', 1, '2022-08-07 12:39:34', '2022-08-07 12:39:34');
INSERT INTO `admin_api` VALUES (39, '/admin/api/all', 'adminApi::all', '接口全部', '全部有效接口列表', 1, '2022-08-07 12:40:07', '2022-08-07 12:40:07');
INSERT INTO `admin_api` VALUES (40, '/admin/role/all', 'adminRole::all', '角色全部', '全部有效的角色列表', 1, '2022-08-07 16:26:45', '2022-08-07 16:26:45');
INSERT INTO `admin_api` VALUES (41, '/admin/menu/modes', 'adminMenu::mode', '菜单页面权限列表', '', 1, '2022-08-08 03:18:10', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (43, '/admin/user/edit_pwd', 'adminUser::editPwd', '账号重置密码', '账号重置密码', 1, '2026-09-04 10:55:34', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (44, '/admin/menu/show', 'adminMenu::show', '菜单显示隐藏', '菜单显示隐藏', 1, '2026-09-04 10:55:34', '2026-09-04 10:55:34');
INSERT INTO `admin_api` VALUES (46, '/admin/account/info', 'adminAccount::info', '用户详情', '用户详情', 1, '2026-09-05 10:16:45', '2026-09-05 10:16:45');
INSERT INTO `admin_api` VALUES (47, '/admin/account/edit', 'adminAccount::edit', '用户账号编辑', '用户账号编辑', 1, '2026-09-05 10:17:57', '2026-09-05 10:17:57');
INSERT INTO `admin_api` VALUES (48, '/admin/account/password', 'adminAccount::password', '用户修改密码', '用户修改密码', 1, '2026-09-05 10:18:25', '2026-09-05 10:18:25');
INSERT INTO `admin_api` VALUES (49, '/admin/upload/list', 'adminUpload::list', '上传记录列表', '上传记录分页列表', 1, '2026-09-05 13:26:54', '2026-09-05 13:26:54');
INSERT INTO `admin_api` VALUES (50, '/admin/upload/upload', 'adminUpload::upload', '上传文件', '上传文件到本地或云存储', 1, '2026-09-05 13:26:54', '2026-09-05 13:26:54');
INSERT INTO `admin_api` VALUES (51, '/admin/upload/delete', 'adminUpload::delete', '删除上传', '删除上传记录与文件', 1, '2026-09-05 13:26:54', '2026-09-05 13:26:54');

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '路径',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限中文名称',
  `key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '菜单的唯一键名',
  `describe` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '路径图标',
  `sort` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序值',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '重定向路径',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '组件名称',
  `is_hide_in_menu` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否隐藏：0显示，1隐藏',
  `is_hide_children_in_menu` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否在children中隐藏：1隐藏，0显示',
  `is_enabled` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '1：启用，0禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `uk_key`(`key` ASC) USING BTREE,
  UNIQUE INDEX `uk_path`(`path` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
INSERT INTO `admin_menu` VALUES (1, 0, '/admin', '系统设置', 'Admin', '系统设置', 'SettingOutlined', 0, '/', '', 0, 0, 1, '2022-06-26 22:28:24', '2026-09-01 17:02:14');
INSERT INTO `admin_menu` VALUES (2, 1, '/admin/user', '账号管理', 'AdminUser', '账号列表', 'UserOutlined', 0, '/', '', 0, 1, 1, '2022-08-06 15:09:09', '2026-09-01 17:02:57');
INSERT INTO `admin_menu` VALUES (3, 1, '/admin/role', '角色管理', 'AdminRole', '角色列表', 'TeamOutlined', 0, '/', '', 0, 1, 1, '2022-08-06 15:08:22', '2026-09-01 17:28:43');
INSERT INTO `admin_menu` VALUES (4, 1, '/admin/menu', '菜单管理', 'AdminMenu', '菜单列表', 'MenuUnfoldOutlined', 0, '/', '', 0, 1, 1, '2022-06-26 22:29:08', '2026-09-01 17:25:25');
INSERT INTO `admin_menu` VALUES (5, 1, '/admin/permission', '权限管理', 'AdminPermission', '权限管理', 'SafetyCertificateOutlined', 0, '/', '', 0, 1, 1, '2022-08-06 15:10:20', '2026-09-01 17:29:01');
INSERT INTO `admin_menu` VALUES (6, 1, '/admin/api', '接口管理', 'AdminApi', '接口管理', 'ApiOutlined', 0, '/', '', 0, 0, 1, '2022-08-06 15:11:37', '2026-09-01 17:25:47');
INSERT INTO `admin_menu` VALUES (26, 0, '/account', '个人中心', 'Account', '', 'UserSwitchOutlined', 0, '/', '', 0, 0, 1, '2026-09-05 10:11:28', '2026-09-05 10:11:28');
INSERT INTO `admin_menu` VALUES (27, 1, '/admin/upload', '上传管理', 'AdminUpload', '上传文件记录', 'CloudUploadOutlined', 0, '/', '', 0, 0, 1, '2026-09-05 13:26:54', '2026-09-05 13:26:54');

-- ----------------------------
-- Table structure for admin_permission
-- ----------------------------
DROP TABLE IF EXISTS `admin_permission`;
CREATE TABLE `admin_permission`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `menu_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属菜单ID',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限唯一标识名称',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限显示名称',
  `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT 'view' COMMENT '权限的操作类型\r\nview：查看（只读）\r\nedit：编辑（读写）\r\ndelete：删除（彻底删除）',
  `describe` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限描述',
  `is_enabled` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否启用：1启用，0禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_key`(`key` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_permission
-- ----------------------------
INSERT INTO `admin_permission` VALUES (1, 2, 'AdminUserView', '查看账号', 'view', '查看账号列表与详情', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (2, 2, 'AdminUserAdd', '新建账号', 'edit', '创建账号', 1, '2026-09-04 11:20:20', '2026-09-04 19:37:12');
INSERT INTO `admin_permission` VALUES (3, 2, 'AdminUserEdit', '编辑账号', 'edit', '编辑账号与启用/禁用', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (4, 2, 'AdminUserResetPwd', '重置密码', 'edit', '重置账号登录密码', 1, '2026-09-04 11:20:20', '2026-09-04 19:37:12');
INSERT INTO `admin_permission` VALUES (5, 2, 'AdminUserBindRoles', '绑定角色', 'edit', '为账号绑定/解绑角色', 1, '2026-09-04 11:20:20', '2026-09-04 19:37:12');
INSERT INTO `admin_permission` VALUES (6, 2, 'AdminUserDelete', '删除账号', 'delete', '删除账号', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (7, 3, 'AdminRoleView', '查看角色', 'view', '查看角色列表与详情', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (8, 3, 'AdminRoleAdd', '新建角色', 'edit', '创建角色', 1, '2026-09-04 11:20:20', '2026-09-04 19:37:12');
INSERT INTO `admin_permission` VALUES (9, 3, 'AdminRoleEdit', '编辑角色', 'edit', '编辑角色与启用/禁用', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (10, 3, 'AdminRoleBindPermissions', '绑定权限', 'edit', '为角色绑定/解绑权限', 1, '2026-09-04 11:20:20', '2026-09-04 19:37:12');
INSERT INTO `admin_permission` VALUES (11, 3, 'AdminRoleDelete', '删除角色', 'delete', '删除角色', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (12, 4, 'AdminMenuView', '查看菜单', 'view', '查看菜单列表与详情', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (13, 4, 'AdminMenuAdd', '新建菜单', 'edit', '创建菜单', 1, '2026-09-04 11:20:20', '2026-09-04 19:37:12');
INSERT INTO `admin_permission` VALUES (14, 4, 'AdminMenuEdit', '编辑菜单', 'edit', '编辑/启停/显示隐藏菜单与权限配置', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (15, 4, 'AdminMenuDelete', '删除菜单', 'delete', '删除菜单', 1, '2026-09-04 11:20:20', '2026-09-04 11:20:20');
INSERT INTO `admin_permission` VALUES (16, 5, 'AdminPermissionView', '查看权限', 'view', '', 1, '2026-09-04 11:20:20', '2026-09-05 10:08:02');
INSERT INTO `admin_permission` VALUES (17, 5, 'AdminPermissionAdd', '新建权限', 'edit', '', 1, '2026-09-04 11:20:20', '2026-09-05 10:08:02');
INSERT INTO `admin_permission` VALUES (18, 5, 'AdminPermissionEdit', '编辑权限', 'edit', '', 1, '2026-09-04 11:20:20', '2026-09-05 10:08:02');
INSERT INTO `admin_permission` VALUES (19, 5, 'AdminPermissionDelete', '删除权限', 'delete', '', 1, '2026-09-04 11:20:20', '2026-09-05 10:08:02');
INSERT INTO `admin_permission` VALUES (20, 6, 'AdminApiView', '查看接口', 'view', '查看接口列表与详情', 1, '2026-09-04 11:20:20', '2026-09-04 19:18:38');
INSERT INTO `admin_permission` VALUES (21, 6, 'AdminApiAdd', '新建接口', 'edit', '创建接口', 1, '2026-09-04 11:20:20', '2026-09-04 19:37:12');
INSERT INTO `admin_permission` VALUES (22, 6, 'AdminApiEdit', '编辑接口', 'edit', '编辑接口与启用/禁用', 1, '2026-09-04 11:20:20', '2026-09-04 19:18:38');
INSERT INTO `admin_permission` VALUES (23, 6, 'AdminApiDelete', '删除接口', 'delete', '删除接口', 1, '2026-09-04 11:20:20', '2026-09-04 19:18:38');
INSERT INTO `admin_permission` VALUES (33, 26, 'AccountView', '查看', 'view', '', 1, '2026-09-05 10:18:55', '2026-09-05 10:18:55');
INSERT INTO `admin_permission` VALUES (34, 26, 'AccountEdit', '编辑', 'edit', '', 1, '2026-09-05 10:18:55', '2026-09-05 10:18:55');
INSERT INTO `admin_permission` VALUES (35, 27, 'AdminUploadView', '查看上传', 'view', '查看上传记录列表', 1, '2026-09-05 13:26:54', '2026-09-05 13:26:54');
INSERT INTO `admin_permission` VALUES (36, 27, 'AdminUploadEdit', '上传文件', 'edit', '上传文件', 1, '2026-09-05 13:26:54', '2026-09-05 13:26:54');
INSERT INTO `admin_permission` VALUES (37, 27, 'AdminUploadDelete', '删除上传', 'delete', '删除上传记录与文件', 1, '2026-09-05 13:26:54', '2026-09-05 13:26:54');

-- ----------------------------
-- Table structure for admin_permission_api
-- ----------------------------
DROP TABLE IF EXISTS `admin_permission_api`;
CREATE TABLE `admin_permission_api`  (
  `permission_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限ID',
  `api_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '接口ID',
  UNIQUE INDEX `uk_permission_api`(`permission_id` ASC, `api_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '接口权限关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_permission_api
-- ----------------------------
INSERT INTO `admin_permission_api` VALUES (1, 1);
INSERT INTO `admin_permission_api` VALUES (1, 3);
INSERT INTO `admin_permission_api` VALUES (2, 2);
INSERT INTO `admin_permission_api` VALUES (3, 4);
INSERT INTO `admin_permission_api` VALUES (3, 5);
INSERT INTO `admin_permission_api` VALUES (4, 43);
INSERT INTO `admin_permission_api` VALUES (5, 7);
INSERT INTO `admin_permission_api` VALUES (5, 40);
INSERT INTO `admin_permission_api` VALUES (6, 6);
INSERT INTO `admin_permission_api` VALUES (7, 8);
INSERT INTO `admin_permission_api` VALUES (7, 10);
INSERT INTO `admin_permission_api` VALUES (7, 15);
INSERT INTO `admin_permission_api` VALUES (7, 41);
INSERT INTO `admin_permission_api` VALUES (8, 9);
INSERT INTO `admin_permission_api` VALUES (9, 11);
INSERT INTO `admin_permission_api` VALUES (9, 12);
INSERT INTO `admin_permission_api` VALUES (10, 14);
INSERT INTO `admin_permission_api` VALUES (10, 15);
INSERT INTO `admin_permission_api` VALUES (10, 41);
INSERT INTO `admin_permission_api` VALUES (11, 13);
INSERT INTO `admin_permission_api` VALUES (12, 16);
INSERT INTO `admin_permission_api` VALUES (12, 17);
INSERT INTO `admin_permission_api` VALUES (12, 19);
INSERT INTO `admin_permission_api` VALUES (13, 18);
INSERT INTO `admin_permission_api` VALUES (14, 20);
INSERT INTO `admin_permission_api` VALUES (14, 21);
INSERT INTO `admin_permission_api` VALUES (14, 23);
INSERT INTO `admin_permission_api` VALUES (14, 31);
INSERT INTO `admin_permission_api` VALUES (14, 44);
INSERT INTO `admin_permission_api` VALUES (15, 22);
INSERT INTO `admin_permission_api` VALUES (16, 16);
INSERT INTO `admin_permission_api` VALUES (16, 25);
INSERT INTO `admin_permission_api` VALUES (16, 27);
INSERT INTO `admin_permission_api` VALUES (16, 39);
INSERT INTO `admin_permission_api` VALUES (17, 16);
INSERT INTO `admin_permission_api` VALUES (17, 24);
INSERT INTO `admin_permission_api` VALUES (17, 26);
INSERT INTO `admin_permission_api` VALUES (17, 39);
INSERT INTO `admin_permission_api` VALUES (18, 28);
INSERT INTO `admin_permission_api` VALUES (18, 29);
INSERT INTO `admin_permission_api` VALUES (18, 32);
INSERT INTO `admin_permission_api` VALUES (18, 39);
INSERT INTO `admin_permission_api` VALUES (19, 30);
INSERT INTO `admin_permission_api` VALUES (20, 33);
INSERT INTO `admin_permission_api` VALUES (20, 35);
INSERT INTO `admin_permission_api` VALUES (21, 34);
INSERT INTO `admin_permission_api` VALUES (22, 36);
INSERT INTO `admin_permission_api` VALUES (22, 37);
INSERT INTO `admin_permission_api` VALUES (23, 38);
INSERT INTO `admin_permission_api` VALUES (33, 46);
INSERT INTO `admin_permission_api` VALUES (34, 47);
INSERT INTO `admin_permission_api` VALUES (34, 48);
INSERT INTO `admin_permission_api` VALUES (35, 49);
INSERT INTO `admin_permission_api` VALUES (36, 50);
INSERT INTO `admin_permission_api` VALUES (37, 51);

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '角色名称',
  `describe` varchar(1024) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '角色描述',
  `modify_admin_id` int(10) UNSIGNED NOT NULL COMMENT '修改人',
  `create_admin_id` int(10) UNSIGNED NOT NULL COMMENT '创建人',
  `is_enabled` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '1：启用，0：禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '管理员角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, '超级管理员', '超级管理员', 9, 1, 1, '2022-08-11 02:09:58', '2026-08-31 17:13:13');
INSERT INTO `admin_role` VALUES (5, '查看', '', 1, 1, 1, '2026-09-03 13:31:58', '2026-09-05 10:04:31');

-- ----------------------------
-- Table structure for admin_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_permission`;
CREATE TABLE `admin_role_permission`  (
  `role_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `permission_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限ID',
  UNIQUE INDEX `uk_role_permission`(`role_id` ASC, `permission_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '角色权限关系表（包含菜单和权限）' ROW_FORMAT = DYNAMIC;

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
INSERT INTO `admin_role_permission` VALUES (1, 9);
INSERT INTO `admin_role_permission` VALUES (1, 10);
INSERT INTO `admin_role_permission` VALUES (1, 11);
INSERT INTO `admin_role_permission` VALUES (1, 12);
INSERT INTO `admin_role_permission` VALUES (1, 13);
INSERT INTO `admin_role_permission` VALUES (1, 14);
INSERT INTO `admin_role_permission` VALUES (1, 15);
INSERT INTO `admin_role_permission` VALUES (1, 16);
INSERT INTO `admin_role_permission` VALUES (1, 17);
INSERT INTO `admin_role_permission` VALUES (1, 18);
INSERT INTO `admin_role_permission` VALUES (1, 19);
INSERT INTO `admin_role_permission` VALUES (1, 20);
INSERT INTO `admin_role_permission` VALUES (1, 21);
INSERT INTO `admin_role_permission` VALUES (1, 22);
INSERT INTO `admin_role_permission` VALUES (1, 23);
INSERT INTO `admin_role_permission` VALUES (1, 33);
INSERT INTO `admin_role_permission` VALUES (1, 34);
INSERT INTO `admin_role_permission` VALUES (1, 35);
INSERT INTO `admin_role_permission` VALUES (1, 36);
INSERT INTO `admin_role_permission` VALUES (1, 37);
INSERT INTO `admin_role_permission` VALUES (5, 1);
INSERT INTO `admin_role_permission` VALUES (5, 7);
INSERT INTO `admin_role_permission` VALUES (5, 12);
INSERT INTO `admin_role_permission` VALUES (5, 16);
INSERT INTO `admin_role_permission` VALUES (5, 20);
INSERT INTO `admin_role_permission` VALUES (5, 33);
INSERT INTO `admin_role_permission` VALUES (5, 34);

-- ----------------------------
-- Table structure for admin_upload
-- ----------------------------
DROP TABLE IF EXISTS `admin_upload`;
CREATE TABLE `admin_upload`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `admin_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传者账号ID',
  `driver` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '存储驱动：local/aliyun/qcloud/s3',
  `upload_group` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '分组（上传路径前缀），如 admin/user',
  `object_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '存储对象键（相对路径）',
  `url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '访问链接（域名前缀+对象键）',
  `original_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '原始文件名',
  `new_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '新文件名',
  `ext` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '扩展名（小写无点）',
  `mime` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件MIME类型',
  `size` bigint(20) NOT NULL DEFAULT 0 COMMENT '文件大小（字节）',
  `md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件内容MD5',
  `upload_date` date NOT NULL COMMENT '上传日期，便于按日归档与统计',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group`(`upload_group` ASC) USING BTREE,
  INDEX `idx_date`(`upload_date` ASC) USING BTREE,
  INDEX `idx_admin_id`(`admin_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '上传文件记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_upload
-- ----------------------------
INSERT INTO `admin_upload` VALUES (8, 1, 'local', 'account', 'account/2026/09/20260905_174557_e4d7b125.ico', 'http://127.0.0.1:3000/upload/account/2026/09/20260905_174557_e4d7b125.ico', 'favicon.ico', '20260905_174557_e4d7b125.ico', 'ico', 'image/x-icon', 4286, '93939b1e4e2d9a9e34efe86c605e19ba', '2026-09-05', '2026-09-05 17:45:57', '2026-09-05 17:45:57');

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '管理员账号',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '管理昵称姓名',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '登录密码',
  `avatar` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户头像',
  `login_total` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录次数',
  `last_login_ip` json NOT NULL COMMENT '上次登录IP',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '上次登录时间',
  `is_enabled` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '账户状态：1正常，0：禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', '骑着八戒游天河', 'ddd@q1.com', '$2a$10$8DN3n4k4C7H3vrYoQ5AUA.5KrOJQYcUCaq5X7J94JmHRH.XeiEm.m', 'http://127.0.0.1:3000/upload/account/2026/09/20260905_174557_e4d7b125.ico', 222, '[\"127.0.0.1\", \"127.0.0.1\"]', '2026-09-05 17:46:33', 1, '2024-06-23 03:31:49', '2026-09-05 17:46:33');
INSERT INTO `admin_user` VALUES (19, 'wei', 'wei', '', '$2a$10$pXjeIKleDlZZPQ6WwER5JeLtn8SiTnDli2PcyRb21mKJuhpqNLDj2', '', 7, '[\"127.0.0.1\", \"127.0.0.1\"]', '2026-09-05 10:08:15', 1, '2026-09-03 13:35:35', '2026-09-05 10:08:15');

-- ----------------------------
-- Table structure for admin_user_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_role`;
CREATE TABLE `admin_user_role`  (
  `admin_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `role_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  UNIQUE INDEX `uk_admin_id_role_id`(`admin_id` ASC, `role_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '管理员-游戏-角色关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user_role
-- ----------------------------
INSERT INTO `admin_user_role` VALUES (5, 1);
INSERT INTO `admin_user_role` VALUES (7, 1);
INSERT INTO `admin_user_role` VALUES (9, 2);
INSERT INTO `admin_user_role` VALUES (10, 1);
INSERT INTO `admin_user_role` VALUES (10, 2);
INSERT INTO `admin_user_role` VALUES (19, 5);

SET FOREIGN_KEY_CHECKS = 1;
