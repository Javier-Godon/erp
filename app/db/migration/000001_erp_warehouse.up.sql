CREATE SCHEMA IF NOT EXISTS WAREHOUSE;
--*******WAREHOUSE**********************

CREATE TABLE IF NOT EXISTS warehouse.receipt(
  receipt_id varchar(256) PRIMARY KEY,
  warehouse_id varchar(256),
  date_of_dispatch TIMESTAMP WITHOUT TIME ZONE,
  reception_date TIMESTAMP WITHOUT TIME ZONE,
  origin_id varchar(256),
  delivery_note_number varchar(500),
  input_reason_id integer,
  person_in_charge_id varchar(256)
);

CREATE TABLE IF NOT EXISTS warehouse.receipt_lines(
  receipt_line_id varchar(256) PRIMARY KEY,
  receipt_id varchar(256),
  item_id varchar(256),
  base_purchase_price integer,
  currency_id varchar(256),
  units integer
);

CREATE TABLE IF NOT EXISTS warehouse.entry_check (
  unique_identifier_id varchar(256) PRIMARY KEY,
  item_id varchar(256),
  receipt_id varchar(256),
  unique_identifier varchar(500),
  digital_identifier bytea,
  person_in_charge_id varchar(256)
);

CREATE TABLE IF NOT EXISTS warehouse.item_distinctive_features (
  unique_identifier_id varchar(256) PRIMARY KEY,
  distinctive_feature_id varchar(256),
  distinctive_feature_value varchar(256),
  notes text
);

--entry_check >caracterisicas >guardarlas y pasarlas al inventario

CREATE TABLE IF NOT EXISTS warehouse.shipment_order(
  shipment_order_id varchar(256) PRIMARY KEY,
  order_id varchar(256),
  date_of_order  TIMESTAMP WITHOUT TIME ZONE,
  client_data text,
  warehouse_id varchar(256),
  reception_date TIMESTAMP WITHOUT TIME ZONE,
  last_delivery_date TIMESTAMP WITHOUT TIME ZONE,
  cancellation_date TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS warehouse.shipment_order_line (
  shipment_order_line_id varchar(256) PRIMARY KEY,
  shipment_order_id varchar(256),
  item_id varchar(256),
  unique_identifier varchar(500),
  requested_units integer,
  shipped_units integer,
  units_of_measure text,
  line_closed boolean,
  discrepancy_reason text,
  notes text,
  person_id varchar(256)
);

CREATE TABLE IF NOT EXISTS warehouse.shipment(
  shipment_id varchar(256) PRIMARY KEY,
  warehouse_id varchar(256),
  delivery_order_id varchar(256),
  date_of_dispatch TIMESTAMP WITHOUT TIME ZONE,
  output_reason_id integer,
  person_in_charge_id varchar(256)
);

CREATE TABLE IF NOT EXISTS warehouse.shipment_line(
  shipment_line_id varchar(256) PRIMARY KEY,
  shipment_order_line_id varchar(256),
  shipment_id varchar(256),
  item_id varchar(256),
  retail_price integer,
  currency_id varchar(256),
  units integer
);

CREATE TABLE IF NOT EXISTS warehouse.exit_check (
  unique_identifier_id varchar(256) PRIMARY KEY,
  item_id varchar(256),
  shipment_id varchar(256),
  unique_identifier varchar(500),
  digital_identifier bytea,
  person_in_charge_id varchar(256)
);

CREATE TABLE IF NOT EXISTS warehouse.item_lambda(
  item_lambda_id varchar(256) PRIMARY KEY,
  flow_id varchar(256),
  item_id varchar(256),
  manufacturer_id varchar(256),
  manufacturer_original_id varchar(256),
  unique_identifier_id varchar(256),
  lambda integer,
  price integer
);

CREATE TABLE IF NOT EXISTS warehouse.stock(
  item_id varchar(256) PRIMARY KEY,
  warehouse_id varchar(256),
  stock integer
);

CREATE TABLE IF NOT EXISTS warehouse.stock_control(
  stock_control_id varchar(256) PRIMARY KEY,
  warehouse_id varchar(256),
  description varchar(120),
  start_date_of_control TIMESTAMP WITHOUT TIME ZONE,
  end_date_of_control TIMESTAMP WITHOUT TIME ZONE,
  person_in_charge_id varchar(256),
  notes text
);

CREATE TABLE IF NOT EXISTS warehouse.stock_control_line(
  stock_control_line_id varchar(256) PRIMARY KEY,
  stock_control_id varchar(256),
  item_id varchar(256),
  unique_identifier varchar(500),
  digital_identifier bytea,
  stock integer
);

--********************WAREHOUSE:DELIVERY DATA*************

CREATE TABLE IF NOT EXISTS warehouse.shipping_address(
  address_id varchar(256) PRIMARY KEY,
  person_id varchar(256),
  country_iso_3 varchar(256),
  administrative_area varchar(120),
  administrative_area_iso varchar(10),
  locality varchar(120),
  dependent_locality varchar(120),
  postal_code varchar(20),
  street varchar(120),
  street_number integer,
  street_type varchar(20),
  premise varchar(120),
  sub_premise varchar(120),
  department varchar(120),
  floor varchar(4),
  letter varchar(20),
  stairs varchar(20),
  decimal_degrees_latitude varchar(20),
  decimal_degrees_longitude varchar(20)
);

CREATE TABLE IF NOT EXISTS warehouse.country (
  country_iso_3 varchar(256) primary key,
  alpha_2 varchar(2),
  english_name varchar(256),
  spanish_name varchar(256)
);

CREATE TABLE IF NOT EXISTS warehouse.email_address (
  email_address_id varchar(256) PRIMARY KEY,
  person_id varchar(256),
  email_address varchar(120),
  notes varchar(500)
);

CREATE TABLE IF NOT EXISTS warehouse.phone_number (
  phone_number_id varchar(256) PRIMARY KEY,
  person_id varchar(256),
  phone_number varchar(120),
  notes varchar(500)
);

INSERT INTO warehouse.country (english_name, spanish_name, alpha_2, country_iso_3) VALUES
  ('','Afganistán','af','afg'),
  ('','Albania','al','alb'),
  ('','Alemania','de','deu'),
  ('','Andorra','ad','and'),
  ('','Angola','ao','ago'),
  ('','Antigua y Barbuda','ag','atg'),
  ('','Arabia Saudita','sa','sau'),
  ('','Argelia','dz','dza'),
  ('','Argentina','ar','arg'),
  ('','Armenia','am','arm'),
  ('','Australia','au','aus'),
  ('','Austria','at','aut'),
  ('','Azerbaiyán','az','aze'),
  ('','Bahamas','bs','bhs'),
  ('','Bangladés','bd','bgd'),
  ('','Barbados','bb','brb'),
  ('','Baréin','bh','bhr'),
  ('','Bélgica','be','bel'),
  ('','Belice','bz','blz'),
  ('','Benín','bj','ben'),
  ('','Bielorrusia','by','blr'),
  ('Bolivia (Plurinational State of)','Bolivia','bo','bol'),
  ('','Bosnia y Herzegovina','ba','bih'),
  ('','Botsuana','bw','bwa'),
  ('Brazil','Brasil','br','bra'),
  ('','Brunéi','bn','brn'),
  ('','Bulgaria','bg','bgr'),
  ('','Burkina Faso','bf','bfa'),
  ('','Burundi','bi','bdi'),
  ('','Bután','bt','btn'),
  ('','Cabo Verde','cv','cpv'),
  ('','Camboya','kh','khm'),
  ('','Camerún','cm','cmr'),
  ('','Canadá','ca','can'),
  ('','Catar','qa','qat'),
  ('','Chad','td','tcd'),
  ('Chile','Chile','cl','chl'),
  ('','China','cn','chn'),
  ('','Chipre','cy','cyp'),
  ('','Colombia','co','col'),
  ('','Comoras','km','com'),
  ('','Corea del Norte','kp','prk'),
  ('','Corea del Sur','kr','kor'),
  ('','Costa de Marfil','ci','civ'),
  ('','Costa Rica','cr','cri'),
  ('','Croacia','hr','hrv'),
  ('','Cuba','cu','cub'),
  ('','Dinamarca','dk','dnk'),
  ('','Dominica','dm','dma'),
  ('','Ecuador','ec','ecu'),
  ('','Egipto','eg','egy'),
  ('','El Salvador','sv','slv'),
  ('','Emiratos Árabes Unidos','ae','are'),
  ('','Eritrea','er','eri'),
  ('','Eslovaquia','sk','svk'),
  ('','Eslovenia','si','svn'),
  ('Spain','España','es','esp'),
  ('','Estados Unidos','us','usa'),
  ('','Estonia','ee','est'),
  ('','Etiopía','et','eth'),
  ('','Filipinas','ph','phl'),
  ('','Finlandia','fi','fin'),
  ('','Fiyi','fj','fji'),
  ('France','Francia','fr','fra'),
  ('','Gabón','ga','gab'),
  ('','Gambia','gm','gmb'),
  ('','Georgia','ge','geo'),
  ('','Ghana','gh','gha'),
  ('','Granada','gd','grd'),
  ('','Grecia','gr','grc'),
  ('','Guatemala','gt','gtm'),
  ('','Guinea','gn','gin'),
  ('','Guinea-Bisáu','gw','gnb'),
  ('','Guinea Ecuatorial','gq','gnq'),
  ('','Guyana','gy','guy'),
  ('','Haití','ht','hti'),
  ('','Honduras','hn','hnd'),
  ('','Hungría','hu','hun'),
  ('','India','in','ind'),
  ('','Indonesia','id','idn'),
  ('','Irak','iq','irq'),
  ('','Irán','ir','irn'),
  ('','Irlanda','ie','irl'),
  ('','Islandia','is','isl'),
  ('','Islas Marshall','mh','mhl'),
  ('','Islas Salomón','sb','slb'),
  ('','Israel','il','isr'),
  ('','Italia','it','ita'),
  ('','Jamaica','jm','jam'),
  ('','Japón','jp','jpn'),
  ('','Jordania','jo','jor'),
  ('','Kazajistán','kz','kaz'),
  ('','Kenia','ke','ken'),
  ('','Kirguistán','kg','kgz'),
  ('','Kiribati','ki','kir'),
  ('','Kuwait','kw','kwt'),
  ('','Laos','la','lao'),
  ('','Lesoto','ls','lso'),
  ('','Letonia','lv','lva'),
  ('','Líbano','lb','lbn'),
  ('','Liberia','lr','lbr'),
  ('','Libia','ly','lby'),
  ('','Liechtenstein','li','lie'),
  ('','Lituania','lt','ltu'),
  ('','Luxemburgo','lu','lux'),
  ('','Macedonia del Norte','mk','mkd'),
  ('','Madagascar','mg','mdg'),
  ('','Malasia','my','mys'),
  ('','Malaui','mw','mwi'),
  ('','Maldivas','mv','mdv'),
  ('','Malí','ml','mli'),
  ('','Malta','mt','mlt'),
  ('','Marruecos','ma','mar'),
  ('','Mauricio','mu','mus'),
  ('','Mauritania','mr','mrt'),
  ('','México','mx','mex'),
  ('','Micronesia','fm','fsm'),
  ('','Moldavia','md','mda'),
  ('','Mónaco','mc','mco'),
  ('','Mongolia','mn','mng'),
  ('','Montenegro','me','mne'),
  ('','Mozambique','mz','moz'),
  ('','Birmania','mm','mmr'),
  ('','Namibia','na','nam'),
  ('','Nauru','nr','nru'),
  ('','Nepal','np','npl'),
  ('','Nicaragua','ni','nic'),
  ('','Níger','ne','ner'),
  ('','Nigeria','ng','nga'),
  ('','Noruega','no','nor'),
  ('','Nueva Zelanda','nz','nzl'),
  ('','Omán','om','omn'),
  ('','Países Bajos','nl','nld'),
  ('','Pakistán','pk','pak'),
  ('','Palaos','pw','plw'),
  ('','Panamá','pa','pan'),
  ('','Papúa Nueva Guinea','pg','png'),
  ('','Paraguay','py','pry'),
  ('','Perú','pe','per'),
  ('','Polonia','pl','pol'),
  ('Portugal','Portugal','pt','prt'),
  ('United Kingdom','Reino Unido','gb','gbr'),
  ('','República Centroafricana','cf','caf'),
  ('','República Checa','cz','cze'),
  ('','República del Congo','cg','cog'),
  ('','República Democrática del Congo','cd','cod'),
  ('','República Dominicana','do','dom'),
  ('','Ruanda','rw','rwa'),
  ('','Rumania','ro','rou'),
  ('','Rusia','ru','rus'),
  ('','Samoa','ws','wsm'),
  ('','San Cristóbal y Nieves','kn','kna'),
  ('','San Marino','sm','smr'),
  ('','San Vicente y las Granadinas','vc','vct'),
  ('','Santa Lucía','lc','lca'),
  ('','Santo Tomé y Príncipe','st','stp'),
  ('','Senegal','sn','sen'),
  ('','Serbia','rs','srb'),
  ('','Seychelles','sc','syc'),
  ('','Sierra Leona','sl','sle'),
  ('','Singapur','sg','sgp'),
  ('','Siria','sy','syr'),
  ('','Somalia','so','som'),
  ('','Sri Lanka','lk','lka'),
  ('','Suazilandia','sz','swz'),
  ('','Sudáfrica','za','zaf'),
  ('','Sudán','sd','sdn'),
  ('','Sudán del Sur','ss','ssd'),
  ('','Suecia','se','swe'),
  ('','Suiza','ch','che'),
  ('','Surinam','sr','sur'),
  ('','Tailandia','th','tha'),
  ('','Tanzania','tz','tza'),
  ('','Tayikistán','tj','tjk'),
  ('','Timor Oriental','tl','tls'),
  ('','Togo','tg','tgo'),
  ('','Tonga','to','ton'),
  ('','Trinidad y Tobago','tt','tto'),
  ('','Túnez','tn','tun'),
  ('','Turkmenistán','tm','tkm'),
  ('','Turquía','tr','tur'),
  ('','Tuvalu','tv','tuv'),
  ('','Ucrania','ua','ukr'),
  ('','Uganda','ug','uga'),
  ('','Uruguay','uy','ury'),
  ('','Uzbekistán','uz','uzb'),
  ('','Vanuatu','vu','vut'),
  ('','Venezuela','ve','ven'),
  ('','Vietnam','vn','vnm'),
  ('','Yemen','ye','yem'),
  ('','Yibuti','dj','dji'),
  ('','Zambia','zm','zmb'),
  ('','Zimbabue','zw','zwe');
