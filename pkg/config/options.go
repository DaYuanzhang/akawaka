package config

type Options struct {
	DirPath         string
	Keywords_File   string
	Keyword         string
	Keywords        []string
	Extension       string
	Extensions      []string
	Extensions_File string
	Is_Filename     bool
	Count           int
	CurrentCount    int
	Thread          int
	Verbose         bool
}
