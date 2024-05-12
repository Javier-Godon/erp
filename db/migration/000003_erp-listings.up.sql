CREATE SCHEMA IF NOT EXISTS LISTINGS;

--***********SALES*******************
CREATE TABLE IF NOT EXISTS LISTINGS.sales_order
(
    order_id                              varchar(256) PRIMARY KEY,
    client_id                             varchar(256),
    date_of_order                         TIMESTAMP WITHOUT TIME ZONE,
    handled                               boolean,
--person
    person_id                             varchar(256),
    is_company                            boolean,
    id_card_number                        varchar(20),
    id_card_number_second_part            varchar(20),
    national_insurance_number             varchar(256),
    national_insurance_number_second_part varchar(256),
    tax_registration_number               varchar(256),
    tax_registration_number_second_part   varchar(256),
    name_line                             varchar(256),
    first_name                            varchar(50),
    middle_name                           varchar(50),
    last_name                             varchar(50),
    fathers_name                          varchar(120),
    mothers_name                          varchar(120),
    date_of_birth                         date,
    organisation_name                     varchar(120),
    --address
    address_id                            varchar(256),
    country_id                            varchar(256),
    administrative_area                   varchar(120),
    administrative_area_iso               varchar(10),
    locality                              varchar(120),
    dependent_locality                    varchar(120),
    postal_code                           varchar(20),
    street                                varchar(120),
    street_number                         integer,
    street_type                           varchar(20),
    premise                               varchar(120),
    sub_premise                           varchar(120),
    department                            varchar(120),
    floor                                 varchar(4),
    letter                                varchar(20),
    stairs                                varchar(20),
    decimal_degrees_latitude              varchar(20),
    decimal_degrees_longitude             varchar(20)
);

CREATE TABLE IF NOT EXISTS LISTINGS.proof_of_delivery
(
    proof_of_delivery_id varchar(256) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS LISTINGS.shipment_order
(
    shipment_order_id varchar(256) PRIMARY KEY
);

--***********FUNCTIONS*****************

CREATE OR REPLACE FUNCTION public.search_order_as_you_type(search_string text)
    RETURNS SETOF character varying
    LANGUAGE sql
AS
$function$
select order_id
from LISTINGS.sales_order
where TRANSLATE(LOWER(concat
                      (id_card_number, name_line, first_name, middle_name, last_name, organisation_name, postal_code,
                       locality, street)
                ), 'âãäåāăąÁÂÃÄÅĀĂĄèééêëēĕėęěĒĔĖĘĚìíîïìĩīĭÌÍÎÏÌĨĪĬóôõöōŏőÒÓÔÕÖŌŎŐùúûüũūŭůÙÚÛÜŨŪŬŮ -_/',
                'aaaaaaaaaaaaaaaeeeeeeeeeeeeeeeiiiiiiiiiiiiiiiiooooooooooooooouuuuuuuuuuuuuuuu')
          ilike
      TRANSLATE(LOWER('%' || search_string || '%'),
                'âãäåāăąÁÂÃÄÅĀĂĄèééêëēĕėęěĒĔĖĘĚìíîïìĩīĭÌÍÎÏÌĨĪĬóôõöōŏőÒÓÔÕÖŌŎŐùúûüũūŭůÙÚÛÜŨŪŬŮ -_/',
                'aaaaaaaaaaaaaaaeeeeeeeeeeeeeeeiiiiiiiiiiiiiiiiooooooooooooooouuuuuuuuuuuuuuuu');
$function$
;

-- CREATE FUNCTION search_person_as_you_type(search_string TEXT)
-- RETURNS setof varchar(256) AS $$
--    select person_id
--    from person where
--    TRANSLATE(LOWER(concat
--    (id_card_number , name_line , first_name , middle_name , last_name , organisation_name)
--    ),'âãäåāăąÁÂÃÄÅĀĂĄèééêëēĕėęěĒĔĖĘĚìíîïìĩīĭÌÍÎÏÌĨĪĬóôõöōŏőÒÓÔÕÖŌŎŐùúûüũūŭůÙÚÛÜŨŪŬŮ -_/',
--    'aaaaaaaaaaaaaaaeeeeeeeeeeeeeeeiiiiiiiiiiiiiiiiooooooooooooooouuuuuuuuuuuuuuuu')
--   ilike
--    TRANSLATE(LOWER('%' || search_string || '%'),'âãäåāăąÁÂÃÄÅĀĂĄèééêëēĕėęěĒĔĖĘĚìíîïìĩīĭÌÍÎÏÌĨĪĬóôõöōŏőÒÓÔÕÖŌŎŐùúûüũūŭůÙÚÛÜŨŪŬŮ -_/',
--    'aaaaaaaaaaaaaaaeeeeeeeeeeeeeeeiiiiiiiiiiiiiiiiooooooooooooooouuuuuuuuuuuuuuuu');
-- $$ LANGUAGE SQL;
