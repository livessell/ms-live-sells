package models

import "github.com/google/uuid"

// InstagramMedia representa uma mídia do Instagram
type InstagramMedia struct {
	ID               string `json:"id"`
	MediaProductType string `json:"media_product_type"`
	CreatedTime      string `json:"created_time"`
}

// InstagramResponse representa a resposta da API
type InstagramResponse struct {
	Data []InstagramMedia `json:"data"`
}

type InstagramComment struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

type InstagramCommentResponse struct {
	Data []InstagramComment `json:"data"`
}

// InstagramMedias represents an Instagram Media object.
type InstagramMedias struct {
	ID               string `json:"id"`
	MediaType        string `json:"media_type"`
	MediaProductType string `json:"media_product_type"`
	Owner            struct {
		ID string `json:"id"`
	} `json:"owner"`
	Username string `json:"username"`
	Comments struct {
		Data []InstagramComments `json:"data"`
	} `json:"comments"`
}

// InstagramComments represents a comment on an Instagram Media object.
type InstagramComments struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
}

type InstagramMonitorRequest struct {
	UserID uuid.UUID `json:"user_id"`
	LiveID uuid.UUID `json:"live_id"`
	Action string    `json:"action"`
}
