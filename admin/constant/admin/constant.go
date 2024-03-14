package admin

const (
	AdministerId int32 = 1 // 超管ID
)

func IsAdministrator(adminId int32) bool {
	return adminId == AdministerId
}
