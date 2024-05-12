CREATE SCHEMA IF NOT EXISTS DELIVERY;
--****************DELIVERY*******************

CREATE TABLE IF NOT EXISTS
    DELIVERY.delivery
(
    delivery_id                   varchar(256) PRIMARY KEY,
    order_id                      varchar(256),
    departure_date                TIMESTAMP WITHOUT TIME ZONE,
    departure_person_in_charge_id varchar(256),
    delivery_date                 TIMESTAMP WITHOUT TIME ZONE,
    delivery_person_in_charge_id  varchar(256),
    receiver_person_id            varchar(256),
    receiver_not_registered       text,
    receiver_signature            bytea,
    notes                         text,
    country_id                    varchar(256),
    administrative_area           varchar(120),
    administrative_area_iso       varchar(10),
    locality                      varchar(120),
    dependent_locality            varchar(120),
    postal_code                   varchar(20),
    street                        varchar(120),
    street_number                 integer,
    street_type                   varchar(20),
    premise                       varchar(120),
    sub_premise                   varchar(120),
    department                    varchar(120),
    floor                         varchar(4),
    letter                        varchar(20),
    stairs                        varchar(20),
    decimal_degrees_latitude      varchar(20),
    decimal_degrees_longitude     varchar(20)
);

CREATE TABLE IF NOT EXISTS
    DELIVERY.delivery_line
(
    delivery_line_id  varchar(256) PRIMARY KEY,
    delivery_id       varchar(256),
    item_id           varchar(256),
    unique_identifier text,
    delivered_units   integer,
    units_of_measure  text,
    line_closed       boolean,
    retail_price      integer,
    currency          varchar(3)
);

CREATE TABLE IF NOT EXISTS
    DELIVERY.delivery_directions
(
    delivery_id        varchar(256) PRIMARY KEY,
    receiver_person_id varchar(256),
    directions         text
);