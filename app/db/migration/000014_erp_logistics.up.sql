CREATE SCHEMA IF NOT EXISTS LOGISTICS;
--************LOGISTICS**********************

CREATE TABLE IF NOT EXISTS
    LOGISTICS.warehouse
(
    warehouse_id varchar(256) PRIMARY KEY,
    person_id    varchar(256)
);

CREATE TABLE IF NOT EXISTS
    LOGISTICS.item_physical_qualities
(
    item_id  varchar(256) PRIMARY KEY,
    volume   varchar(120),
    height   varchar(120),
    width    varchar(120),
    depth    varchar(120),
    diameter varchar(120),
    weight   varchar(120),
    notes    text
);

CREATE TABLE IF NOT EXISTS
    LOGISTICS.item_warehouse_categorization
(
    item_warehouse_categorization_id varchar(256) PRIMARY KEY,
    item_id                          varchar(256),
    warehouse_id                     varchar(256),
    traceable                        boolean,
    valuable_id                      integer
);