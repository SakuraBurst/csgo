package memory

const (
	MAX_MODULE_NAME32        = 255
	MAX_PATH                 = 260
	TH32CS_SNAPHEAPLIST      = 0x00000001
	TH32CS_SNAPPROCESS       = 0x00000002
	TH32CS_SNAPTHREAD        = 0x00000004
	TH32CS_SNAPMODULE        = 0x00000008
	TH32CS_SNAPMODULE32      = 0x00000010
	TH32CS_INHERIT           = 0x80000000
	TH32CS_SNAPALL           = TH32CS_SNAPHEAPLIST | TH32CS_SNAPMODULE | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD
	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	SYNCHRONIZE              = 0x00100000
	PROCESS_ALL_ACCESS       = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xffff
)
