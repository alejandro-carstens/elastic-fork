package elasticfork

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/alejandro-carstens/elasticfork/uritemplates"
)

type IndicesRecoveryService struct {
	client     *Client
	indices    []string
	human      *bool
	detailed   *bool
	activeOnly *bool
}

func NewIndicesRecoveryService(client *Client) *IndicesRecoveryService {
	return &IndicesRecoveryService{
		client: client,
	}
}

func (irs *IndicesRecoveryService) Indices(indices ...string) *IndicesRecoveryService {
	irs.indices = indices
	return irs
}

func (irs *IndicesRecoveryService) Human(human bool) *IndicesRecoveryService {
	irs.human = &human
	return irs
}

func (irs *IndicesRecoveryService) buildURL() (string, url.Values, error) {
	var err error
	var path string

	if len(irs.indices) > 0 {
		path, err = uritemplates.Expand("/{indices}/_recovery", map[string]string{
			"indices": strings.Join(irs.indices, ","),
		})
	} else {
		path, err = uritemplates.Expand("/_recovery", map[string]string{})
	}

	if err != nil {
		return "", url.Values{}, err
	}

	params := url.Values{}

	if irs.human != nil {
		params.Set("human", fmt.Sprintf("%v", *irs.human))
	}

	if irs.detailed != nil {
		params.Set("detailed", fmt.Sprintf("%v", *irs.detailed))
	}

	if irs.activeOnly != nil {
		params.Set("detailed", fmt.Sprintf("%v", *irs.activeOnly))
	}

	return path, params, nil
}

func (irs *IndicesRecoveryService) Do(ctx context.Context) (*Response, error) {
	path, params, err := irs.buildURL()

	if err != nil {
		return nil, err
	}

	res, err := irs.client.PerformRequest(ctx, PerformRequestOptions{
		Method: "GET",
		Path:   path,
		Params: params,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

type IndexRecoveryResponse struct {
	Id                string `json:"id"`
	Type              string `json:"type"`
	Stage             string `json:"stage"`
	Primary           bool   `json:"primary"`
	StartTime         string `json:"start_time"`
	StartTimeInMillis string `json:"start_time_in_millis"`
	StopTime          string `json:"stop_time"`
	StopTimeInMillis  string `json:"stop_time_in_millis"`
	TotalTime         string `json:"total_time"`
	TotalTimeInMillis int64  `json:"total_time_in_millis"`
	Source            struct {
		Id               string `json:"id"`
		Host             string `json:"host"`
		TransportAddress string `json:"transport_address"`
		Ip               string `json:"ip"`
		Name             string `json:"name"`
	} `json:"source"`
}
