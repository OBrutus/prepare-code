package types

type PlatformName string

type Instance interface {
	// platform name
	GetPlatformName() PlatformName

	// code name of the file
	GetCodeName() string
}
