package sqlite

const (
	QueryCreateOrderTable = `
	create table if not exists experiment
(
    id          INTEGER not null
        constraint orders_pk
            primary key autoincrement,
    SSID TEXT,
    RSSI INTEGER,
    CurrentTime TEXT
);`
	QueryGetData = `
SELECT id, SSID, RSSI, CurrentTime FROM experiment`
	QueryInsertData = `
	INSERT INTO experiment (SSID, RSSI, CurrentTime)
	VALUES ($1, $2, $3);`
)
