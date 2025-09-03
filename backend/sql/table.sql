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
-- Table structure for table `Category`
--

DROP TABLE IF EXISTS `Category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Category` (
  `categoryid` int NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `categoryname` varchar(100) NOT NULL COMMENT '分类名称',
  `type` tinyint(1) NOT NULL COMMENT '收支类型，1=收入，0=支出',
  `note` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`categoryid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='收支分类表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `Family`
--

DROP TABLE IF EXISTS `Family`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Family` (
  `familyid` int NOT NULL AUTO_INCREMENT COMMENT '家庭ID',
  `familyname` varchar(100) NOT NULL COMMENT '家庭名称',
  PRIMARY KEY (`familyid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='家庭表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `TransactionRecord`
--

DROP TABLE IF EXISTS `TransactionRecord`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `TransactionRecord` (
  `transactionrecordid` int NOT NULL AUTO_INCREMENT COMMENT '收支记录ID',
  `familyid` int NOT NULL COMMENT '所属家庭ID',
  `userid` int DEFAULT NULL COMMENT '操作用户ID，可为空表示家庭整体',
  `categoryid` int NOT NULL COMMENT '分类ID',
  `amount` decimal(10,2) NOT NULL COMMENT '金额',
  `occurred_at` datetime NOT NULL COMMENT '发生时间',
  `note` varchar(255) DEFAULT NULL COMMENT '备注',
  `merchant` varchar(100) DEFAULT NULL COMMENT '商家名称',
  `location` varchar(100) DEFAULT NULL COMMENT '消费地点',
  `paymentmethod` varchar(50) DEFAULT NULL COMMENT '支付方式（现金/银行卡/微信/支付宝等）',
  PRIMARY KEY (`transactionrecordid`),
  KEY `FK_Transaction_Family` (`familyid`),
  KEY `FK_Transaction_User` (`userid`),
  KEY `FK_Transaction_Category` (`categoryid`),
  CONSTRAINT `FK_Transaction_Category` FOREIGN KEY (`categoryid`) REFERENCES `Category` (`categoryid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_Transaction_Family` FOREIGN KEY (`familyid`) REFERENCES `Family` (`familyid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_Transaction_User` FOREIGN KEY (`userid`) REFERENCES `Users` (`userid`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='家庭收支记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `Users`
--

DROP TABLE IF EXISTS `Users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Users` (
  `userid` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '加密密码',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `familyid` int DEFAULT NULL COMMENT '所属家庭ID',
  PRIMARY KEY (`userid`),
  UNIQUE KEY `email` (`email`),
  KEY `FK_User_Family` (`familyid`),
  CONSTRAINT `FK_User_Family` FOREIGN KEY (`familyid`) REFERENCES `Family` (`familyid`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-09-03  8:39:12


CREATE TABLE Budget (
    familyid INT NOT NULL COMMENT '所属家庭编号',
    time DATE NOT NULL COMMENT '预算对应的时间（按月或按年）',
    amount DECIMAL(10,2) NOT NULL COMMENT '预算金额',
    PRIMARY KEY (familyid, time),
    CONSTRAINT FK_Budget_Family FOREIGN KEY (familyid) REFERENCES Family(familyid)
        ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='家庭预算表';
