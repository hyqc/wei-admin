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
INSERT INTO `admin_api` VALUES (1, '/admin/user/list', 'adminUser::list', '账号列表', '账号列表', 1, '2024-06-23 13:53:39', '2024-06-23 13:57:51');
INSERT INTO `admin_api` VALUES (2, '/admin/user/add', 'adminUser::add', '账号创建', '账号创建', 1, '2022-08-07 11:01:21', '2022-08-07 11:01:21');
INSERT INTO `admin_api` VALUES (3, '/admin/user/info', 'adminUser::detail', '账号详情', '账号详情', 1, '2022-08-07 10:59:05', '2022-08-07 10:59:05');
INSERT INTO `admin_api` VALUES (4, '/admin/user/edit', 'adminUser::edit', '账号编辑', '账号编辑', 1, '2022-08-07 11:24:04', '2022-08-07 11:24:04');
INSERT INTO `admin_api` VALUES (5, '/admin/user/enable', 'adminUser::enable', '账号启用禁用', '账号启用禁用', 1, '2022-08-07 11:31:53', '2022-08-07 11:31:53');
INSERT INTO `admin_api` VALUES (6, '/admin/user/delete', 'adminUser::delete', '账号删除', '账号删除', 1, '2022-08-07 11:31:58', '2022-08-07 11:31:58');
INSERT INTO `admin_api` VALUES (7, '/admin/user/bind_roles', 'adminUser::bindRoles', '账号绑定角色', '账号绑定角色', 1, '2022-08-07 11:37:20', '2022-08-07 11:37:20');
INSERT INTO `admin_api` VALUES (8, '/admin/role/list', 'adminRole::list', '角色列表', '角色列表', 1, '2022-08-07 11:48:30', '2022-08-07 11:48:30');
INSERT INTO `admin_api` VALUES (9, '/admin/role/add', 'adminRole::add', '角色创建', '', 1, '2022-08-07 11:48:55', '2022-08-07 11:48:56');
INSERT INTO `admin_api` VALUES (10, '/admin/role/info', 'adminRole::detail', '角色详情', '', 1, '2022-08-07 11:50:29', '2022-08-07 11:50:29');
INSERT INTO `admin_api` VALUES (11, '/admin/role/edit', 'adminRole::edit', '角色编辑', '', 1, '2022-08-07 12:14:07', '2022-08-07 12:14:07');
INSERT INTO `admin_api` VALUES (12, '/admin/role/enable', 'adminRole::enable', '角色禁用启用', '', 1, '2022-08-07 12:14:38', '2022-08-07 12:14:38');
INSERT INTO `admin_api` VALUES (13, '/admin/role/delete', 'adminRole::delete', '角色删除', '', 1, '2022-08-07 12:15:21', '2022-08-07 12:15:21');
INSERT INTO `admin_api` VALUES (14, '/admin/role/bind_permissions', 'adminRole::bindPermissions', '角色绑定权限', '', 1, '2022-08-07 12:16:03', '2022-08-07 12:16:03');
INSERT INTO `admin_api` VALUES (15, '/admin/role/permissions', 'adminRole::permissions', '角色权限列表', '', 1, '2022-08-07 12:16:56', '2022-08-07 12:16:56');
INSERT INTO `admin_api` VALUES (16, '/admin/menu/tree', 'adminMenu::tree', '菜单树', '', 1, '2022-08-07 12:27:17', '2022-08-07 12:27:17');
INSERT INTO `admin_api` VALUES (17, '/admin/menu/list', 'adminMenu::list', '菜单列表', '', 1, '2022-08-07 12:27:40', '2022-08-07 12:27:40');
INSERT INTO `admin_api` VALUES (18, '/admin/menu/add', 'adminMenu::add', '菜单创建', '', 1, '2022-08-07 12:28:02', '2022-08-07 12:28:02');
INSERT INTO `admin_api` VALUES (19, '/admin/menu/info', 'adminMenu::detail', '菜单详情', '', 1, '2022-08-07 12:28:17', '2022-08-07 12:28:17');
INSERT INTO `admin_api` VALUES (20, '/admin/menu/edit', 'adminMenu::edit', '菜单编辑', '', 1, '2022-08-07 12:28:38', '2022-08-07 12:28:38');
INSERT INTO `admin_api` VALUES (21, '/admin/menu/enable', 'adminMenu::enable', '菜单禁用启用', '', 1, '2022-08-07 12:28:52', '2022-08-07 12:28:52');
INSERT INTO `admin_api` VALUES (22, '/admin/menu/delete', 'adminMenu::delete', '菜单删除', '', 1, '2022-08-07 12:29:06', '2022-08-07 12:29:06');
INSERT INTO `admin_api` VALUES (23, '/admin/menu/permissions', 'adminMenu::permissions', '菜单权限', '', 1, '2022-08-07 14:18:02', '2022-08-07 14:18:03');
INSERT INTO `admin_api` VALUES (24, '/admin/menu/pages', 'adminMenu::pages', '菜单页面列表', '', 1, '2022-08-07 12:30:16', '2022-08-07 12:30:16');
INSERT INTO `admin_api` VALUES (25, '/admin/permission/list', 'adminPermission::list', '权限列表', '', 1, '2022-08-07 12:31:00', '2022-08-07 12:31:00');
INSERT INTO `admin_api` VALUES (26, '/admin/permission/add', 'adminPermission::add', '权限创建', '', 1, '2022-08-07 12:31:15', '2022-08-07 12:31:15');
INSERT INTO `admin_api` VALUES (27, '/admin/permission/info', 'adminPermission::detail', '权限详情', '', 1, '2022-08-07 12:31:29', '2022-08-07 12:31:29');
INSERT INTO `admin_api` VALUES (28, '/admin/permission/edit', 'adminPermission::edit', '权限编辑', '', 1, '2022-08-07 12:36:02', '2022-08-07 12:36:02');
INSERT INTO `admin_api` VALUES (29, '/admin/permission/enable', 'adminPermission::enable', '权限禁用启用', '', 1, '2022-08-07 12:36:18', '2022-08-07 12:36:18');
INSERT INTO `admin_api` VALUES (30, '/admin/permission/delete', 'adminPermission::delete', '权限删除', '', 1, '2022-08-07 12:36:28', '2022-08-07 12:36:28');
INSERT INTO `admin_api` VALUES (31, '/admin/permission/add_menu_permissions', 'adminPermission::addMenuPermissions', '权限指定菜单批量创建权限', '给指定的菜单创建查看，编辑，删除权限', 1, '2022-08-07 12:37:22', '2022-08-07 12:37:22');
INSERT INTO `admin_api` VALUES (32, '/admin/permission/bind_apis', 'adminPermission::bindApis', '权限绑定接口', '', 1, '2022-08-07 12:37:40', '2022-08-07 12:37:40');
INSERT INTO `admin_api` VALUES (33, '/admin/api/list', 'adminApi::list', '接口列表', '', 1, '2022-08-07 12:38:24', '2022-08-07 12:38:24');
INSERT INTO `admin_api` VALUES (34, '/admin/api/add', 'adminApi::add', '接口创建', '', 1, '2022-08-07 12:38:44', '2022-08-07 12:38:44');
INSERT INTO `admin_api` VALUES (35, '/admin/api/info', 'adminApi::detail', '接口详情', '', 1, '2022-08-07 12:38:59', '2022-08-07 12:38:59');
INSERT INTO `admin_api` VALUES (36, '/admin/api/edit', 'adminApi::edit', '接口编辑', '', 1, '2022-08-07 12:39:10', '2022-08-07 12:39:10');
INSERT INTO `admin_api` VALUES (37, '/admin/api/enable', 'adminApi::enable', '接口禁用启用', '', 1, '2022-08-07 12:39:23', '2022-08-07 12:39:23');
INSERT INTO `admin_api` VALUES (38, '/admin/api/delete', 'adminApi::delete', '接口删除', '', 1, '2022-08-07 12:39:34', '2022-08-07 12:39:34');
INSERT INTO `admin_api` VALUES (39, '/admin/api/all', 'adminApi::all', '接口全部', '全部有效接口列表', 1, '2022-08-07 12:40:07', '2022-08-07 12:40:07');
INSERT INTO `admin_api` VALUES (40, '/admin/role/all', 'adminRole::all', '角色全部', '全部有效的角色列表', 1, '2022-08-07 16:26:45', '2022-08-07 16:26:45');
INSERT INTO `admin_api` VALUES (41, '/admin/menu/modes', 'adminMenu::mode', '菜单页面权限列表', '', 1, '2022-08-08 03:18:10', '2022-08-08 03:18:10');
INSERT INTO `admin_api` VALUES (43, '/admin/user/edit_pwd', 'adminUser::editPwd', '账号重置密码', '账号重置密码', 1, '2026-09-04 10:00:00', '2026-09-04 10:00:00');
INSERT INTO `admin_api` VALUES (44, '/admin/menu/show', 'adminMenu::show', '菜单显示隐藏', '菜单显示隐藏', 1, '2026-09-04 10:00:00', '2026-09-04 10:00:00');

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

-- ----------------------------
-- Table structure for admin_permission
-- ----------------------------
DROP TABLE IF EXISTS `admin_permission`;
CREATE TABLE `admin_permission`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `menu_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属菜单ID',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限唯一标识名称（与前端按钮权限码一致）',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限显示名称',
  `type` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL DEFAULT 'view' COMMENT '权限的动作类型\r\nview：查看\r\nadd：新增\r\nedit：编辑\r\nenable：启停\r\nbind：绑定\r\nreset：重置\r\ndelete：删除',
  `describe` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限描述',
  `is_enabled` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否启用：1启用，0禁用',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_key`(`key` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_permission
-- ----------------------------
INSERT INTO `admin_permission` VALUES (1, 2, 'AdminUserView', '查看账号', 'view', '查看账号列表与详情', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (2, 2, 'AdminUserAdd', '新建账号', 'edit', '创建账号', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (3, 2, 'AdminUserEdit', '编辑账号', 'edit', '编辑账号与启用/禁用', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (4, 2, 'AdminUserResetPwd', '重置密码', 'edit', '重置账号登录密码', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (5, 2, 'AdminUserBindRoles', '绑定角色', 'edit', '为账号绑定/解绑角色', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (6, 2, 'AdminUserDelete', '删除账号', 'delete', '删除账号', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (7, 3, 'AdminRoleView', '查看角色', 'view', '查看角色列表与详情', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (8, 3, 'AdminRoleAdd', '新建角色', 'edit', '创建角色', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (9, 3, 'AdminRoleEdit', '编辑角色', 'edit', '编辑角色与启用/禁用', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (10, 3, 'AdminRoleBindPermissions', '绑定权限', 'edit', '为角色绑定/解绑权限', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (11, 3, 'AdminRoleDelete', '删除角色', 'delete', '删除角色', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (12, 4, 'AdminMenuView', '查看菜单', 'view', '查看菜单列表与详情', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (13, 4, 'AdminMenuAdd', '新建菜单', 'edit', '创建菜单', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (14, 4, 'AdminMenuEdit', '编辑菜单', 'edit', '编辑/启停/显示隐藏菜单与权限配置', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (15, 4, 'AdminMenuDelete', '删除菜单', 'delete', '删除菜单', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (16, 5, 'AdminPermissionView', '查看权限', 'view', '查看权限列表与详情', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (17, 5, 'AdminPermissionAdd', '新建权限', 'edit', '创建权限', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (18, 5, 'AdminPermissionEdit', '编辑权限', 'edit', '编辑/启停权限与绑定接口', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (19, 5, 'AdminPermissionDelete', '删除权限', 'delete', '删除权限', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (20, 6, 'AdminApiView', '查看接口', 'view', '查看接口列表与详情', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (21, 6, 'AdminApiAdd', '新建接口', 'edit', '创建接口', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (22, 6, 'AdminApiEdit', '编辑接口', 'edit', '编辑接口与启用/禁用', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_permission` VALUES (23, 6, 'AdminApiDelete', '删除接口', 'delete', '删除接口', 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');

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
INSERT INTO `admin_permission_api` VALUES (16, 25);
INSERT INTO `admin_permission_api` VALUES (16, 27);
INSERT INTO `admin_permission_api` VALUES (16, 16);
INSERT INTO `admin_permission_api` VALUES (17, 26);
INSERT INTO `admin_permission_api` VALUES (17, 24);
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
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '管理员角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, '超级管理员', '超级管理员', 9, 1, 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');
INSERT INTO `admin_role` VALUES (5, '测试', '', 1, 1, 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');

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
INSERT INTO `admin_role_permission` VALUES (5, 1);
INSERT INTO `admin_role_permission` VALUES (5, 2);
INSERT INTO `admin_role_permission` VALUES (5, 3);
INSERT INTO `admin_role_permission` VALUES (5, 4);
INSERT INTO `admin_role_permission` VALUES (5, 5);
INSERT INTO `admin_role_permission` VALUES (5, 6);
INSERT INTO `admin_role_permission` VALUES (5, 7);
INSERT INTO `admin_role_permission` VALUES (5, 8);
INSERT INTO `admin_role_permission` VALUES (5, 9);
INSERT INTO `admin_role_permission` VALUES (5, 10);
INSERT INTO `admin_role_permission` VALUES (5, 11);
INSERT INTO `admin_role_permission` VALUES (5, 12);
INSERT INTO `admin_role_permission` VALUES (5, 13);
INSERT INTO `admin_role_permission` VALUES (5, 14);
INSERT INTO `admin_role_permission` VALUES (5, 15);
INSERT INTO `admin_role_permission` VALUES (5, 16);
INSERT INTO `admin_role_permission` VALUES (5, 17);
INSERT INTO `admin_role_permission` VALUES (5, 18);
INSERT INTO `admin_role_permission` VALUES (5, 19);
INSERT INTO `admin_role_permission` VALUES (5, 20);
INSERT INTO `admin_role_permission` VALUES (5, 21);
INSERT INTO `admin_role_permission` VALUES (5, 22);
INSERT INTO `admin_role_permission` VALUES (5, 23);

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
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', '骑着八戒游天河', 'ddd@q1.com', '$2a$10$8DN3n4k4C7H3vrYoQ5AUA.5KrOJQYcUCaq5X7J94JmHRH.XeiEm.m', '', 39, '["127.0.0.1", "127.0.0.1"]', '2024-06-23 11:31:49', 1, '2024-06-23 03:31:49', '2024-06-23 11:31:49');
INSERT INTO `admin_user` VALUES (19, 'wei', 'wei', '', '$2a$10$pXjeIKleDlZZPQ6WwER5JeLtn8SiTnDli2PcyRb21mKJuhpqNLDj2', '', 0, '[]', NULL, 1, '2026-09-04 12:00:00', '2026-09-04 12:00:00');

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
INSERT INTO `admin_user_role` VALUES (19, 5);

SET FOREIGN_KEY_CHECKS = 1;
