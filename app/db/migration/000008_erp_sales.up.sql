CREATE SCHEMA IF NOT EXISTS SALES;
--**************SALES************************

CREATE TABLE IF NOT EXISTS
    SALES.sales_order
(
    order_id      varchar(256) PRIMARY KEY,
    client_id     varchar(256),
    date_of_order TIMESTAMP WITHOUT TIME ZONE,
    person_id     varchar(256),
    handled       boolean
);

CREATE TABLE IF NOT EXISTS
    SALES.order_line
(
    order_line_id        varchar(256) PRIMARY KEY,
    order_id             varchar(256),
    item_id              varchar(256),
    unique_identifier_id varchar(256),
    warehouse_id         varchar(250),
    retail_price         integer,
    currency_id          varchar(256),
    units                integer,
    units_of_measure     text
);