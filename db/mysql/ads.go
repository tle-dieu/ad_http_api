package mysql

import (
	"log"

	model "github.com/tle-dieu/gql_test/pkg/protobuf"
)

func (cli *Client) CreateAd(ad *model.Ad) error {
	stmt, err := cli.db.Prepare("INSERT INTO Ads(ref,brand,model,price,bluetooth,gps) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ad.Ref, ad.Brand, ad.Model, ad.Price, ad.Options.Bluetooth, ad.Options.Gps)
	// should not fatal if duplicate ref
	if err != nil {
		return err
	}
	log.Println("Row inserted!")
	return nil
}

func (cli *Client) UpdateAd(ad *model.Ad) error {
	stmt, err := cli.db.Prepare("UPDATE Ads SET brand = ?, model = ?, price = ?, bluetooth = ?, gps = ? WHERE ref = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ad.Brand, ad.Model, ad.Price, ad.Options.Bluetooth, ad.Options.Gps, ad.Ref)
	if err != nil {
		return err
	}
	log.Println("Row updated!")
	return nil
}

func (cli *Client) DeleteAd(ref string) error {
	stmt, err := cli.db.Prepare("DELETE FROM Ads WHERE ref = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ref)
	if err != nil {
		return err
	}
	log.Println("Row deleted!")
	return nil
}

func (cli *Client) GetAllAds() (*model.Ads, error) {
	stmt, err := cli.db.Prepare("SELECT ref,brand,model,price,bluetooth,gps FROM Ads")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ads := &model.Ads{}
	for rows.Next() {
		var ad model.Ad
		ad.Options = &model.Ad_Options{}
		err := rows.Scan(&ad.Ref, &ad.Brand, &ad.Model, &ad.Price, &ad.Options.Bluetooth, &ad.Options.Gps)
		if err != nil {
			return nil, err
		}
		ads.Ads = append(ads.Ads, &ad)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ads, nil
}
