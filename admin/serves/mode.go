package serves

const (
	ModeDefaultValue = iota
	ModeLocalValue   = iota
	ModeDevelopValue
	ModeTestValue
	ModeProductionValue
)

const (
	ModeDefaultName    = ""
	ModeLocalName      = "local"
	ModeDevelopName    = "dev"
	ModeTestName       = "test"
	ModeProductionName = "prod"
)

var (
	ModeValueMap = map[int]string{
		ModeDefaultValue:    ModeDefaultName,
		ModeLocalValue:      ModeLocalName,
		ModeDevelopValue:    ModeDevelopName,
		ModeTestValue:       ModeTestName,
		ModeProductionValue: ModeProductionName,
	}

	ModeNameMap = map[string]int{
		ModeDefaultName:    ModeDefaultValue,
		ModeLocalName:      ModeLocalValue,
		ModeDevelopName:    ModeDevelopValue,
		ModeTestName:       ModeTestValue,
		ModeProductionName: ModeProductionValue,
	}
)
