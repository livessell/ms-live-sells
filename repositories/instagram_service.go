package repositories

import (
	"ms-live-sells/models"
	"ms-live-sells/provider"
)

type InstagramService struct {
	Provider provider.InstagramProvier
}

func (i *InstagramService) GetUserMedias(userID string) ([]models.InstagramMedia, error) {
	return i.Provider.GetMedias(userID)
}

func (i *InstagramService) GetMediaComments(mediaID string) (*models.InstagramCommentResponse, error) {
	return i.Provider.GetComments(mediaID)
}
