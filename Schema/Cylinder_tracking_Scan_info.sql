-- MySQL dump 10.13  Distrib 8.0.20, for Linux (x86_64)
--
-- Host: 192.168.43.140    Database: Cylinder_tracking
-- ------------------------------------------------------
-- Server version	8.0.20-0ubuntu0.20.04.1

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
-- Table structure for table `Scan_info`
--

DROP TABLE IF EXISTS `Scan_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Scan_info` (
  `SI` int NOT NULL,
  `Scanned Location` varchar(45) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `Email` varchar(45) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `Matched location` varchar(45) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `Product ID` varchar(45) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `Product SKU` varchar(45) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `System Batched No` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `System filled date` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `System Delivered location` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `System Delivered Date` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `Checked` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`Product ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Scan_info`
--

LOCK TABLES `Scan_info` WRITE;
/*!40000 ALTER TABLE `Scan_info` DISABLE KEYS */;
INSERT INTO `Scan_info` VALUES (1,'Dhaka','kajol@gmail.com','YES','11212','123123','3232','12-12-20','dhaka','14-12-20','YES');
/*!40000 ALTER TABLE `Scan_info` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-06-24 14:18:46
