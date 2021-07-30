package types

type Format struct {
	Filename string
	Filesize uint
	Language string
}

type AudioFormat struct {
	Bitrate  uint
	Encoding string
	Format
}

type VideoFormat struct {
	Bitrate  uint
	Encoding string
	Format
}

type SubtitleFormat struct {
	Format
}
