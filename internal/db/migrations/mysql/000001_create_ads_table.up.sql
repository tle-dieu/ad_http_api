CREATE TABLE IF NOT EXISTS Ads(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    brand VARCHAR (127),
    model VARCHAR (127),
    price INT,
    bluetooth TINYINT(1),
    gps TINYINT(1)
)