package constant

const (
	AdministratorId = 1
)

// IsAdministrator 是否是超管
func IsAdministrator(adminId int) bool {
	return AdministratorId == adminId
}
