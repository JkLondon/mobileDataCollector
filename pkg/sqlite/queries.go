package sqlite

const (
	QueryCreateOrderTable = `
	create table if not exists experiment
(
    id          INTEGER not null
        constraint orders_pk
            primary key autoincrement,
    RSSI INTEGER,
    CurrentTime TEXT
);`
	QueryGetData = `
SELECT id, RSSI, CurrentTime FROM experiment`
	QueryInsertData = `
	INSERT INTO experiment (RSSI, CurrentTime)
	VALUES ($1, $2);`
)
