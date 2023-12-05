package constant

const (
	AdministratorId int32 = 1
)

// IsAdministrator 是否是超管
func IsAdministrator(adminId int32) bool {
	return AdministratorId == adminId
}
