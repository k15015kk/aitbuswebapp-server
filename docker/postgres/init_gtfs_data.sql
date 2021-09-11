
COPY agency FROM '/gtfs/agency.txt' WITH csv header;
COPY calendar_dates FROM '/gtfs/calendar_dates.txt' WITH csv header;
COPY fare_attributes FROM '/gtfs/fare_attributes.txt' WITH csv header;
COPY fare_rules FROM '/gtfs/fare_rules.txt' WITH csv header;
COPY routes FROM '/gtfs/routes.txt' WITH csv header;
COPY stop_times FROM '/gtfs/stop_times.txt' WITH csv header;
COPY stops FROM '/gtfs/stops.txt' WITH csv header;
COPY trips FROM '/gtfs/trips.txt' WITH csv header;

