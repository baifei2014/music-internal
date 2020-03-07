package music

type SongResourceGetDetailResp struct {
	Privileges []*Privilege `json:"privileges"`
	Songs []*Song `json:"songs"`
}

type Song struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Duration int64 `json:"duration"`
	Url string `json:"url"`
	Al *Album `json:"al"`
	Artists []*Artist `json:"ar"`
}

type Album struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	PicUrl string `json:"picUrl"`
}

type AlbumGroup struct {
	Album *Album
}

type Artist struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

type ArtistGroup struct {
	Artists []*Artist
}

type Privilege struct {
	Id int64 `json:"id"`
	Maxbr int64 `json:"maxbr"`
	St int64 `json:"st"`
}

type SongPlayerInfoResp struct {
	Data []*SongPlayerInfo `json:"data"`
	Code int64 `json:"code"`
}

type SongPlayerInfo struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Duration int64 `json:"duration"`
	Url string `json:"url"`
}