CREATE SCHEMA IF NOT EXISTS RETURNS;
--***************RETURNS*********************
CREATE TABLE IF NOT EXISTS
    RETURNS.return
(
    return_id      varchar(256) PRIMARY KEY,
    invoice_number varchar(500),
    customer_id    varchar(256),
    notes          text
);

CREATE TABLE IF NOT EXISTS
    RETURNS.return_lines
(
    return_line_id           varchar(256) PRIMARY KEY,
    return_id                varchar(256),
    manufacturer_id          varchar(256),
    manufacturer_original_id varchar(512),
    units                    integer
);