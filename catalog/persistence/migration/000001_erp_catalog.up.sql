CREATE SCHEMA IF NOT EXISTS CATALOG;
--**************CATALOG****************
CREATE TABLE IF NOT EXISTS
    CATALOG.item
(
    item_id                  UUID PRIMARY KEY,
    item_title               text,
    manufacturer_id          UUID,
    manufacturer_name_line   text,
    manufacturer_original_id text
);

CREATE TABLE IF NOT EXISTS
    CATALOG.category
(
    category_id          UUID PRIMARY KEY,
    category_name        text,
    category_description text
);

CREATE TABLE IF NOT EXISTS
    CATALOG.category_link
(
    category_link_id   UUID PRIMARY KEY,
    main_category_id   UUID,
    linked_category_id UUID
);

CREATE TABLE IF NOT EXISTS
    CATALOG.item_categorized
(
    item_categorized_id UUID PRIMARY KEY,
    item_id             UUID,
    category_leaf_id    UUID
);

CREATE TABLE IF NOT EXISTS
    CATALOG.item_retail_price
(
    item_id      UUID PRIMARY KEY,
    retail_price integer,
    currency_id  UUID
);

CREATE TABLE IF NOT EXISTS
    CATALOG.item_image
(
    item_image_id UUID PRIMARY KEY,
    item_id       UUID,
    image         bytea
);

CREATE TABLE IF NOT EXISTS
    CATALOG.item_description
(
    item_id     UUID PRIMARY KEY,
    description text
);