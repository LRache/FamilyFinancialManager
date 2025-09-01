CREATE SCHEMA `jiating` ;
-- 1. 家庭表
CREATE TABLE Family (
    FamilyID INT AUTO_INCREMENT PRIMARY KEY COMMENT '家庭编号',
    FamilyName VARCHAR(100) NOT NULL COMMENT '家庭名称',
    Address VARCHAR(255) COMMENT '住址',
    Contact VARCHAR(100) COMMENT '联系方式'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='家庭信息表';


-- 2. 成员表
CREATE TABLE Member (
    MemberID INT AUTO_INCREMENT PRIMARY KEY COMMENT '成员编号',
    FamilyID INT NOT NULL COMMENT '所属家庭编号',
    Name VARCHAR(50) NOT NULL COMMENT '成员姓名',
    Gender ENUM('男','女','其他') COMMENT '性别',
    Age INT COMMENT '年龄',
    Relation VARCHAR(50) COMMENT '与家庭关系（如：父亲、母亲、子女等）',
    CONSTRAINT FK_Member_Family FOREIGN KEY (FamilyID) REFERENCES Family(FamilyID)
        ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='家庭成员表';


-- 3. 分类表（不考虑多级分类）
CREATE TABLE Category (
    CategoryID INT AUTO_INCREMENT PRIMARY KEY COMMENT '分类编号',
    CategoryName VARCHAR(100) NOT NULL COMMENT '分类名称（如：餐饮、购物、工资等）',
    Type ENUM('收入','支出') NOT NULL COMMENT '分类类型'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收支分类表';


-- 4. 收支记录表
CREATE TABLE TransactionRecord (
    TransactionID INT AUTO_INCREMENT PRIMARY KEY COMMENT '收支记录编号',
    FamilyID INT NOT NULL COMMENT '所属家庭编号',
    MemberID INT DEFAULT NULL COMMENT '成员编号，可为空表示整个家庭',
    CategoryID INT NOT NULL COMMENT '分类编号',
    Amount DECIMAL(10,2) NOT NULL COMMENT '金额',
    Date DATE NOT NULL COMMENT '收支日期',
    Type ENUM('收入','支出') NOT NULL COMMENT '收支类型',
    Merchant VARCHAR(100) COMMENT '商家名称（可选）',
    Location VARCHAR(100) COMMENT '消费地点（可选）',
    PaymentMethod VARCHAR(50) COMMENT '支付方式（现金/银行卡/微信/支付宝等）',
    Remark VARCHAR(255) COMMENT '备注',
    CONSTRAINT FK_Transaction_Family FOREIGN KEY (FamilyID) REFERENCES Family(FamilyID)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT FK_Transaction_Member FOREIGN KEY (MemberID) REFERENCES Member(MemberID)
        ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT FK_Transaction_Category FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID)
        ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='家庭收支记录表';
