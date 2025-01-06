CREATE SCHEMA IF NOT EXISTS SALESFORCE;
--***********SALESFORCE**********************

CREATE TABLE IF NOT EXISTS
    SALESFORCE.promo_in_discount
(
    promo_id varchar(256) PRIMARY KEY,
    discount integer
);

CREATE TABLE IF NOT EXISTS
    SALESFORCE.promo_in_price
(
    promo_id    varchar(256) PRIMARY KEY,
    final_price integer
);

CREATE TABLE IF NOT EXISTS
    SALESFORCE.promo_in_units
(
    promo_id          varchar(256) PRIMARY KEY,
    units_required    integer,
    units_to_take_off integer
);

CREATE TABLE IF NOT EXISTS
    SALESFORCE.categories_in_promo
(
    category_in_promo_id varchar(256) PRIMARY KEY,
    promo_id             varchar(256),
    category_id          varchar(256)
);

CREATE TABLE IF NOT EXISTS
    SALESFORCE.items_in_promo
(
    item_in_promo_id varchar(256) PRIMARY KEY,
    promo_id         varchar(256),
    category_id      varchar(256)
);

CREATE TABLE IF NOT EXISTS
    SALESFORCE.promo
(
    promo_id    varchar(256) PRIMARY KEY,
    term        TIMESTAMP WITHOUT TIME ZONE,
    active      boolean,
    description text
);