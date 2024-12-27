package repositories

import (
	"github.com/google/uuid"
	"ms-live-sells/database"
	"ms-live-sells/models"
	"time"
)

type LiveRepository struct {
}

func (r *LiveRepository) FindByID(liveID uuid.UUID) (*models.Live, error) {
	var live models.Live
	err := database.DB.First(&live, "id = ?", liveID).Error
	if err != nil {
		return nil, err
	}
	return &live, nil
}

// UpdateLiveStatusToStart update the live status
func (r *LiveRepository) UpdateLiveStatusToStart(live *models.Live) error {
	live.Status = "start"       // CHange the live status to start
	live.StartTime = time.Now() // Setting the live starting time
	return database.DB.Save(live).Error
}

// UpdateLiveStatusToEnd update the live status
func (r *LiveRepository) UpdateLiveStatusToEnd(live *models.Live) error {
	live.Status = "ended"     // CHange the live status to start
	live.EndTime = time.Now() // Setting the live ending time
	return database.DB.Save(live).Error
}
