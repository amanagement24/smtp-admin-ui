/*M!999999\- enable the sandbox mode */
-- MariaDB dump 10.19  Distrib 10.11.14-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: smtp
-- ------------------------------------------------------
-- Server version	10.11.14-MariaDB-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `domain`
--

DROP TABLE IF EXISTS `domain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `domain` (
                          `domain_id` varchar(64) NOT NULL,
                          `name` varchar(255) NOT NULL,
                          `catchall_ind` varchar(1) NOT NULL DEFAULT 'N',
                          `catchall_login` varchar(255) NOT NULL DEFAULT '',
                          PRIMARY KEY (`domain_id`),
                          UNIQUE KEY `ix_domain_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mailbox`
--

DROP TABLE IF EXISTS `mailbox`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `mailbox` (
                           `mailbox_id` varchar(64) NOT NULL,
                           `user_id` varchar(64) NOT NULL,
                           `name` varchar(255) NOT NULL,
                           `flag_non_existent` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_no_inferiors` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_no_select` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_marked` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_subscribed` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_remote` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_archive` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_drafts` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_flagged` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_junk` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_sent` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_trash` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_important` varchar(1) NOT NULL DEFAULT 'N',
                           PRIMARY KEY (`mailbox_id`),
                           KEY `fk_mailbox_mailbox_user` (`user_id`),
                           CONSTRAINT `fk_mailbox_mailbox_user` FOREIGN KEY (`user_id`) REFERENCES `mailbox_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mailbox_user`
--

DROP TABLE IF EXISTS `mailbox_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `mailbox_user` (
                                `user_id` varchar(64) NOT NULL,
                                `domain_id` varchar(64) NOT NULL,
                                `login` varchar(255) NOT NULL,
                                `password` varchar(255) NOT NULL,
                                PRIMARY KEY (`user_id`),
                                UNIQUE KEY `ix_mailbox_user` (`user_id`,`login`),
                                KEY `fk_mailbox_user_domain` (`domain_id`),
                                CONSTRAINT `fk_mailbox_user_domain` FOREIGN KEY (`domain_id`) REFERENCES `domain` (`domain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `message` (
                           `message_id` varchar(64) NOT NULL,
                           `mailbox_id` varchar(64) NOT NULL,
                           `body` longblob DEFAULT NULL,
                           `uid` int(11) NOT NULL DEFAULT 0,
                           `created_date` datetime NOT NULL DEFAULT current_timestamp(),
                           `flag_seen` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_answered` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_flagged` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_deleted` varchar(1) NOT NULL DEFAULT 'N',
                           `flag_draft` varchar(1) NOT NULL DEFAULT 'N',
                           `spam_score` decimal(5,2) NOT NULL DEFAULT 0.00,
                           PRIMARY KEY (`message_id`),
                           KEY `fk_message_mailbox` (`mailbox_id`),
                           CONSTRAINT `fk_message_mailbox` FOREIGN KEY (`mailbox_id`) REFERENCES `mailbox` (`mailbox_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `queue`
--

DROP TABLE IF EXISTS `queue`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `queue` (
                         `queue_id` varchar(64) NOT NULL,
                         `from_addr` varchar(255) NOT NULL,
                         `body` longtext DEFAULT NULL,
                         PRIMARY KEY (`queue_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `queue_recipient`
--

DROP TABLE IF EXISTS `queue_recipient`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `queue_recipient` (
                                   `queue_recipient_id` varchar(64) NOT NULL,
                                   `queue_id` varchar(64) NOT NULL,
                                   `to_addr` varchar(255) NOT NULL,
                                   `attempts` int(11) NOT NULL DEFAULT 0,
                                   `last_attempted_dt` datetime DEFAULT NULL,
                                   `success_ind` varchar(1) NOT NULL DEFAULT 'N',
                                   PRIMARY KEY (`queue_recipient_id`),
                                   KEY `fk_qrec_queue` (`queue_id`),
                                   CONSTRAINT `fk_qrec_queue` FOREIGN KEY (`queue_id`) REFERENCES `queue` (`queue_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-06-01 22:03:18
