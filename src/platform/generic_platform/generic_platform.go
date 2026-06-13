package generic_platform

import (
	"prepare-code/src/constants"
	"prepare-code/src/types"
)

type GenericInstance struct {
	name     types.PlatformName
	codeName string
}

// GetCodeName implements types.Platform.
func (g GenericInstance) GetCodeName() string {
	return g.codeName
}

// GetPlatformName implements types.Platform.
func (g GenericInstance) GetPlatformName() types.PlatformName {
	return constants.PlatformGeneric
}

func GetPlatform() types.Instance {
	return GenericInstance{
		name:     constants.PlatformGeneric,
		codeName: "generic",
	}
}
