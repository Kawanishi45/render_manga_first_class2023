package repository

type PageRepository interface {
	GetPageImageUrlsByTitleId(titleId int) (pageImages []string, err error)
}

type pageRepository struct {
}

func NewPageRepository() PageRepository {
	return &pageRepository{}
}

func (c pageRepository) GetPageImageUrlsByTitleId(titleId int) (pageImageUrls []string, err error) {
	return []string{
		"https://placehold.jp/500x800.png",
		"https://placehold.jp/500x800.png",
		"https://placehold.jp/500x800.png",
	}, nil
}
