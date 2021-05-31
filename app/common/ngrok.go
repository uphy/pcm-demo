package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type (
	NgrokClient struct {
		endpoint string
	}
	NgrokTunnelsResponse struct {
		Tunnels []NgrokTunnel `json:"tunnels"`
	}
	NgrokTunnel struct {
		Name      string `json:"name"`
		PublicURL string `json:"public_url"`
		Protocol  string `json:"proto"`
	}
	NgrokEndpoints struct {
		Media      string
		Advertiser string
	}
)

func WaitForNgrok() (*NgrokEndpoints, error) {
	log.Println("Waiting for ngrok endpoints...")
	media, err := waitForNgrok("ngrok-media:4040")
	if err != nil {
		return nil, err
	}
	advertiser, err := waitForNgrok("ngrok-advertiser:4040")
	if err != nil {
		return nil, err
	}
	log.Printf("ngrok ready: media=%s, advertiser=%s", media, advertiser)
	return &NgrokEndpoints{media, advertiser}, nil
}

func waitForNgrok(endpoint string) (string, error) {
	client := NewNgrokClient(endpoint)
	for i := 0; i < 60; i++ {
		resp, err := client.Tunnels()
		if err != nil {
			log.Printf("Waiting ngrok endpoint: %s", endpoint)
			time.Sleep(time.Second)
			continue
		}
		tunnel := resp.FindTunnelByName("app")
		if tunnel == nil {
			time.Sleep(time.Second)
			continue
		}
		return tunnel.PublicURL, nil
	}
	return "", errors.New("ngrok not started")
}

func NewNgrokClient(endpoint string) *NgrokClient {
	return &NgrokClient{endpoint}
}

func (n *NgrokClient) Tunnels() (*NgrokTunnelsResponse, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/tunnels", n.endpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to get ngrok tunnels: %w", err)
	}
	defer resp.Body.Close()
	var decoded NgrokTunnelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return nil, fmt.Errorf("unable to parse ngrok tunnel response: %w", err)
	}
	return &decoded, nil
}

func (r *NgrokTunnelsResponse) FindTunnelByName(name string) *NgrokTunnel {
	for _, t := range r.Tunnels {
		if t.Name == name {
			return &t
		}
	}
	return nil
}
