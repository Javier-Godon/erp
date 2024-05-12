CREATE SCHEMA IF NOT EXISTS OPERATOR;
--**************OPERATOR*********************
CREATE TABLE IF NOT EXISTS
    OPERATOR.currencies
(
    currency_id varchar(256) PRIMARY KEY,
    iso_4217    varchar(3)
);

CREATE TABLE IF NOT EXISTS
    OPERATOR.item_input_reason
(
    input_reason_id integer PRIMARY KEY,
    input_reason    varchar(120)
);

CREATE TABLE IF NOT EXISTS
    OPERATOR.item_output_reason
(
    output_reason_id integer PRIMARY KEY,
    output_reason    varchar(120)
);

CREATE TABLE IF NOT EXISTS
    OPERATOR.distinctive_features
(
    distinctive_feature_id          varchar(256) PRIMARY KEY,
    distinctive_feature_name        varchar(120),
    distinctive_feature_description varchar(500)
);

CREATE TABLE IF NOT EXISTS
    OPERATOR.units_of_measure
(
    unit_of_measure_id integer PRIMARY KEY,
    unit_of_measure    varchar(120)
);

CREATE TABLE IF NOT EXISTS
    OPERATOR.warehouse_item_category_value
(
    valuable_id integer PRIMARY KEY,
    level       varchar(10)
);

insert into OPERATOR.item_input_reason (input_reason_id, input_reason)
values (1, 'external purchase');
insert into OPERATOR.item_input_reason (input_reason_id, input_reason)
values (2, 'internal purchase');
insert into OPERATOR.item_input_reason (input_reason_id, input_reason)
values (3, 'refund');
insert into OPERATOR.item_input_reason (input_reason_id, input_reason)
values (4, 'count');
insert into OPERATOR.item_input_reason (input_reason_id, input_reason)
values (5, 'unknown');

insert into OPERATOR.item_output_reason (output_reason_id, output_reason)
values (1, 'sale');
insert into OPERATOR.item_output_reason (output_reason_id, output_reason)
values (2, 'waste');
insert into OPERATOR.item_output_reason (output_reason_id, output_reason)
values (3, 'robbery');
insert into OPERATOR.item_output_reason (output_reason_id, output_reason)
values (4, 'expiration');
insert into OPERATOR.item_output_reason (output_reason_id, output_reason)
values (5, 'count');
insert into OPERATOR.item_output_reason (output_reason_id, output_reason)
values (6, 'unknown');

insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (1, 'KG');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (2, 'TON');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (3, 'POUND');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (4, 'STONE');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (5, 'METER');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (6, 'FEET');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (7, 'CENTIMETER');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (8, 'INCH');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (9, 'LITER');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (10, 'GALLON');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (11, 'OUNCE');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (12, 'PHYSICAL UNIT');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (13, 'KWH');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (14, 'PACKAGE');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (15, 'PALE');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (16, 'CONTAINER');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (17, 'BOX');
insert into OPERATOR.units_of_measure (unit_of_measure_id, unit_of_measure)
values (18, 'OTHER');

insert into OPERATOR.warehouse_item_category_value (valuable_id, level)
values (1, 'A');
insert into OPERATOR.warehouse_item_category_value (valuable_id, level)
values (2, 'B');
insert into OPERATOR.warehouse_item_category_value (valuable_id, level)
values (3, 'C');

insert into OPERATOR.currencies(currency_id, iso_4217)
values ('1', 'EUR');
insert into OPERATOR.currencies(currency_id, iso_4217)
values ('2', 'USD');
insert into OPERATOR.currencies(currency_id, iso_4217)
values ('3', 'CLF');
insert into OPERATOR.currencies(currency_id, iso_4217)
values ('4', 'BRL');
insert into OPERATOR.currencies(currency_id, iso_4217)
values ('5', 'MXV');
