package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tle-dieu/ad_http_api/domain/model"
	"github.com/tle-dieu/ad_http_api/internal/db/mysql"
)

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
		if err := db.CreateAd(ad); err != nil {
			http.Error(w, "Error while creating ad: "+err.Error(), http.StatusBadRequest)
			return
		}
	}
}
