create:
	sqlite3 contact.sqlite < ./create_table.sql
load:
	sqlite3 contact.sqlite < ./load_data.sql