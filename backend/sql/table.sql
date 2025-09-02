-- MySQL dump 10.13  Distrib 8.0.43, for macos15 (arm64)
--
-- Host: localhost    Database: family_finance
-- ------------------------------------------------------
-- Server version	9.4.0

-- create database family;
-- use family;


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
-- Table structure for table `Budget`
--

DROP TABLE IF EXISTS `Budget`;
CREATE TABLE `Budget` (
  `BudgetID` int NOT NULL AUTO_INCREMENT COMMENT '预算ID',
  `FamilyID` int NOT NULL COMMENT '所属家庭ID',
  `CategoryID` int DEFAULT NULL COMMENT '对应分类，可为空表示整体预算',
  `Amount` decimal(10,2) NOT NULL COMMENT '预算金额',
  `StartDate` date NOT NULL COMMENT '预算开始日期',
  `EndDate` date DEFAULT NULL COMMENT '预算结束日期',
  `Note` varchar(255) DEFAULT NULL COMMENT '备注',
  `CreatedAt` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `UpdatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`BudgetID`),
  KEY `FK_Budget_Family` (`FamilyID`),
  KEY `FK_Budget_Category` (`CategoryID`),
  CONSTRAINT `FK_Budget_Category` FOREIGN KEY (`CategoryID`) REFERENCES `Category` (`CategoryID`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `FK_Budget_Family` FOREIGN KEY (`FamilyID`) REFERENCES `Family` (`FamilyID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='家庭预算表';

--
-- Table structure for table `Category`
--

DROP TABLE IF EXISTS `Category`;
CREATE TABLE `Category` (
  `CategoryID` int NOT NULL AUTO_INCREMENT COMMENT '分类编号',
  `CategoryName` varchar(100) NOT NULL COMMENT '分类名称',
  `Type` tinyint(1) NOT NULL COMMENT '收支类型，1=收入，0=支出',
  `CreatedAt` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`CategoryID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='收支分类表';

--
-- Table structure for table `Family`
--

DROP TABLE IF EXISTS `Family`;
CREATE TABLE `Family` (
  `FamilyID` int NOT NULL AUTO_INCREMENT COMMENT '家庭编号',
  `FamilyName` varchar(100) NOT NULL COMMENT '家庭名称',
  `Address` varchar(255) DEFAULT NULL COMMENT '住址',
  `Contact` varchar(100) DEFAULT NULL COMMENT '联系方式',
  `CreatedAt` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`FamilyID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='家庭信息表';

--
-- Table structure for table `Member`
--

DROP TABLE IF EXISTS `Member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Member` (
  `MemberID` int NOT NULL AUTO_INCREMENT COMMENT '成员编号',
  `FamilyID` int NOT NULL COMMENT '所属家庭编号',
  `UserID` int DEFAULT NULL COMMENT '关联用户ID，可为空表示未注册成员',
  `Name` varchar(50) NOT NULL COMMENT '成员姓名',
  `Email` varchar(100) DEFAULT NULL COMMENT '邮箱，用于邀请',
  `Status` enum('pending','accepted') DEFAULT 'accepted' COMMENT '邀请状态',
  `Gender` enum('男','女','其他') DEFAULT NULL COMMENT '性别',
  `Age` int DEFAULT NULL COMMENT '年龄',
  `Relation` varchar(50) DEFAULT NULL COMMENT '与家庭关系（父亲、母亲、子女等）',
  PRIMARY KEY (`MemberID`),
  KEY `FK_Member_Family` (`FamilyID`),
  KEY `FK_Member_User` (`UserID`),
  CONSTRAINT `FK_Member_Family` FOREIGN KEY (`FamilyID`) REFERENCES `Family` (`FamilyID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_Member_User` FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='家庭成员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `RecurringTransaction`
--

DROP TABLE IF EXISTS `RecurringTransaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `RecurringTransaction` (
  `RecurringID` int NOT NULL AUTO_INCREMENT COMMENT '定期账单ID',
  `FamilyID` int NOT NULL COMMENT '所属家庭ID',
  `MemberID` int DEFAULT NULL COMMENT '成员ID',
  `CategoryID` int NOT NULL COMMENT '分类ID',
  `Type` tinyint(1) NOT NULL COMMENT '收支类型，1=收入，0=支出',
  `Amount` decimal(10,2) NOT NULL COMMENT '账单金额',
  `OccurredAt` datetime NOT NULL COMMENT '第一次发生时间',
  `IntervalType` enum('daily','weekly','monthly') NOT NULL COMMENT '周期类型',
  `Note` varchar(255) DEFAULT NULL COMMENT '备注',
  `CreatedAt` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `UpdatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`RecurringID`),
  KEY `FK_Recurring_Family` (`FamilyID`),
  KEY `FK_Recurring_Member` (`MemberID`),
  KEY `FK_Recurring_Category` (`CategoryID`),
  CONSTRAINT `FK_Recurring_Category` FOREIGN KEY (`CategoryID`) REFERENCES `Category` (`CategoryID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_Recurring_Family` FOREIGN KEY (`FamilyID`) REFERENCES `Family` (`FamilyID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_Recurring_Member` FOREIGN KEY (`MemberID`) REFERENCES `Member` (`MemberID`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='定期账单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `TransactionRecord`
--

DROP TABLE IF EXISTS `TransactionRecord`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `TransactionRecord` (
  `TransactionID` int NOT NULL AUTO_INCREMENT COMMENT '收支记录编号',
  `FamilyID` int NOT NULL COMMENT '所属家庭编号',
  `MemberID` int DEFAULT NULL COMMENT '成员编号，可为空表示家庭整体',
  `CategoryID` int NOT NULL COMMENT '分类编号',
  `Type` tinyint(1) NOT NULL COMMENT '收支类型，1=收入，0=支出',
  `Amount` decimal(10,2) NOT NULL COMMENT '金额',
  `OccurredAt` datetime NOT NULL COMMENT '发生时间',
  `Note` varchar(255) DEFAULT NULL COMMENT '备注',
  `Merchant` varchar(100) DEFAULT NULL COMMENT '商家名称（可选）',
  `Location` varchar(100) DEFAULT NULL COMMENT '消费地点（可选）',
  `PaymentMethod` varchar(50) DEFAULT NULL COMMENT '支付方式（现金/银行卡/微信/支付宝等）',
  `CreatedAt` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `UpdatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`TransactionID`),
  KEY `FK_Transaction_Family` (`FamilyID`),
  KEY `FK_Transaction_Member` (`MemberID`),
  KEY `FK_Transaction_Category` (`CategoryID`),
  CONSTRAINT `FK_Transaction_Category` FOREIGN KEY (`CategoryID`) REFERENCES `Category` (`CategoryID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_Transaction_Family` FOREIGN KEY (`FamilyID`) REFERENCES `Family` (`FamilyID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_Transaction_Member` FOREIGN KEY (`MemberID`) REFERENCES `Member` (`MemberID`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='家庭收支记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `Users`
--

DROP TABLE IF EXISTS `Users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Users` (
  `UserID` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `UserName` varchar(50) NOT NULL COMMENT '用户名',
  `Password` varchar(255) NOT NULL COMMENT '加密密码',
  `Email` varchar(100) DEFAULT NULL COMMENT '用户邮箱',
  `FamilyID` int DEFAULT NULL COMMENT '所属家庭ID',
  `CreatedAt` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `UpdatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`UserID`),
  UNIQUE KEY `UserName` (`UserName`),
  UNIQUE KEY `Email` (`Email`),
  KEY `FK_User_Family` (`FamilyID`),
  CONSTRAINT `FK_User_Family` FOREIGN KEY (`FamilyID`) REFERENCES `Family` (`FamilyID`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-09-02  9:27:39
