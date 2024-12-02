drop database maladie_chronique;
create database maladie_chronique;
use maladie_chronique;
-- MySQL dump 10.13  Distrib 8.0.40, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: maladie_chronique
-- ------------------------------------------------------
-- Server version	8.0.40

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
-- Table structure for table `assurance_maladie_medicament`
--

DROP TABLE IF EXISTS `assurance_maladie_medicament`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assurance_maladie_medicament` (
  `id` int NOT NULL AUTO_INCREMENT,
  `assure_id` int DEFAULT NULL,
  `maladie_id` int DEFAULT NULL,
  `medicament_id` int DEFAULT NULL,
  `prix_medic` decimal(10,2) DEFAULT NULL,
  `date_achat_medic` varchar(15) DEFAULT NULL,
  `commune` varchar(100) DEFAULT NULL,
  `wilaya` varchar(100) DEFAULT NULL,
  `region` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `assure_id` (`assure_id`),
  KEY `maladie_id` (`maladie_id`),
  KEY `medicament_id` (`medicament_id`),
  CONSTRAINT `assurance_maladie_medicament_ibfk_1` FOREIGN KEY (`assure_id`) REFERENCES `assuré` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `assurance_maladie_medicament_ibfk_2` FOREIGN KEY (`maladie_id`) REFERENCES `maladie` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `assurance_maladie_medicament_ibfk_3` FOREIGN KEY (`medicament_id`) REFERENCES `medicament` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assurance_maladie_medicament`
--

LOCK TABLES `assurance_maladie_medicament` WRITE;
/*!40000 ALTER TABLE `assurance_maladie_medicament` DISABLE KEYS */;
/*!40000 ALTER TABLE `assurance_maladie_medicament` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `assurance_maladie_prestation`
--

DROP TABLE IF EXISTS `assurance_maladie_prestation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assurance_maladie_prestation` (
  `id` int NOT NULL AUTO_INCREMENT,
  `assure_id` int DEFAULT NULL,
  `maladie_id` int DEFAULT NULL,
  `prestation_id` int DEFAULT NULL,
  `prix_prest` decimal(10,2) DEFAULT NULL,
  `date_prestation` varchar(15) DEFAULT NULL,
  `commune` varchar(100) DEFAULT NULL,
  `wilaya` varchar(100) DEFAULT NULL,
  `region` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `assure_id` (`assure_id`),
  KEY `maladie_id` (`maladie_id`),
  KEY `prestation_id` (`prestation_id`),
  CONSTRAINT `assurance_maladie_prestation_ibfk_1` FOREIGN KEY (`assure_id`) REFERENCES `assuré` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `assurance_maladie_prestation_ibfk_2` FOREIGN KEY (`maladie_id`) REFERENCES `maladie` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `assurance_maladie_prestation_ibfk_3` FOREIGN KEY (`prestation_id`) REFERENCES `prestation` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assurance_maladie_prestation`
--

LOCK TABLES `assurance_maladie_prestation` WRITE;
/*!40000 ALTER TABLE `assurance_maladie_prestation` DISABLE KEYS */;
/*!40000 ALTER TABLE `assurance_maladie_prestation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `assuré`
--

DROP TABLE IF EXISTS `assuré`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assuré` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nom_prenom` varchar(100) DEFAULT NULL,
  `dateN` varchar(100) DEFAULT NULL,
  `lieuN` varchar(100) DEFAULT NULL,
  `genre` varchar(50) DEFAULT NULL,
  `lieu_res` varchar(150) DEFAULT NULL,
  `commune` varchar(100) DEFAULT NULL,
  `wilaya` varchar(100) DEFAULT NULL,
  `region` varchar(100) DEFAULT NULL,
  `num_tel` varchar(15) DEFAULT NULL,
  `profession` varchar(100) DEFAULT NULL,
  `secteur` varchar(100) DEFAULT NULL,
  `revenu` decimal(10,2) DEFAULT NULL,
  `etat_civil` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assuré`
--

LOCK TABLES `assuré` WRITE;
/*!40000 ALTER TABLE `assuré` DISABLE KEYS */;
INSERT INTO `assuré` VALUES (1,'rachid LARBI','199439','Aïn El Hammam','Femme','Taourirt','Taourirt','Khenchela','Est','0662716150','Médecin','technologie',556834.00,'Divorcé(e)');
/*!40000 ALTER TABLE `assuré` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maladie`
--

DROP TABLE IF EXISTS `maladie`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `maladie` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code_maladie` varchar(20) DEFAULT NULL,
  `libelle` varchar(100) DEFAULT NULL,
  `taux` decimal(5,2) DEFAULT NULL,
  `type_maladie` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `code_maladie` (`code_maladie`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maladie`
--

LOCK TABLES `maladie` WRITE;
/*!40000 ALTER TABLE `maladie` DISABLE KEYS */;
/*!40000 ALTER TABLE `maladie` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maladie_medicament`
--

DROP TABLE IF EXISTS `maladie_medicament`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `maladie_medicament` (
  `id` int NOT NULL AUTO_INCREMENT,
  `maladie_id` int DEFAULT NULL,
  `medicament_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `maladie_id` (`maladie_id`),
  KEY `medicament_id` (`medicament_id`),
  CONSTRAINT `maladie_medicament_ibfk_1` FOREIGN KEY (`maladie_id`) REFERENCES `maladie` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `maladie_medicament_ibfk_2` FOREIGN KEY (`medicament_id`) REFERENCES `medicament` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maladie_medicament`
--

LOCK TABLES `maladie_medicament` WRITE;
/*!40000 ALTER TABLE `maladie_medicament` DISABLE KEYS */;
/*!40000 ALTER TABLE `maladie_medicament` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `medicament`
--

DROP TABLE IF EXISTS `medicament`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `medicament` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(100) DEFAULT NULL,
  `famille` varchar(50) DEFAULT NULL,
  `TR` decimal(5,2) DEFAULT NULL,
  `marque` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `medicament`
--

LOCK TABLES `medicament` WRITE;
/*!40000 ALTER TABLE `medicament` DISABLE KEYS */;
/*!40000 ALTER TABLE `medicament` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `prestation`
--

DROP TABLE IF EXISTS `prestation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `prestation` (
  `id` int NOT NULL AUTO_INCREMENT,
  `libelle` varchar(100) DEFAULT NULL,
  `type_pres` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prestation`
--

LOCK TABLES `prestation` WRITE;
/*!40000 ALTER TABLE `prestation` DISABLE KEYS */;
/*!40000 ALTER TABLE `prestation` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-12-02  9:57:35
