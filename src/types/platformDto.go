package types

type PlatformName string

type Platform interface {
	// platform name
	GetPlatformName() PlatformName

	// code name of the file
	GetCodeName() string
}
