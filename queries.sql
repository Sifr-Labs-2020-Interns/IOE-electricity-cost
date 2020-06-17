USE ioe;

SET @@time_zone='SYSTEM';

/* this will delete transactions every 1 second */
CREATE EVENT delete_transaction_every_second
ON SCHEDULE EVERY 1 SECOND
STARTS NOW()
ON COMPLETION PRESERVE
DO DELETE FROM transactions ;

INSERT INTO users (`NAME`,`EMAIL_ID`, `USERNAME`,`KEY`) VALUES ("John Evans","john123@gmail.com","jevans","231");
INSERT INTO users (`NAME`,`EMAIL_ID`, `USERNAME`,`KEY`) VALUES ("Chris George","chrisg@gmail.com","chrisg1","452");

/*TIME STAMP FORMAT yyyy-mm-dd hh:mm:ss*/
INSERT INTO transactions (`USER_ID`,`WATT/SECOND`, `TYPE`,`TIMESTAMP`) VALUES (1,123.25,"GIVE",'2020-06-08 06:08:57');

INSERT INTO transactions (`USER_ID`,`WATT/SECOND`, `TYPE`,`TIMESTAMP`) VALUES (1,243.25,"TAKE",'2020-06-08 06:08:57');
/*DELETE FROM transactions WHERE `TIMESTAMP` < (NOW() - INTERVAL 1 SECOND);*/
INSERT INTO transactions (`USER_ID`,`WATT/SECOND`, `TYPE`) VALUES (1,243.25,"TAKE"); 

      
      
INSERT INTO transactions (`USER_ID`,`WATT/SECOND`, `TYPE`) VALUES (2,12.5,"TAKE"); 




















