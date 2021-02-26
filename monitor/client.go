package monitor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

type (
	Commitment string
)

const (
	// Most recent block confirmed by supermajority of the cluster as having reached maximum lockout.
	CommitmentMax Commitment = "max"
	// Most recent block having reached maximum lockout on this node.
	CommitmentRoot Commitment = "root"
	// Most recent block that has been voted on by supermajority of the cluster (optimistic confirmation).
	CommitmentSingleGossiper Commitment = "singleGossip"
	// The node will query its most recent block. Note that the block may not be complete.
	CommitmentRecent Commitment = "recent"
)

func addQueryParameters(req *http.Request, queryParams types.QueryParams) {
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	req.URL.RawQuery = params.Encode()
}

//newHTTPRequest to make a new http request
func newHTTPRequest(ops types.HTTPOptions) (*http.Request, error) {
	// make new request
	payloadBytes, _ := json.Marshal(ops.Body)
	req, err := http.NewRequest(ops.Method, ops.Endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Add any query parameters to the URL.
	if len(ops.QueryParams) != 0 {
		addQueryParameters(req, ops.QueryParams)
	}

	return req, nil
}

func makeResponse(res *http.Response) (*types.PingResp, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &types.PingResp{}, err
	}

	response := &types.PingResp{
		StatusCode: res.StatusCode,
		Body:       body,
	}
	_ = res.Body.Close()
	return response, nil
}

// HitHTTPTarget to hit the target and get response
func HitHTTPTarget(ops types.HTTPOptions) (*types.PingResp, error) {
	req, err := newHTTPRequest(ops)
	if err != nil {
		return nil, err
	}

	httpcli := http.Client{Timeout: time.Duration(5 * time.Second)}
	resp, err := httpcli.Do(req)
	if err != nil {
		return nil, err
	}

	res, err := makeResponse(resp)
	if err != nil {
		return nil, err
	}

	return res, nil
}
