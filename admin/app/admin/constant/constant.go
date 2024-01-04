package constant

const (
	Administer = 1 // 超管ID
)

func IsAdministrator(adminId int) bool {
	return adminId == Administer
}
