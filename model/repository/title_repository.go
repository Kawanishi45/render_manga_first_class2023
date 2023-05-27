package repository

import (
	"github.com/render_manga_api/model/entity"
)

const (
	PickUp = "pick_up"
	Normal = "normal"
)

type TitleRepository interface {
	GetTitles() (todos []entity.TitleEntity, err error)
}

type titleRepository struct {
}

func NewTitleRepository() TitleRepository {
	return &titleRepository{}
}

func (t *titleRepository) GetTitles() (titles []entity.TitleEntity, err error) {
	return []entity.TitleEntity{{
		Id:           1,
		Name:         "title1",
		Type:         PickUp,
		ThumbnailUrl: "https://placehold.jp/400x250.png",
	}, {
		Id:           2,
		Name:         "title2",
		Type:         PickUp,
		ThumbnailUrl: "https://placehold.jp/400x250.png",
	}, {
		Id:           3,
		Name:         "title3",
		Type:         Normal,
		ThumbnailUrl: "https://placehold.jp/400x250.png",
	}, {
		Id:           4,
		Name:         "title4",
		Type:         Normal,
		ThumbnailUrl: "https://placehold.jp/400x250.png",
	}, {
		Id:           5,
		Name:         "title5",
		Type:         Normal,
		ThumbnailUrl: "https://placehold.jp/400x250.png",
	}}, nil
}
