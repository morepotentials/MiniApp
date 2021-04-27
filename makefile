DATA_DIR=DB
refresh_db: stop_db clean_db create_db start_db init_db

start_db: 
	pg_ctl -D $(DATA_DIR) -l logfile start

stop_db:
	pg_ctl -D $(DATA_DIR) stop

create_db:
	initdb $(DATA_DIR)

clean_db:
	rm -rf $(DATA_DIR)

init_db:
	psql -d postgres -f migration_app.sql
	psql -d postgres -f dummy_data_app.sql

login:
	psql postgres