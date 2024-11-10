package qobuz

import (
	"encoding/json"
	"fmt"
	"github.com/szerookii/goquobuz/qobuz/types"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (s *QobuzClient) Album(id string) (*types.FullAlbum, error) {
	params := url.Values{}
	params.Set("album_id", id)
	params.Set("offset", "0")
	params.Set("limit", "50")
	params.Set("extras", "track_ids,albumsFromSameArtist")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/album/get?%s", APIBaseURL, params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-User-Auth-Token", s.authToken)
	req.Header.Set("X-App-Id", s.app_id)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	album := &types.FullAlbum{}
	err = json.Unmarshal(body, album)
	if err != nil {
		return nil, err
	}

	return album, nil
}
