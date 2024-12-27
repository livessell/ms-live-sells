package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ms-live-sells/models"
	"net/http"
	"net/url"
	"sync"
)

// const apiURL = "https://graph.instagram.com/"

type InstagramProvider struct {
	LongLivedToken string
	AccessToken    string
	mu             sync.Mutex // to avoid run conditions
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

const baseURL = "https://graph.instagram.com/v21.0"

// GenerateLongLivedAccessToken troca o token de curta duração por um token de longa duração.
func (ip *InstagramProvider) GenerateLongLivedAccessToken() error {
	apiURL := "https://graph.instagram.com/refresh_access_token"
	params := url.Values{}
	params.Set("grant_type", "ig_refresh_token")
	params.Set("access_token", ip.LongLivedToken)

	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return fmt.Errorf("failed to request access token: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var result AccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	ip.mu.Lock()
	defer ip.mu.Unlock()
	ip.AccessToken = result.AccessToken
	return nil
}

// GetUserByUsername busca os dados do usuário pelo username
func (p *InstagramProvider) GetUserByUsername(username string) (*models.UserInfo, error) {
	// Endpoint público do Instagram para buscar usuários
	url := fmt.Sprintf("https://www.instagram.com/web/search/topsearch/?query=%s", username)
	fmt.Printf("Imprimindo a URL de GET by USERNAME: %s", url)
	// Fazendo a requisição HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	// Estrutura para decodificar a resposta completa do Instagram
	var result struct {
		Users []struct {
			User struct {
				Username string `json:"username"`
				FullName string `json:"full_name"`
				FbidV2   string `json:"fbid_v2"`
			} `json:"user"`
		} `json:"users"`
	}

	// Decodificar a resposta JSON
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Verifica se encontrou o usuário
	if len(result.Users) == 0 {
		return nil, fmt.Errorf("user not found for username: %s", username)
	}

	// Retorna os dados do primeiro usuário encontrado
	user := result.Users[0].User
	return &models.UserInfo{
		Username: user.Username,
		FullName: user.FullName,
		FbidV2:   user.FbidV2,
	}, nil
}

// GetLiveMedia retrieves the live video IG Media for an IG User.
func (ip *InstagramProvider) GetLiveMedia(socialNetworkID string) ([]models.InstagramMedias, error) {
	apiURL := fmt.Sprintf("%s/%s/live_media?fields=id,media_type,media_product_type,owner,username,comments&access_token=%s", baseURL, socialNetworkID, ip.AccessToken)
	// Realiza a requisição para o endpoint
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch live media: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o código de status da resposta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decodifica o corpo da resposta JSON
	var response struct {
		Data   []models.InstagramMedias `json:"data"`
		Paging struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"paging"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return response.Data, nil
}

// GetMedias Get Instagram User Medias
func (ip *InstagramProvider) GetMedias(userID string) ([]models.InstagramMedia, error) {
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
func (ip *InstagramProvider) GetComments(mediaID string) (*models.InstagramCommentResponse, error) {

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
