package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tle-dieu/gql_test/internal/db/mysql"
	"github.com/tle-dieu/gql_test/pkg/protobuf"
	"google.golang.org/protobuf/proto"
)

func getProtobufRequest(req *http.Request, message proto.Message) error {
	if req.Body == nil || req.Body == http.NoBody {
		return errors.New("Please send a request body")
	}
	if req.Header.Get("Content-Type") != "application/protobuf" {
		return errors.New("Content-Type must be application/protobuf")
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return errors.New("Unable to read message from request : " + err.Error())
	}
	if err := proto.Unmarshal(body, message); err != nil {
		return err
	}
	return nil
}

func GetAds(db mysql.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "application/protobuf") {
			panic("bad accept encoding")
		}
		ads, err := db.GetAllAds()
		if err != nil {
			panic(err) // @FIXME
		}
		response, err := proto.Marshal(ads)
		if err != nil {
			panic(err) // @FIXME
		}
		_, err = w.Write(response)
		if err != nil {
			log.Fatalf("Unable to write data into HTTP response: %v", err)
		}
	}
}

func CreateAd(db mysql.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ad := &protobuf.Ad{}
		err := getProtobufRequest(r, ad)
		if err != nil {
			panic(err)
		}
		err = db.CreateAd(ad)
		if err != nil {
			panic(err) // @FIXME
		}
	}
}

func UpdateAd(db mysql.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ad := &protobuf.Ad{}
		err := getProtobufRequest(r, ad)
		if err != nil {
			panic(err)
		}
		err = db.UpdateAd(ad)
		if err != nil {
			panic(err) // @FIXME
		}
	}
}

func DeleteAd(db mysql.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ref := mux.Vars(r)["ad_ref"]
		err := db.DeleteAd(ref)
		if err != nil {
			panic(err)
		}
	}
}
