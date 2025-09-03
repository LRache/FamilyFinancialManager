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
