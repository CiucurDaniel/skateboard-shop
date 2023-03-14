-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: catalogdb
-- ------------------------------------------------------
-- Server version	8.0.28

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
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `price` bigint DEFAULT NULL,
  `quantity` bigint DEFAULT NULL,
  `image_url` longtext,
  `description` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'2022-06-28 13:27:01.106','2022-06-28 13:27:01.106',NULL,'Cliche Pro',140,4,'skate_cliche.PNG','A nice skateboard'),(2,'2022-06-28 13:39:49.643','2022-06-28 13:39:49.643',NULL,'Baker',200,4,'skate_baker.PNG','A nice skateboard'),(3,'2022-06-28 13:40:12.475','2022-06-28 13:40:12.475',NULL,'Toy Machine',240,4,'skate_toy_machine.PNG','A nice skateboard'),(4,'2022-06-28 13:42:34.149','2022-06-28 13:42:34.149',NULL,'Element',130,4,'skate_element.PNG','A nice skateboard'),(5,'2022-06-28 13:53:37.552','2022-06-28 13:53:37.552',NULL,'Jart',140,4,'skate_jart.PNG','A nice skateboard');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-06  1:01:11
