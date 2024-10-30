package logic

type AdminLogic struct {
	AdminUser       IAdminUserLogic
	AdminAPI        IAdminAPILogic
	AdminMenu       IAdminMenuLogic
	AdminPermission IAdminPermissionLogic
	AdminRole       IAdminRoleLogic
}

func newAdminLogic() *AdminLogic {
	return &AdminLogic{
		AdminUser:       newAdminUserLogic(),
		AdminAPI:        newAdminAPILogic(),
		AdminMenu:       newAdminMenuLogic(),
		AdminPermission: newAdminPermissionLogic(),
		AdminRole:       newAdminRoleLogic(),
	}
}

var H = newAdminLogic()
