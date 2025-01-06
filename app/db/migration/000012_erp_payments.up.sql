CREATE SCHEMA IF NOT EXISTS PAYMENTS;
--****************PAYMENTS*******************

CREATE TABLE IF NOT EXISTS
    PAYMENTS.purchases_payments
(
    payment_id          varchar(256) PRIMARY KEY,
    purchase_id         varchar(256),
    receipt_id          varchar(256),
    amount              integer,
    currency_id         varchar(256),
    due_date            date,
    payment_date        TIMESTAMP WITHOUT TIME ZONE,
    amount_paid         integer,
    amount_renegotiated integer
);

CREATE TABLE IF NOT EXISTS
    PAYMENTS.payments_renegotiation
(
    payment_renegotiated_id varchar(256) PRIMARY KEY,
    payment_origin_id       varchar(256),
    renegotiation_date      TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS
    PAYMENTS.structural_payments_definition
(
    structural_payment_id varchar(256) PRIMARY KEY,
    frequency_in_days     integer,
    first_due_date        date,
    last_due_date         date,
    fixed_due_dates       boolean,
    description           text
);

CREATE TABLE IF NOT EXISTS
    PAYMENTS.structural_payments_fixed_due_date
(
    payment_id            varchar(256) PRIMARY KEY,
    structural_payment_id varchar(256),
    due_date              date
);

CREATE TABLE IF NOT EXISTS
    PAYMENTS.structural_payments_due_date
(
    payment_id            varchar(256) PRIMARY KEY,
    structural_payment_id varchar(256),
    next_due_date         date
);

CREATE TABLE IF NOT EXISTS
    PAYMENTS.structural_payments
(
    payment_id            varchar(256) PRIMARY KEY,
    structural_payment_id varchar(256),
    contract_id           varchar(512),
    amount                integer,
    currency_id           varchar(256),
    short_description     varchar(60)
);