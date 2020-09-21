package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/tle-dieu/ad_http_api/domain/model"
	"github.com/tle-dieu/ad_http_api/internal/db/mysql"
)

// func getProtobufRequest(req *http.Request, message proto.Message) error {
// 	if req.Body == nil || req.Body == http.NoBody {
// 		return errors.New("Please send a request body")
// 	}
// 	if req.Header.Get("Content-Type") != "application/protobuf" {
// 		return errors.New("Content-Type must be application/protobuf")
// 	}
// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		return errors.New("Unable to read message from request : " + err.Error())
// 	}
// 	if err := proto.Unmarshal(body, message); err != nil {
// 		return err
// 	}
// 	return nil
// }

func CreateAd(db mysql.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ad := &model.Ad{}
		// r.Body = nil
		if err := json.NewDecoder(r.Body).Decode(ad); err != nil {
			panic(err)
		}
		if err := db.CreateAd(ad); err != nil {
			panic(err) // @FIXME
		}
	}
}
