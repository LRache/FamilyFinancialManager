-- MySQL dump 10.13  Distrib 8.0.43, for macos15 (arm64)
--
-- Host: localhost    Database: newfamily
-- ------------------------------------------------------
-- Server version	9.4.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Temporary view structure for view `categorysummary`
--

DROP TABLE IF EXISTS `categorysummary`;
/*!50001 DROP VIEW IF EXISTS `categorysummary`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `categorysummary` AS SELECT 
 1 AS `familyid`,
 1 AS `familyname`,
 1 AS `categoryid`,
 1 AS `categoryname`,
 1 AS `type`,
 1 AS `total_amount`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `familysummary`
--

DROP TABLE IF EXISTS `familysummary`;
/*!50001 DROP VIEW IF EXISTS `familysummary`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `familysummary` AS SELECT 
 1 AS `familyid`,
 1 AS `familyname`,
 1 AS `total_income`,
 1 AS `total_expense`,
 1 AS `balance`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `membersummary`
--

DROP TABLE IF EXISTS `membersummary`;
/*!50001 DROP VIEW IF EXISTS `membersummary`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `membersummary` AS SELECT 
 1 AS `userid`,
 1 AS `username`,
 1 AS `familyid`,
 1 AS `familyname`,
 1 AS `income`,
 1 AS `expense`,
 1 AS `balance`*/;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `categorysummary`
--

/*!50001 DROP VIEW IF EXISTS `categorysummary`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `categorysummary` AS select `f`.`familyid` AS `familyid`,`f`.`familyname` AS `familyname`,`c`.`categoryid` AS `categoryid`,`c`.`categoryname` AS `categoryname`,`c`.`type` AS `type`,sum(`t`.`amount`) AS `total_amount` from ((`transactionrecord` `t` join `category` `c` on((`t`.`categoryid` = `c`.`categoryid`))) join `family` `f` on((`t`.`familyid` = `f`.`familyid`))) group by `f`.`familyid`,`f`.`familyname`,`c`.`categoryid`,`c`.`categoryname`,`c`.`type` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `familysummary`
--

/*!50001 DROP VIEW IF EXISTS `familysummary`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `familysummary` AS select `f`.`familyid` AS `familyid`,`f`.`familyname` AS `familyname`,sum((case when (`c`.`type` = 1) then `t`.`amount` else 0 end)) AS `total_income`,sum((case when (`c`.`type` = 0) then `t`.`amount` else 0 end)) AS `total_expense`,sum((case when (`c`.`type` = 1) then `t`.`amount` else -(`t`.`amount`) end)) AS `balance` from ((`family` `f` left join `transactionrecord` `t` on((`f`.`familyid` = `t`.`familyid`))) left join `category` `c` on((`t`.`categoryid` = `c`.`categoryid`))) group by `f`.`familyid`,`f`.`familyname` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `membersummary`
--

/*!50001 DROP VIEW IF EXISTS `membersummary`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `membersummary` AS select `u`.`userid` AS `userid`,`u`.`username` AS `username`,`f`.`familyid` AS `familyid`,`f`.`familyname` AS `familyname`,sum((case when (`c`.`type` = 1) then `t`.`amount` else 0 end)) AS `income`,sum((case when (`c`.`type` = 0) then `t`.`amount` else 0 end)) AS `expense`,sum((case when (`c`.`type` = 1) then `t`.`amount` else -(`t`.`amount`) end)) AS `balance` from (((`users` `u` join `family` `f` on((`u`.`familyid` = `f`.`familyid`))) left join `transactionrecord` `t` on((`u`.`userid` = `t`.`userid`))) left join `category` `c` on((`t`.`categoryid` = `c`.`categoryid`))) group by `u`.`userid`,`u`.`username`,`f`.`familyid`,`f`.`familyname` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-09-03  8:54:31
