package ads

import (
	"net/http"

	"github.com/gemalto/requester"
	"github.com/tle-dieu/ad_http_api/domain/model"
)

type Client struct {
	requester *requester.Requester
}

type createAdResponse struct {
	ref string `json:"ref"`
}

func NewClient(apiUrl string, httpClient *http.Client) *Client {
	return &Client{
		requester: requester.MustNew(
			requester.JSON(true),
			requester.URL(apiUrl),
			requester.WithDoer(httpClient),
			requester.ExpectCode(http.StatusAccepted),
		),
	}
}

func (cli *Client) CreateAd(ad *model.Ad) (createAdReponse, error) {
	requesterOptions := []requester.Option{
		requester.Post("/createAd"),
		requester.Body(ad),
	}
	var resp createAdReponse
	_, _, err := cli.requester.Receive(
		&resp,
		requesterOptions...,
	)
	return resp, err
}
