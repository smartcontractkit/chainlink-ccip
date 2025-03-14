package config

var (
	AccessControllerProgram = GetProgramID("access_controller")
	AccSpace                = uint64(8 + 32 + 32 + ((32 * 64) + 8)) // discriminator + owner + proposed owner + access_list (64 max addresses + length)
)
