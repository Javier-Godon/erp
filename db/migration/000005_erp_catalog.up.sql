CREATE SCHEMA IF NOT EXISTS CATALOG;
--**************CATALOG****************
CREATE TABLE IF NOT EXISTS
    CATALOG.item
(
    item_id                  varchar(256) PRIMARY KEY,
    item_title               varchar(256),
    manufacturer_id          varchar(256),
    manufacturer_name_line   varchar(512),
    manufacturer_original_id varchar(512)
);

CREATE TABLE IF NOT EXISTS
    CATALOG.category
(
    category_id          varchar(256) PRIMARY KEY,
    category_name        varchar(120),
    category_description text
);

CREATE TABLE IF NOT EXISTS
    CATALOG.category_link
(
    category_link_id   varchar(256) PRIMARY KEY,
    main_category_id   varchar(256),
    linked_category_id varchar(256)
);

CREATE TABLE IF NOT EXISTS
    CATALOG.item_categorized
(
    item_categorized_id varchar(256) PRIMARY KEY,
    item_id             varchar(256),
    category_leaf_id    varchar(256)
);

CREATE TABLE IF NOT EXISTS
    CATALOG.item_retail_price
(
    item_id      varchar(256) PRIMARY KEY,
    retail_price integer,
    currency_id  varchar(256)
);

CREATE TABLE IF NOT EXISTS
    CATALOG.item_image
(
    item_image_id varchar(256) PRIMARY KEY,
    item_id       varchar(256),
    image         bytea
);

CREATE TABLE IF NOT EXISTS
    CATALOG.item_description
(
    item_id     varchar(256) PRIMARY KEY,
    description text
);