package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ms-live-sells/models"
	"net/http"
	"time"
)

//const apiURL = "https://graph.instagram.com/"

type InstagramProvier struct {
	AccessToken string
}

const baseURL = "https://graph.instagram.com/v21.0"

// CheckLive verifica se há uma live ativa
func (ip *InstagramProvier) CheckLive(userID string) (models.InstagramMedia, bool, error) {
	url := fmt.Sprintf("%s/%s/media?fields=id,media_product_type,created_time&access_token=%s", baseURL, userID, ip.AccessToken)
	resp, err := http.Get(url)
	if err != nil {
		return models.InstagramMedia{}, false, err
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	fmt.Printf("%#v", resp.Body)
	var result models.InstagramResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.InstagramMedia{}, false, err
	}

	fmt.Println(result.Data)
	// Verifica se existe uma mídia do tipo "LIVE" criada recentemente
	for _, media := range result.Data {
		if media.MediaProductType == "LIVE" {
			createdTime, _ := time.Parse(time.RFC3339, media.CreatedTime)
			if time.Since(createdTime) < time.Hour { // Live criada na última hora
				return media, true, nil
			}
		}
	}
	return models.InstagramMedia{}, false, nil
}

// GetMedias Get Instagram User Medias
func (ip *InstagramProvier) GetMedias(userID string) ([]models.InstagramMedia, error) {
	url := fmt.Sprintf("%s/%s/media?fields=id,media_product_type,created_time&access_token=%s", baseURL, userID, ip.AccessToken)
	resp, err := http.Get(url)
	if err != nil {
		return []models.InstagramMedia{}, err
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	fmt.Printf("%#v", resp.Body)
	var result models.InstagramResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []models.InstagramMedia{}, err
	}

	return result.Data, nil
}

// GetComments Get Instagram Media comments
func (ip *InstagramProvier) GetComments(mediaID string) (*models.InstagramCommentResponse, error) {

	apiURL := fmt.Sprintf("https://graph.instagram.com/v21.0/media/%s/comments?fields=username,text&access_token=%s", mediaID, ip.AccessToken)

	// Fazer a requisição para a API do Instagram
	resp, err := http.Get(apiURL)
	if err != nil {
		return &models.InstagramCommentResponse{}, err
	}
	defer resp.Body.Close()

	// Ler a resposta JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &models.InstagramCommentResponse{}, err
	}

	// Decodificar os comentários
	var instagramResp models.InstagramCommentResponse
	err = json.Unmarshal(body, &instagramResp)
	if err != nil {
		return &models.InstagramCommentResponse{}, err
	}

	return &instagramResp, nil
}
