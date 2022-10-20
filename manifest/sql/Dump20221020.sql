CREATE DATABASE  IF NOT EXISTS `election` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `election`;
-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: election
-- ------------------------------------------------------
-- Server version	8.0.31

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
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(45) DEFAULT NULL COMMENT '用户名',
  `password` varchar(45) DEFAULT NULL COMMENT '密码',
  `role` int DEFAULT '1' COMMENT '账号角色\n0 - 超级管理员\n1 - 管理员',
  `status` int DEFAULT '1' COMMENT '账号状态 ： \n0 - 禁用      \n1 - 可用',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (1,'admin','660806b01437a89ad01f5abca8c1d099',0,1,'2022-10-16 07:00:54','2022-10-16 07:00:54');
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `candidates`
--

DROP TABLE IF EXISTS `candidates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `candidates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL COMMENT '候选人名称',
  `introduction` varchar(1024) DEFAULT NULL COMMENT '候选人简介',
  `image` varchar(255) DEFAULT NULL COMMENT '候选人简介',
  `status` int DEFAULT '1' COMMENT '候选人状态\n1 - 可用\n0 - 不可用',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='候选人';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `candidates`
--

LOCK TABLES `candidates` WRITE;
/*!40000 ALTER TABLE `candidates` DISABLE KEYS */;
INSERT INTO `candidates` VALUES (1,'候选人1','候选人1','http://cdn.cocheer.cn/imgs/15655790132407c8b75ecb36ca2933b4073f9325bb217.png',1,'2022-10-18 15:49:51','2022-10-18 15:49:51'),(2,'候选人2','候选人2','http://cdn.cocheer.cn/imgs/15655790132407c8b75ecb36ca2933b4073f9325bb217.png',1,'2022-10-19 15:57:29','2022-10-19 15:57:29');
/*!40000 ALTER TABLE `candidates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `election_config_candidates`
--

DROP TABLE IF EXISTS `election_config_candidates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `election_config_candidates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `candidateId` int DEFAULT '0',
  `electionId` int DEFAULT '0',
  `voteCount` int DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `index_candidate_election` (`candidateId`,`electionId`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `election_config_candidates`
--

LOCK TABLES `election_config_candidates` WRITE;
/*!40000 ALTER TABLE `election_config_candidates` DISABLE KEYS */;
INSERT INTO `election_config_candidates` VALUES (7,1,7,3),(8,2,7,0);
/*!40000 ALTER TABLE `election_config_candidates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `election_user_details`
--

DROP TABLE IF EXISTS `election_user_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `election_user_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `electionId` int DEFAULT NULL,
  `candidateId` int DEFAULT NULL,
  `idCard` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `index_details` (`electionId`,`candidateId`,`email`,`idCard`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='选举用户详情';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `election_user_details`
--

LOCK TABLES `election_user_details` WRITE;
/*!40000 ALTER TABLE `election_user_details` DISABLE KEYS */;
INSERT INTO `election_user_details` VALUES (1,7,1,'a111111(1)','1160627439@qq.com','2022-10-19 12:55:27'),(2,7,1,'a111111(2)','zhangzheming@cocheer.net','2022-10-19 12:57:34'),(3,7,1,'a111111(3)','1153@qq.com','2022-10-20 06:14:00');
/*!40000 ALTER TABLE `election_user_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `elections`
--

DROP TABLE IF EXISTS `elections`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `elections` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(45) DEFAULT NULL COMMENT '选举标题',
  `introduction` varchar(1024) DEFAULT NULL COMMENT '''简介''',
  `status` int DEFAULT '0' COMMENT '状态\n0 - 未开始\n1 - 进行中\n2 - 已结束',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `elections`
--

LOCK TABLES `elections` WRITE;
/*!40000 ALTER TABLE `elections` DISABLE KEYS */;
INSERT INTO `elections` VALUES (7,'选举测试1','选举测试2',1);
/*!40000 ALTER TABLE `elections` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-20 14:49:14
