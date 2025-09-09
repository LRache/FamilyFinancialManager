package service

import (
	"backend/internal/repository"
	"backend/pkg/config"
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/wonderivan/logger"
)

func SendOverBudgetAlert(familyID int, amount float64, category string, currentSpent float64, budgetLimit float64) error {
	admin, err := repository.GetFamilyAdminByFamilyID(familyID)
	if err != nil {
		logger.Warn("Failed to get family admin:", err.Error())
		return err
	}

	if admin.Email == nil || *admin.Email == "" {
		logger.Warn("Family admin has no email address")
		return fmt.Errorf("家庭管理员未设置邮箱地址")
	}

	subject := "【家庭财务管理】预算超支提醒"
	body := fmt.Sprintf(`
尊敬的家庭管理员，

您的家庭在 "%s" 分类下的支出已超过预算：

本次支出金额：%.2f 元
当前总支出：%.2f 元
预算限额：%.2f 元
超支金额：%.2f 元

请及时关注家庭财务状况。

此致
家庭财务管理系统
`, category, amount/100.0, currentSpent/100.0, budgetLimit/100.0, (currentSpent-budgetLimit)/100.0)

	// 发送邮件
	err = sendEmail(*admin.Email, subject, body)
	if err != nil {
		logger.Warn("Failed to send email:", err.Error())
		return err
	}

	logger.Info("Over budget alert email sent to:", admin.Email)
	return nil
}

// sendEmail 发送邮件的内部函数
func sendEmail(to, subject, body string) error {
	// 检查邮件服务是否启用
	if !config.Email.Enabled {
		logger.Info("Email service is disabled, logging email instead:")
		logger.Info("To:", to)
		logger.Info("Subject:", subject)
		logger.Info("Body:", body)
		return nil
	}

	// 使用配置文件中的设置
	emailConfig := config.Email

	// 构造邮件
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body))

	// 连接SMTP服务器
	auth := smtp.PlainAuth("", emailConfig.Username, emailConfig.Password, emailConfig.SMTPHost)
	addr := emailConfig.SMTPHost + ":" + strconv.Itoa(emailConfig.SMTPPort)

	err := smtp.SendMail(addr, auth, emailConfig.FromEmail, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
