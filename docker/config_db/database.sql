CREATE TABLE IF NOT EXISTS public."History"	(		"ID" 			serial PRIMARY KEY NOT NULL,		"Time"			TIMESTAMP  NULL,		"Multplicator"	DECIMAL  NULL,		"Bet"			DECIMAL  NULL,		"BetCurrency"	DECIMAL  NULL,		"Details"		varchar(255) NULL,		"Currency"		varchar(16) NULL,		"Win"			DECIMAL  NULL,	)	TABLESPACE pg_default;		;  				CREATE TABLE IF NOT EXISTS public."Player"	(		"ID" 			serial PRIMARY KEY NOT NULL,		"Time"			TIMESTAMP  NULL,		"UserName"		varchar(255) NULL,		"Cashout"		DECIMAL NULL,		"Bet"			DECIMAL  NULL,		"Flg"			varchar(16) NULL,		"IsCashedout"	BOOLEAN  NULL,	)	TABLESPACE pg_default;		;  