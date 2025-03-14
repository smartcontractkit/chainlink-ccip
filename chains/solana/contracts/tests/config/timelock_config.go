package config

var (
	TimelockProgram = GetProgramID("timelock")
	// [0,0,0,...'t','e','s','t','-','t','i','m','e','l','o','c','k']
	TestTimelockID = [32]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b,
	}
	NumAccountsPerRole      = 63 // max 64 accounts per role(access list) * 4 - 1(to keep test accounts fits single funding)
	BatchAddAccessChunkSize = 24
	MinDelay                = uint64(1)
	TimelockEmptyOpID       = [32]byte{}
	MaxFunctionSelectorLen  = 128

	// operational constraints
	AppendIxDataChunkSize = 491
)
