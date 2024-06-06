package outerapi

import (
	"github.com/go-resty/resty/v2"
)

type Outerapi struct {
	client *resty.Client
}

func New() Outerapi {
	C := resty.New()
	return Outerapi{client: C}
}

func (cl *Outerapi) GetMovieDataIMDB(IMDBid string) ([]byte, bool) {
	resp, err := cl.client.R().Get("https://www.omdbapi.com/?i=" + IMDBid + "&apikey=e55bdc80")
	if err != nil {
		return nil, false
	}
	return resp.Body(), true
}
