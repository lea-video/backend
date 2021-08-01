package types

type Format struct {
	Filename string `json:"filename"`
	Filesize uint   `json:"filesize"`
	Language string `json:"lang"`
}

type AudioFormat struct {
	*Format
	Bitrate  uint   `json:"bitrate"`
	Encoding string `json:"encoding"`
}

type VideoFormat struct {
	*Format
	Bitrate  uint   `json:"bitrate"`
	Encoding string `json:"encoding"`
}

type SubtitleFormat struct {
	*Format
}
