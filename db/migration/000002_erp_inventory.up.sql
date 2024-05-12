CREATE SCHEMA IF NOT EXISTS INVENTORY;
--**************INVENTORY**************

CREATE TABLE IF NOT EXISTS INVENTORY.global_stock
(
    warehouse_stock_id   varchar(256) PRIMARY KEY,
    item_id              varchar(256),
    unique_identifier_id varchar(256),
    warehouse_id         varchar(250),
    stock                integer,
    CONSTRAINT item_by_warehouse UNIQUE (warehouse_id, item_id)
);

CREATE TABLE IF NOT EXISTS INVENTORY.item_distinctive_feature
(
    unique_identifier_id      varchar(256) PRIMARY KEY,
    distinctive_feature_id    varchar(256),
    distinctive_feature_value varchar(256),
    notes                     text
);

CREATE TABLE IF NOT EXISTS INVENTORY.purchase_order
(
    order_id           varchar(256) PRIMARY KEY,
    date_of_order      TIMESTAMP WITHOUT TIME ZONE,
    cancellation_date  TIMESTAMP WITHOUT TIME ZONE,
    cancellation_notes text
);

CREATE TABLE IF NOT EXISTS INVENTORY.order_line
(
    order_line_id        varchar(256) PRIMARY KEY,
    order_id             varchar(256),
    item_id              varchar(256),
    unique_identifier_id varchar(256),
    warehouse_id         varchar(250),
    retail_price         integer,
    currency_id          varchar(256),
    units                integer
);

CREATE TABLE IF NOT EXISTS INVENTORY.child_order
(
    child_order_id    varchar(256) PRIMARY KEY,
    order_id          varchar(256),
    warehouse_id      varchar(250),
    date_of_order     TIMESTAMP WITHOUT TIME ZONE,
    date_of_shipment  TIMESTAMP WITHOUT TIME ZONE,
    cancellation_date TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS INVENTORY.child_order_line
(
    child_order_line_id  varchar(256),
    order_id             varchar(256),
    item_id              varchar(256),
    unique_identifier_id varchar(256),
    retail_price         integer,
    currency_id          varchar(256),
    units                integer
);
