package model

type BannerResourceGetResp struct {
	CurrentPage int32 `json:"current_page"`
	TotalCount int32 `json:"total_count"`
	List []*BannerResourceGetResp_List `json:"list"`
}

type BannerResourceGetResp_List struct {
	Pic string `json:"pic"`
}