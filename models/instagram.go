package models

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
