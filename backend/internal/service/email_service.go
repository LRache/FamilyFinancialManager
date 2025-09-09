package service

import (
	"backend/internal/repository"
	"backend/pkg/config"
	"crypto/tls"
	"fmt"

	"github.com/wonderivan/logger"
	"gopkg.in/gomail.v2"
)

func SendOverBudgetAlert(familyID int, amount float64, category string) error {
	logger.Info("Preparing to send over budget alert email...")

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

请及时关注家庭财务状况。

此致
家庭财务管理系统
`, category, amount)

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

	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.Account)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(emailConfig.SMTPHost, emailConfig.SMTPPort, emailConfig.Account, emailConfig.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// gomail v2 默认支持 587 端口的 STARTTLS，无需额外配置
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
