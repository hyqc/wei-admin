package constant

const (
	AdministerId = 1 // 超管ID
)

func IsAdministrator(adminId int) bool {
	return adminId == AdministerId
}
