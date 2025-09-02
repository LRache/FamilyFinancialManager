-- 家庭收支汇总视图
CREATE VIEW v_family_income_expense AS
SELECT 
    f.FamilyID,
    f.FamilyName,
    YEAR(t.OccurredAt) AS year,
    MONTH(t.OccurredAt) AS month,
    SUM(CASE WHEN t.Type=1 THEN t.Amount ELSE 0 END) AS total_income,
    SUM(CASE WHEN t.Type=0 THEN t.Amount ELSE 0 END) AS total_expense,
    SUM(CASE WHEN t.Type=1 THEN t.Amount ELSE 0 END) -
    SUM(CASE WHEN t.Type=0 THEN t.Amount ELSE 0 END) AS net_income
FROM Family f
LEFT JOIN TransactionRecord t ON f.FamilyID = t.FamilyID
GROUP BY f.FamilyID, f.FamilyName, YEAR(t.OccurredAt), MONTH(t.OccurredAt);

-- 分类收支汇总视图
CREATE VIEW v_category_summary AS
SELECT 
    f.FamilyID,
    c.CategoryID,
    c.CategoryName,
    c.Type AS category_type,
    YEAR(t.OccurredAt) AS year,
    MONTH(t.OccurredAt) AS month,
    SUM(t.Amount) AS total_amount
FROM TransactionRecord t
JOIN Category c ON t.CategoryID = c.CategoryID
JOIN Family f ON t.FamilyID = f.FamilyID
GROUP BY f.FamilyID, c.CategoryID, YEAR(t.OccurredAt), MONTH(t.OccurredAt);

-- 成员收支汇总视图
CREATE VIEW v_member_summary AS
SELECT 
    m.MemberID,
    m.Name AS member_name,
    m.FamilyID,
    YEAR(t.OccurredAt) AS year,
    MONTH(t.OccurredAt) AS month,
    SUM(CASE WHEN t.Type=1 THEN t.Amount ELSE 0 END) AS income,
    SUM(CASE WHEN t.Type=0 THEN t.Amount ELSE 0 END) AS expense
FROM Member m
LEFT JOIN TransactionRecord t ON m.MemberID = t.MemberID
GROUP BY m.MemberID, m.Name, m.FamilyID, YEAR(t.OccurredAt), MONTH(t.OccurredAt);


