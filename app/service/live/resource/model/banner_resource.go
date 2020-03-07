package model

type BannerResource struct {
	Pic string
}

func (c BannerResource) TableName() string {
	return "banner_resource"
}