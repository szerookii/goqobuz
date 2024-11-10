package qobuz

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type TrackURLResponse struct {
	TrackId      int    `json:"track_id"`
	Duration     int    `json:"duration"`
	Url          string `json:"url"`
	FormatId     int    `json:"format_id"`
	MimeType     string `json:"mime_type"`
	Restrictions []struct {
		Code string `json:"code"`
	} `json:"restrictions"`
	SamplingRate float64 `json:"sampling_rate"`
	BitDepth     int     `json:"bit_depth"`
	Blob         string  `json:"blob"`
}

func (s *QobuzClient) DownloadFileLink(trackID string, quality int) (*TrackURLResponse, error) {
	unixTS := strconv.FormatInt(time.Now().Unix(), 10)
	rSig := fmt.Sprintf("trackgetFileUrlformat_id%dintentstreamtrack_id%s%s%s", quality, trackID, unixTS, s.secret)

	hasher := md5.New()
	hasher.Write([]byte(rSig))
	rSigHashed := hex.EncodeToString(hasher.Sum(nil))

	params := url.Values{}
	params.Set("request_ts", unixTS)
	params.Set("request_sig", rSigHashed)
	params.Set("track_id", trackID)
	params.Set("format_id", strconv.Itoa(quality))
	params.Set("intent", "stream")

	req, err := http.NewRequest("GET", APIBaseURL+"/track/getFileUrl?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-App-Id", s.app_id)
	req.Header.Set("X-User-Auth-Token", s.authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var trackURLResponse *TrackURLResponse
	err = json.NewDecoder(resp.Body).Decode(&trackURLResponse)
	if err != nil {
		return nil, err
	}

	return trackURLResponse, nil
}
