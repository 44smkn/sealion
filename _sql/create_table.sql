
CREATE TABLE IF NOT EXISTS Tasks 
(
	 id        INTEGER PRIMARY KEY AUTO_INCREMENT
	,category  VARCHAR(255)
    ,name      VARCHAR(255)
	,do_today  boolean
	,deadline  DATE
	,ticket_id VARCHAR(255)
	,archive   boolean
);