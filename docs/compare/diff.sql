-- Patch script: bring old database schema up to new.sql for the tables present in new.sql
-- Differences found:
--   mailbox_user : missing column admin_ind
--   session      : table does not exist

ALTER TABLE `mailbox_user`
    ADD COLUMN `admin_ind` varchar(1) NOT NULL DEFAULT 'N';

CREATE TABLE `session`
(
    `session_id`  varchar(64)  NOT NULL,
    `user_id`     varchar(64)  NOT NULL,
    `token`       varchar(128) NOT NULL,
    `expired_ind` varchar(1)   NOT NULL,
    `expiry_date` datetime     NOT NULL,
    `data`        mediumtext   NOT NULL,
    CONSTRAINT `fk_session_user`
        FOREIGN KEY (`user_id`) REFERENCES `mailbox_user` (`user_id`)
) COLLATE = utf8mb4_general_ci;
