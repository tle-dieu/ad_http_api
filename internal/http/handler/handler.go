package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tle-dieu/ad_http_api/domain/model"
	"github.com/tle-dieu/ad_http_api/internal/db/mysql"
)

type createAdResponse struct {
	Ref string `json:"ref"`
}

func CreateAd(db mysql.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ad := &model.Ad{}
		if r.Body == nil || r.Body == http.NoBody {
			http.Error(w, "Please send a body", http.StatusBadRequest)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(ad); err != nil {
			http.Error(w, "Cannot decode data: "+err.Error(), http.StatusBadRequest)
			return
		}
		ref, err := db.CreateAd(ad)
		if err != nil {
			http.Error(w, "Error while creating ad: "+err.Error(), http.StatusInternalServerError)
			return
		}
		respBody, err := json.Marshal(createAdResponse{Ref: ref})
		if err != nil {
			http.Error(w, "Error while transforming response: "+err.Error(), http.StatusInternalServerError)
		}
		if _, err = w.Write(respBody); err != nil {
			http.Error(w, "Error while writing response: "+err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusAccepted)
	}
}
