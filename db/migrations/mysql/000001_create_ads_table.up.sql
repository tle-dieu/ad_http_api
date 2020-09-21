CREATE TABLE IF NOT EXISTS Ads(
    ref VARCHAR (127) UNIQUE,
    brand VARCHAR (127),
    model VARCHAR (127),
    price INT,
    bluetooth TINYINT(1),
    gps TINYINT(1),
    PRIMARY KEY (ref)
)