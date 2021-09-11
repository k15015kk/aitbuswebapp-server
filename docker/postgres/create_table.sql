DROP SCHEMA IF EXISTS gtfs;

CREATE SCHEMA gtfs;

CREATE TABLE IF NOT EXISTS agency (
    agency_id VARCHAR(64) NOT NULL,
    agency_name VARCHAR(255) NOT NULL ,
    agency_url VARCHAR(255) NOT NULL ,
    agency_timezone VARCHAR(32) NOT NULL,
    agency_lang VARCHAR(32) NOT NULL,
    agency_phone VARCHAR(32) NOT NULL,

    PRIMARY KEY (agency_id)
);

CREATE TABLE IF NOT EXISTS calendar_dates (
    service_id VARCHAR(64) NOT NULL,
    date VARCHAR(8),
    exception_type SMALLINT,

    PRIMARY KEY (service_id,date)
);

CREATE TABLE IF NOT EXISTS fare_attributes (
    fare_id VARCHAR(64) NOT NULL,
    price INTEGER NOT NULL,
    currency_type VARCHAR(8) NOT NULL,
    payment_method SMALLINT NOT NULL,
    transfers SMALLINT NOT NULL,

    PRIMARY KEY (fare_id)
);

CREATE TABLE IF NOT EXISTS fare_rules (
    fare_id VARCHAR(64) NOT NULL,
    route_id VARCHAR(64),
    origin_id VARCHAR(64),
    destination_id VARCHAR(64),

    PRIMARY KEY (fare_id, route_id)
);

CREATE TABLE IF NOT EXISTS routes (
    route_id VARCHAR(64) NOT NULL,
    agency_id VARCHAR(64) NOT NULL,
    route_short_name VARCHAR(64) NOT NULL,
    route_long_name VARCHAR(255) NOT NULL,
    route_type SMALLINT NOT NULL,
    route_color VARCHAR(6),
    route_text_color VARCHAR(6),

    PRIMARY KEY (route_id)
);

CREATE INDEX route__route_id
    ON routes(route_id);

CREATE INDEX route__agency_id
    ON routes(agency_id);

CREATE TABLE IF NOT EXISTS stop_times (
    trip_id VARCHAR(64) NOT NULL,
    arrival_time VARCHAR(8) NOT NULL,
    departure_time VARCHAR(8) NOT NULL,
    stop_id  VARCHAR(64),
    stop_sequence INTEGER NOT NULL,
    pickup_type SMALLINT,
    drop_off_type SMALLINT,
    
    PRIMARY KEY (trip_id, stop_sequence, arrival_time)
);

CREATE INDEX stop_times__stop_id
    ON stop_times(stop_id);

CREATE TABLE IF NOT EXISTS stops (
    stop_id VARCHAR(64) NOT NULL,
    stop_name VARCHAR(64) NOT NULL,
    stop_desc VARCHAR(255),
    stop_lat NUMERIC NOT NULL,
    stop_lon NUMERIC NOT NULL,
    zone_id VARCHAR(64),
    location_type SMALLINT,
    parent_station VARCHAR(64),
    platform_code SMALLINT,

    PRIMARY KEY (stop_id)
);

CREATE TABLE IF NOT EXISTS trips (
    route_id VARCHAR(64) NOT NULL,
    service_id VARCHAR(64) NOT NULL,
    trip_id VARCHAR(64) NOT NULL,
    trip_headsign VARCHAR(64),
    direction_id SMALLINT,
    wheelchair_accessible SMALLINT,
    bikes_allowed SMALLINT,

    PRIMARY KEY (trip_id, service_id, direction_id)
);

CREATE INDEX trips__route_id
    ON trips(route_id);

CREATE INDEX trips__service_id
    ON trips(service_id);

CREATE INDEX trips__trip_id
    ON trips(trip_id);