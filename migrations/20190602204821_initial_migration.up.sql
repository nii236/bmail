CREATE TABLE IF NOT EXISTS users (
	id	    			INTEGER UNIQUE NOT NULL PRIMARY KEY,
	username			VARCHAR(40) UNIQUE NOT NULL,
	bitmessage_id      	VARCHAR(40) UNIQUE NOT NULL,
	archived 			BOOL NOT NULL DEFAULT false,
	created_at			TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS processed_messages (
	id	    		INTEGER UNIQUE NOT NULL PRIMARY KEY,
	encoding_type 	INTEGER,
	to_address    	VARCHAR(200),
	read         	INTEGER,
	msgid        	VARCHAR(200),
	message      	VARCHAR(200),
	from_address  	VARCHAR(200),
	received_time 	VARCHAR(200),
	subject      	VARCHAR(200),
	processed		BOOL NOT NULL DEFAULT FALSE
);