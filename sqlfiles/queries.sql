USE ioe;

/*
SET @@time_zone='SYSTEM';

/* this will delete transactions every 1 second 
CREATE EVENT delete_transaction_every_second
ON SCHEDULE EVERY 1 SECOND
STARTS NOW()
ON COMPLETION PRESERVE
DO DELETE FROM transactions ;
*/

INSERT INTO users (`NAME`,`EMAIL_ID`, `USERNAME`, `PASSWORD`,`USER_KEY`) VALUES ("John Evans","john123@gmail.com","jevans","jeeee","231");

/*TIME STAMP FORMAT yyyy-mm-dd hh:mm:ss*/
INSERT INTO transactions (`USER_ID`,`WATT/SECOND`, `TYPE`,`TIMESTAMP`) VALUES (1,123.25,"GIVE",'2020-06-08 06:08:57');

INSERT INTO transactions (`USER_ID`,`WATT/SECOND`, `TYPE`,`TIMESTAMP`) VALUES (1,243.25,"TAKE",'2020-06-08 06:08:57');
/*DELETE FROM transactions WHERE `TIMESTAMP` < (NOW() - INTERVAL 1 SECOND);*/
INSERT INTO transactions (`USER_ID`,`WATT/SECOND`, `TYPE`) VALUES (1,243.25,"TAKE"); 

INSERT INTO transactions (`USER_ID`,`WATT/SECOND`, `TYPE`) VALUES (2,12.5,"TAKE"); 

/* For the unit tests to work create a separate database called test and add all the tables from the database.sql file.
After adding the tables add the following admin and user into the database*/

INSERT INTO users (`NAME`,`EMAIL_ID`, `USERNAME`,`PASSWORD`,`USER_KEY`) VALUES ("Chris George","chrisg@gmail.com","chrisg1","MyNameIsChris","ETU3rSTYCnqM51lsgiZMXI5Y4B4sDNxBodsnaImKBgesNcbpf09JbwnFKurCL5zObqihGyDJEJrLXZxPXjvYM1Pe1psqbh4jpHADxjSZYZ8Pey2loQDByDBzdtYyDp8skkD7c3M5tVwWGSzu05zoJxOA8scQtwbgryFhErrGHKTAvUQ3hbRgnOEaj191mP4A7swVOQKqorU8OBTrlmj6W49IPzd0Cp85ZJKtXb4H1HVzR9v39wLFzBeaRGjOQ0EKIGdy3iiKzzLZeIKzy58PjgK2UF8aDw3YaRU9TJILy4q93xNBQJA9xh59HZ3mqJGUfyEUOC15sqEimxPflwrurewHBrc0GO1AjBYwYw4fLOmzgXUXrPBjCsxpTtHkDXzIdf9FqSG4q5BqmdqsDVU5FGcllHvKmhb9Gm2U9DHRniNJ9bLwLMNX1DpIQxBrgrT1Bnzrn1o80fDOqZwSc8KjRWzQpqxxbchlEQCqH8fz12KABRSPzs0k");
INSERT INTO admins (`ADMIN_NAME`,`EMAIL_ID`, `USERNAME`, `PASSWORD`, `ADMIN_KEY`) VALUES ("Jack Black","jb12@gmail.com","jbl","jackbl","ETU3rSTYCnqM51lsgiZMXI5Y4B4sDNxBodsnaImKBgesNcbpf09JbwnFKurCL5zObqihGyDJEJrLXZxPXjvYM1Pe1psqbh4jpHADxjSZYZ8Pey2loQDByDBzdtYyDp8skkD7c3M5tVwWGSzu05zoJxOA8scQtwbgryFhErrGHKTAvUQ3hbRgnOEaj191mP4A7swVOQKqorU8OBTrlmj6W49IPzd0Cp85ZJKtXb4H1HVzR9v39wLFzBeaRGjOQ0EKIGdy3iiKzzLZeIKzy58PjgK2UF8aDw3YaRU9TJILy4q93xNBQJA9xh59HZ3mqJGUfyEUOC15sqEimxPflwrurewHBrc0GO1AjBYwYw4fLOmzgXUXrPBjCsxpTtHkDXzIdf9FqSG4q5BqmdqsDVU5FGcllHvKmhb9Gm2U9DHRniNJ9bLwLMNX1DpIQxBrgrT1Bnzrn1o80fDOqZwSc8KjRWzQpqxxbchlEQCqH8fz12KABRSPzs0k");
















