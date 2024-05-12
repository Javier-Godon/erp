CREATE SCHEMA IF NOT EXISTS COLLECTIONS;
--************COLLECTIONS********************

CREATE TABLE IF NOT EXISTS
    COLLECTIONS.collections
(
    collection_id       varchar(256) PRIMARY KEY,
    order_id            varchar(256),
    delivery_id         varchar(256),
    amount              integer,
    currency_id         varchar(256),
    due_date            date,
    collection_date     TIMESTAMP WITHOUT TIME ZONE,
    amount_collected    integer,
    amount_renegotiated integer
);

CREATE TABLE IF NOT EXISTS
    COLLECTIONS.collections_renegotiation
(
    collection_renegotiated_id varchar(256) PRIMARY KEY,
    collection_origin_id       varchar(256),
    renegotiation_date         TIMESTAMP WITHOUT TIME ZONE

);