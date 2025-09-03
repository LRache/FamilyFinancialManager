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

