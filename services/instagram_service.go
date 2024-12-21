package services

import (
	"fmt"
	"github.com/google/uuid"
	"ms-live-sells/models"
	"ms-live-sells/provider"
	"ms-live-sells/repositories"
	"ms-live-sells/social"
	"ms-live-sells/utils"
	"os"
	"time"
)

type InstagramService struct {
	Provider *provider.InstagramProvider
	UserRepo *repositories.UserRepository

	SocialService *social.SocialService
}

func (i *InstagramService) GenerateLongLivedAccessToken() error {
	return i.Provider.GenerateLongLivedAccessToken()
}

func (i *InstagramService) GetUserByUsername(username string) (*models.UserInfo, error) {
	return i.Provider.GetUserByUsername(username)
}

func (i *InstagramService) GetUserMedias(userID string) ([]models.InstagramMedia, error) {
	return i.Provider.GetMedias(userID)
}

func (i *InstagramService) GetLiveMedias(userID string) ([]models.InstagramMedias, error) {
	return i.Provider.GetLiveMedia(userID)
}

func (i *InstagramService) GetMediaComments(mediaID string) (*models.InstagramCommentResponse, error) {
	return i.Provider.GetComments(mediaID)
}

// StartInstagramMonitoring begins the monitoring process for the instagram user by his ID
func (i *InstagramService) StartInstagramMonitoring(userID uuid.UUID) error {
	// Retrieve the user from the repository
	user, err := i.UserRepo.GetUsersWithInstagramByID(userID)
	if err != nil {
		return fmt.Errorf("error retrieving user: %v", err)
	}

	// Initialize Instagram service
	instagramService := InstagramService{
		Provider: &provider.InstagramProvider{
			LongLivedToken: os.Getenv("LONGLIVEDTOKEN"),
		},
	}

	err = instagramService.GenerateLongLivedAccessToken()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	//socialNetworkID := "17841401499820518"
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Monitor while the user has an active live
	for {
		select {
		case <-ticker.C:
			// Fetch live media for the user
			liveMedia, err := instagramService.GetLiveMedias(user.UsersSocialNetwork[0].SocialNetworkID)
			if err != nil {
				fmt.Printf("Error fetching live media: %v\n", err)
				continue
			}

			// Check if the live is still active
			if len(liveMedia) == 0 {
				fmt.Printf("Live ended for user: %s\n", user.UsersSocialNetwork[0].SocialNetworkID)
				return nil
			}

			// Process comments
			for _, media := range liveMedia {
				fmt.Printf("Live Media ID: %s, Username: %s\n", media.ID, media.Username)

				for _, comment := range media.Comments.Data {
					fmt.Printf("Username: %s - Comment: %s\n", comment.Username, comment.Text)
					productCode := utils.ExtractProductCode(comment.Text)
					err = i.SocialService.ProcessProductCode(productCode, media.Username, comment.Username, "instagram")
					if err != nil {
						return err
					}
				}
			}
		}
	}
}
