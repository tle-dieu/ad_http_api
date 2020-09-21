package mysql

import (
	"github.com/tle-dieu/ad_http_api/domain/model"
)

func (cli *Client) CreateAd(ad *model.Ad) error {
	stmt, err := cli.db.Prepare("INSERT INTO Ads(brand,model,price,bluetooth,gps) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ad.Brand, ad.Model, ad.Price, ad.Options.Bluetooth, ad.Options.Gps)
	if err != nil {
		return err
	}
	return nil
}
