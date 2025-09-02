DELIMITER //

-- 1) 用户注册：返回 uuid 和 username
CREATE PROCEDURE sp_register_api (
    IN p_username VARCHAR(50),
    IN p_password VARCHAR(255),
    OUT p_uuid CHAR(36),
    OUT p_user_name VARCHAR(50)
)
BEGIN
    DECLARE v_count INT;

    -- 检查用户名是否已存在
    SELECT COUNT(*) INTO v_count FROM Users WHERE UserName = p_username;
    IF v_count > 0 THEN
        SET p_uuid = NULL;
        SET p_user_name = NULL;
    ELSE
        -- 插入新用户
        SET p_uuid = UUID();
        INSERT INTO Users (UUID, UserName, Password) VALUES (p_uuid, p_username, p_password);
        SET p_user_name = p_username;
    END IF;
END//
  
-- 2) 用户登录：返回 uuid 和 token
CREATE PROCEDURE sp_login_api (
    IN p_username VARCHAR(50),
    IN p_password VARCHAR(255),
    OUT p_uuid CHAR(36),
    OUT p_token VARCHAR(255)
)
BEGIN
    DECLARE v_count INT;

    -- 检查用户是否存在
    SELECT COUNT(*) INTO v_count
    FROM Users
    WHERE UserName = p_username AND Password = p_password;

    IF v_count = 0 THEN
        SET p_uuid = NULL;
        SET p_token = NULL;
    ELSE
        -- 取 UUID
        SELECT UUID INTO p_uuid
        FROM Users
        WHERE UserName = p_username AND Password = p_password
        LIMIT 1;

        -- 生成 token 并保存
        SET p_token = UUID();
        UPDATE Users SET Token = p_token WHERE UUID = p_uuid;
    END IF;
END//

-- 3) 自动登录：根据 token 获取 uuid 和 username
CREATE PROCEDURE sp_autologin_api (
    IN p_token VARCHAR(255),
    OUT p_uuid CHAR(36),
    OUT p_user_name VARCHAR(50)
)
BEGIN
    DECLARE v_count INT;

    SELECT COUNT(*) INTO v_count
    FROM Users
    WHERE Token = p_token;

    IF v_count = 0 THEN
        SET p_uuid = NULL;
        SET p_user_name = NULL;
    ELSE
        SELECT UUID, UserName INTO p_uuid, p_user_name
        FROM Users
        WHERE Token = p_token
        LIMIT 1;
    END IF;
END//

DELIMITER ;
