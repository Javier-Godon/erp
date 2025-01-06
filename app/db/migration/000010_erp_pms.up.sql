CREATE SCHEMA IF NOT EXISTS PMS;
--*****************PMS***********************
CREATE TABLE IF NOT EXISTS
    PMS.reorder_point
(
    reorder_point_id varchar(256) PRIMARY KEY,
    item_id          varchar(256),
    warehouse_id     varchar(256),
    units            integer
);

CREATE TABLE IF NOT EXISTS
    PMS.default_order
(
    default_order_id varchar(256) PRIMARY KEY,
    item_id          varchar(256),
    warehouse_id     varchar(256),
    units            integer
);