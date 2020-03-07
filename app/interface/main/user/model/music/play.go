package music

type PlaySongResp struct {
	Data []*PlaySong `json:"data"`
}

type PlaySong struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Duration int64 `json:"duration"`
	Url string `json:"url"`
}