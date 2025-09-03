-- 用户注册
DELIMITER //

CREATE PROCEDURE sp_register_api (
    IN p_username VARCHAR(50),
    IN p_password VARCHAR(255),
    OUT p_userid INT
)
BEGIN
    DECLARE v_count INT;

    -- 检查用户名是否已存在
    SELECT COUNT(*) INTO v_count 
    FROM Users 
    WHERE username = p_username;

    IF v_count > 0 THEN
        SET p_userid = NULL;
    ELSE
        -- 插入新用户（未指定familyid）
        INSERT INTO Users (username, password)
        VALUES (p_username, p_password);

        SET p_userid = LAST_INSERT_ID();
    END IF;
END//

DELIMITER ;

-- 创建家庭
DELIMITER //

CREATE PROCEDURE sp_create_family (
    IN p_userid INT,               -- 创建家庭的用户ID
    IN p_familyname VARCHAR(100),  -- 家庭名称
    OUT p_familyid INT             -- 输出新家庭ID
)
BEGIN
    DECLARE v_count INT;

    -- 1. 检查该用户是否已加入某个家庭
    SELECT COUNT(*) INTO v_count
    FROM Users
    WHERE userid = p_userid AND familyid IS NOT NULL;

    IF v_count > 0 THEN
        -- 如果已属于某个家庭，则抛出错误
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = '该用户已属于某个家庭，无法再次创建家庭';
    ELSE
        -- 2. 插入新家庭
        INSERT INTO Family (familyname)
        VALUES (p_familyname);

        -- 3. 获取新插入的家庭ID
        SET p_familyid = LAST_INSERT_ID();

        -- 4. 更新用户为该家庭管理员
        UPDATE Users
        SET familyid = p_familyid,
            role = 1
        WHERE userid = p_userid;
    END IF;
END //

DELIMITER ;

-- 邀请成员
DELIMITER //

CREATE PROCEDURE sp_invite_user_to_family (
    IN p_inviter_id INT,   -- 邀请人（必须是管理员）
    IN p_invitee_id INT    -- 被邀请人
)
BEGIN
    DECLARE v_familyid INT;
    DECLARE v_role TINYINT;
    DECLARE v_count INT;

    -- 1. 检查邀请人是否存在并且是管理员
    SELECT familyid, role INTO v_familyid, v_role
    FROM Users
    WHERE userid = p_inviter_id;

    IF v_familyid IS NULL OR v_role <> 1 THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = '只有家庭管理员才能邀请用户加入家庭';
    END IF;

    -- 2. 检查被邀请人是否已经有家庭
    SELECT COUNT(*) INTO v_count
    FROM Users
    WHERE userid = p_invitee_id AND familyid IS NOT NULL;

    IF v_count > 0 THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = '该用户已属于某个家庭，无法被邀请';
    END IF;

    -- 3. 更新被邀请用户的家庭和角色（成员）
    UPDATE Users
    SET familyid = v_familyid,
        role = 0
    WHERE userid = p_invitee_id;
END //

DELIMITER ;

-- 设置预算
DELIMITER //

CREATE PROCEDURE sp_set_family_budget (
    IN p_userid INT,           -- 操作用户ID（必须是管理员）
    IN p_time DATE,            -- 预算对应的时间（按月或按年）
    IN p_amount DECIMAL(10,2)  -- 预算金额
)
BEGIN
    DECLARE v_familyid INT;
    DECLARE v_role TINYINT;
    DECLARE v_count INT;

    -- 1. 检查用户是否是管理员
    SELECT familyid, role INTO v_familyid, v_role
    FROM Users
    WHERE userid = p_userid;

    IF v_familyid IS NULL OR v_role <> 1 THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = '只有家庭管理员才能设置家庭预算';
    END IF;

    -- 2. 检查是否已有该时间的预算记录
    SELECT COUNT(*) INTO v_count
    FROM Budget
    WHERE familyid = v_familyid AND time = p_time;

    -- 3. 如果存在则更新，否则插入
    IF v_count > 0 THEN
        UPDATE Budget
        SET amount = p_amount
        WHERE familyid = v_familyid AND time = p_time;
    ELSE
        INSERT INTO Budget (familyid, time, amount)
        VALUES (v_familyid, p_time, p_amount);
    END IF;
END //

DELIMITER ;

DELIMITER $$
-- 根据日期查询收支
CREATE PROCEDURE GetFamilyFinanceByUser(
    IN p_userid INT,
    IN p_start_date DATE,
    IN p_end_date DATE
)
BEGIN
    DECLARE v_familyid INT;

    -- 找到用户所属家庭
    SELECT familyid INTO v_familyid
    FROM Users
    WHERE userid = p_userid;

    -- 如果用户没有家庭，直接返回空
    IF v_familyid IS NULL THEN
        SELECT '该用户没有所属家庭' AS message;
    ELSE
        -- 查询该家庭在时间范围内的收支情况
        SELECT 
            f.familyid,
            f.familyname,
            c.type AS income_or_expense, -- 1=收入, 0=支出
            SUM(t.amount) AS total_amount
        FROM TransactionRecord t
        JOIN Category c ON t.categoryid = c.categoryid
        JOIN Family f ON t.familyid = f.familyid
        WHERE t.familyid = v_familyid
          AND DATE(t.occurred_at) BETWEEN p_start_date AND p_end_date
        GROUP BY f.familyid, f.familyname, c.type;
    END IF;
END $$

DELIMITER ;


DELIMITER $$
-- 家庭收入/支出的录入
CREATE PROCEDURE AddTransactionRecord(
    IN p_userid INT,
    IN p_categoryid INT,
    IN p_amount DECIMAL(10,2),
    IN p_occurred_at DATETIME,
    IN p_note VARCHAR(255),
    IN p_merchant VARCHAR(100),
    IN p_location VARCHAR(100),
    IN p_paymentmethod VARCHAR(50)
)
BEGIN
    DECLARE v_familyid INT;
    DECLARE v_category_type TINYINT;

    -- 检查分类是否存在，并获取类型
    SELECT type INTO v_category_type
    FROM Category
    WHERE categoryid = p_categoryid;

    IF v_category_type IS NULL THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '分类不存在';
    END IF;

    -- 如果指定用户，则获取其所属家庭
    IF p_userid IS NOT NULL THEN
        SELECT familyid INTO v_familyid
        FROM Users
        WHERE userid = p_userid;

        IF v_familyid IS NULL THEN
            SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '用户不存在或未关联家庭';
        END IF;
    ELSE
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '必须指定用户ID或手动设置家庭ID';
    END IF;

    -- 插入记录
    INSERT INTO TransactionRecord(familyid, userid, categoryid, amount, occurred_at, note, merchant, location, paymentmethod)
    VALUES (v_familyid, p_userid, p_categoryid, p_amount, p_occurred_at, p_note, p_merchant, p_location, p_paymentmethod);
END $$

DELIMITER ;


DELIMITER $$
-- 修改已有的收支记录
CREATE PROCEDURE EditTransactionRecord(
    IN p_transactionrecordid INT,
    IN p_userid INT,
    IN p_categoryid INT,
    IN p_amount DECIMAL(10,2),
    IN p_occurred_at DATETIME,
    IN p_note VARCHAR(255),
    IN p_merchant VARCHAR(100),
    IN p_location VARCHAR(100),
    IN p_paymentmethod VARCHAR(50)
)
BEGIN
    DECLARE v_familyid INT;
    DECLARE v_category_type TINYINT;

    -- 检查账单是否存在，并获取所属家庭
    SELECT familyid INTO v_familyid
    FROM TransactionRecord
    WHERE transactionrecordid = p_transactionrecordid;

    IF v_familyid IS NULL THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '账单不存在';
    END IF;

    -- 如果修改用户ID，则验证用户是否属于同一家庭
    IF p_userid IS NOT NULL THEN
        IF NOT EXISTS (
            SELECT 1 FROM Users 
            WHERE userid = p_userid AND familyid = v_familyid
        ) THEN
            SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '用户不存在或不属于同一家庭';
        END IF;
    END IF;

    -- 如果修改分类，则验证分类是否存在
    IF p_categoryid IS NOT NULL THEN
        SELECT type INTO v_category_type
        FROM Category
        WHERE categoryid = p_categoryid;

        IF v_category_type IS NULL THEN
            SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '分类不存在';
        END IF;
    END IF;

    -- 执行更新
    UPDATE TransactionRecord
    SET 
        userid = IFNULL(p_userid, userid),
        categoryid = IFNULL(p_categoryid, categoryid),
        amount = IFNULL(p_amount, amount),
        occurred_at = IFNULL(p_occurred_at, occurred_at),
        note = IFNULL(p_note, note),
        merchant = IFNULL(p_merchant, merchant),
        location = IFNULL(p_location, location),
        paymentmethod = IFNULL(p_paymentmethod, paymentmethod)
    WHERE transactionrecordid = p_transactionrecordid;

END $$

DELIMITER ;


DELIMITER $$
-- 删除已有的收支记录
CREATE PROCEDURE DeleteTransactionRecord(
    IN p_transactionrecordid INT
)
BEGIN
    DECLARE v_exists INT;

    -- 检查账单是否存在
    SELECT COUNT(*) INTO v_exists
    FROM TransactionRecord
    WHERE transactionrecordid = p_transactionrecordid;

    IF v_exists = 0 THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '账单不存在';
    END IF;

    -- 执行删除
    DELETE FROM TransactionRecord
    WHERE transactionrecordid = p_transactionrecordid;

END $$

DELIMITER ;
