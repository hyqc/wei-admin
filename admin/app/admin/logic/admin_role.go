package logic

type AdminRoleLogic struct {
}

type IAdminRoleLogic interface {
}

func newAdminRoleLogic() IAdminRoleLogic {
	return &AdminRoleLogic{}
}
