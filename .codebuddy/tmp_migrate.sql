-- 方案A：权限点=页面操作。动作类型：view/add/edit/enable/bind/reset/delete
START TRANSACTION;

ALTER TABLE `admin_permission` DROP INDEX `uk_permission`;

DELETE FROM `admin_permission_api`;
DELETE FROM `admin_role_permission`;
DELETE FROM `admin_permission`;
ALTER TABLE `admin_permission` AUTO_INCREMENT = 24;

INSERT INTO `admin_permission` (`id`,`menu_id`,`key`,`name`,`type`,`describe`,`is_enabled`,`created_at`,`updated_at`) VALUES
(1,2,'AdminUserView','查看账号','view','查看账号列表与详情',1,NOW(),NOW()),
(2,2,'AdminUserAdd','新建账号','add','创建账号',1,NOW(),NOW()),
(3,2,'AdminUserEdit','编辑账号','edit','编辑账号与启用/禁用',1,NOW(),NOW()),
(4,2,'AdminUserResetPwd','重置密码','reset','重置账号登录密码',1,NOW(),NOW()),
(5,2,'AdminUserBindRoles','绑定角色','bind','为账号绑定/解绑角色',1,NOW(),NOW()),
(6,2,'AdminUserDelete','删除账号','delete','删除账号',1,NOW(),NOW()),
(7,3,'AdminRoleView','查看角色','view','查看角色列表与详情',1,NOW(),NOW()),
(8,3,'AdminRoleAdd','新建角色','add','创建角色',1,NOW(),NOW()),
(9,3,'AdminRoleEdit','编辑角色','edit','编辑角色与启用/禁用',1,NOW(),NOW()),
(10,3,'AdminRoleBindPermissions','绑定权限','bind','为角色绑定/解绑权限',1,NOW(),NOW()),
(11,3,'AdminRoleDelete','删除角色','delete','删除角色',1,NOW(),NOW()),
(12,4,'AdminMenuView','查看菜单','view','查看菜单列表与详情',1,NOW(),NOW()),
(13,4,'AdminMenuAdd','新建菜单','add','创建菜单',1,NOW(),NOW()),
(14,4,'AdminMenuEdit','编辑菜单','edit','编辑/启停/显示隐藏菜单与权限配置',1,NOW(),NOW()),
(15,4,'AdminMenuDelete','删除菜单','delete','删除菜单',1,NOW(),NOW()),
(16,5,'AdminPermissionView','查看权限','view','查看权限列表与详情',1,NOW(),NOW()),
(17,5,'AdminPermissionAdd','新建权限','add','创建权限',1,NOW(),NOW()),
(18,5,'AdminPermissionEdit','编辑权限','edit','编辑/启停权限与绑定接口',1,NOW(),NOW()),
(19,5,'AdminPermissionDelete','删除权限','delete','删除权限',1,NOW(),NOW()),
(20,6,'AdminApiView','查看接口','view','查看接口列表与详情',1,NOW(),NOW()),
(21,6,'AdminApiAdd','新建接口','add','创建接口',1,NOW(),NOW()),
(22,6,'AdminApiEdit','编辑接口','edit','编辑接口与启用/禁用',1,NOW(),NOW()),
(23,6,'AdminApiDelete','删除接口','delete','删除接口',1,NOW(),NOW());

INSERT INTO `admin_permission_api` (`permission_id`,`api_id`) VALUES
(1,1),(1,3),
(2,2),
(3,4),(3,5),
(4,43),
(5,7),(5,40),
(6,6),
(7,8),(7,10),(7,15),(7,41),
(8,9),
(9,11),(9,12),
(10,14),(10,15),(10,41),
(11,13),
(12,16),(12,17),(12,19),
(13,18),
(14,20),(14,21),(14,23),(14,31),(14,44),
(15,22),
(16,25),(16,27),(16,16),
(17,26),(17,24),
(18,28),(18,29),(18,32),(18,39),
(19,30),
(20,33),(20,35),
(21,34),
(22,36),(22,37),
(23,38);

-- 超级管理员(1)与测试(5)角色均授予全部权限点（测试账号 wei 用于联调，需全量可用）
INSERT INTO `admin_role_permission` (`role_id`,`permission_id`) VALUES
(1,1),(1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8),(1,9),(1,10),(1,11),(1,12),
(1,13),(1,14),(1,15),(1,16),(1,17),(1,18),(1,19),(1,20),(1,21),(1,22),(1,23),
(5,1),(5,2),(5,3),(5,4),(5,5),(5,6),(5,7),(5,8),(5,9),(5,10),(5,11),(5,12),
(5,13),(5,14),(5,15),(5,16),(5,17),(5,18),(5,19),(5,20),(5,21),(5,22),(5,23);

COMMIT;
