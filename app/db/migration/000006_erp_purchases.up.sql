CREATE SCHEMA IF NOT EXISTS PURCHASES;
--*********PURCHASES*******************

CREATE TABLE IF NOT EXISTS
    purchases.purchase
(
    purchase_id    varchar(256) PRIMARY KEY,
    supplier_id    varchar(256),
    dateOfPurchase TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS
    purchases.purchase_line
(
    purchase_line_id    varchar(256) PRIMARY KEY,
    purchase_id         varchar(256),
    item_id             varchar(256),
    base_purchase_price integer,
    currency_id         varchar(256),
    units               integer
);

CREATE TABLE IF NOT EXISTS
    purchases.manufacturer
(
    manufacturer_id varchar(256) PRIMARY KEY,
    person_id       varchar(256)
);