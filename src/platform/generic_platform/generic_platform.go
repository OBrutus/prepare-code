package generic_platform

import (
	"prepare-code/src/constants"
	"prepare-code/src/types"
)

type GenericPlatform struct {
	name     types.PlatformName
	codeName string
}

// GetCodeName implements types.Platform.
func (g GenericPlatform) GetCodeName() string {
	return g.codeName
}

// GetPlatformName implements types.Platform.
func (g GenericPlatform) GetPlatformName() types.PlatformName {
	return constants.PlatformGeneric
}

func GetPlatform() types.Platform {
	return GenericPlatform{
		name:     constants.PlatformGeneric,
		codeName: "generic",
	}
}
