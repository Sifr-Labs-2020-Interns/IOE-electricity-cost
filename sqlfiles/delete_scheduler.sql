USE ioe;

SET @@time_zone='SYSTEM';

/* this will delete transactions every 1 second */
CREATE EVENT delete_transaction_every_second
ON SCHEDULE EVERY 1 SECOND
STARTS NOW()
ON COMPLETION PRESERVE
DO DELETE FROM transactions ;
