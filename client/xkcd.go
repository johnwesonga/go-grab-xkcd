package client

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/johnwesonga/go-grab-xkcd/model"
)

type ComicNumber int

const (
	// BaseURL of xkcd
	BaseURL string = "https://xkcd.com"
	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout time.Duration = 30 * time.Second
	// LatestComic is the latest comic number according to the xkcd API
	LatestComic ComicNumber = 0
)

// XKCDClient is the client for XKCD
type XKCDClient struct {
	client  *http.Client
	baseURL string
}

// NewXKCDClient creates a new XKCDClient
func NewXKCDClient() *XKCDClient {
	return &XKCDClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
	}
}

// SetTimeout overrides the default ClientTimeout
func (hc *XKCDClient) SetTimeout(d time.Duration) {
	hc.client.Timeout = d
}

func (hc *XKCDClient) buildURL(n ComicNumber) string {
	var finalURL string
	finalURL = fmt.Sprintf("%s/%d/info.0.json", hc.baseURL, n)
	return finalURL
}

func randomComicNumber() ComicNumber {
	min, max := 1, 999
	rand.Seed(time.Now().UTC().UnixNano())
	return ComicNumber(min + rand.Intn(max-min))
}

// Fetch retrieves the comic as per provided comic number
func (hc *XKCDClient) Fetch() (model.Comic, error) {
	n := randomComicNumber()
	resp, err := hc.client.Get(hc.buildURL(n))
	if err != nil {
		return model.Comic{}, err
	}
	defer resp.Body.Close()

	var comicResp model.ComicResponse
	if err := json.NewDecoder(resp.Body).Decode(&comicResp); err != nil {
		return model.Comic{}, err
	}

	return comicResp.Comic(), nil
}
