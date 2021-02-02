-- MySQL dump 10.13  Distrib 8.0.18, for Win64 (x86_64)
--
-- Host: localhost    Database: household
-- ------------------------------------------------------
-- Server version	8.0.16

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
-- Table structure for table `familymember`
--

DROP TABLE IF EXISTS `familymember`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `familymember` (
  `householdID` int(11) NOT NULL,
  `familyMemberID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `gender` varchar(1) NOT NULL,
  `maritalStatus` varchar(10) NOT NULL,
  `spouse` varchar(45) DEFAULT NULL,
  `occupationType` varchar(30) NOT NULL,
  `DOB` date NOT NULL,
  `annualIncome` int(11) NOT NULL,
  PRIMARY KEY (`familyMemberID`),
  KEY `householdID_idx` (`householdID`),
  CONSTRAINT `householdID` FOREIGN KEY (`householdID`) REFERENCES `household` (`householdID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `familymember`
--

LOCK TABLES `familymember` WRITE;
/*!40000 ALTER TABLE `familymember` DISABLE KEYS */;
INSERT INTO `familymember` VALUES (6,8,'child','F','single','','student','2003-02-28',0),(6,9,'mother','F','married','10','employed','2000-01-31',50000),(6,10,'father','F','married','9','employed','2000-01-31',50000),(9,16,'mother2','F','married','father2','employed','1988-01-31',50000),(9,17,'father2','M','married','mother2','employed','1988-01-31',50000),(9,18,'child','M','single','','student','2005-02-28',0),(9,19,'kid','M','single','','student','2016-02-28',0),(11,22,'mary','F','married','john','employed','1990-01-31',100000),(11,23,'john','F','married','mary','employed','1988-01-31',100000),(10,24,'elderly','M','single','','unemployed','1945-01-31',0),(11,25,'child','M','single','','student','2005-02-28',0),(11,29,'elderly2','M','single','','unemployed','1971-01-31',0);
/*!40000 ALTER TABLE `familymember` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-02-02 10:00:26
