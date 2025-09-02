DELIMITER //

-- 新注册存储过程（指定 FamilyID）
CREATE PROCEDURE sp_register_api (
    IN p_username VARCHAR(50),
    IN p_password VARCHAR(255),
    OUT p_user_id INT
)
BEGIN
    DECLARE v_count INT;

    -- 检查用户名是否已存在
    SELECT COUNT(*) INTO v_count FROM Users WHERE UserName = p_username;
    IF v_count > 0 THEN
        SET p_user_id = NULL;
    ELSE
        -- 插入新用户
        INSERT INTO Users (UserName, Password)
        VALUES (p_username, p_password);
        SET p_user_id = LAST_INSERT_ID();
    END IF;
END//

DELIMITER ;

