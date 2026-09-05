package logic

type AdminLogic struct {
	AdminUser       IAdminUserLogic
	AdminAPI        IAdminAPILogic
	AdminMenu       IAdminMenuLogic
	AdminPermission IAdminPermissionLogic
	AdminRole       IAdminRoleLogic
	AdminUpload     IAdminUploadLogic
}

func newAdminLogic() *AdminLogic {
	return &AdminLogic{
		AdminUser:       newAdminUserLogic(),
		AdminAPI:        newAdminAPILogic(),
		AdminMenu:       newAdminMenuLogic(),
		AdminPermission: newAdminPermissionLogic(),
		AdminRole:       newAdminRoleLogic(),
		AdminUpload:     newAdminUploadLogic(),
	}
}

var H = newAdminLogic()
