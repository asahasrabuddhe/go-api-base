package mail

import (
	"github.com/asahasrabuddhe/pigeon"
	"github.com/asahasrabuddhe/pigeon/email"
	"github.com/asahasrabuddhe/pigeon/smtp"
	"github.com/asahasrabuddhe/pigeon/themes/flat"
	"github.com/spf13/viper"
)

var p pigeon.Pigeon
var d *smtp.Dialer

type Mail interface {
	Name() string
	Email() email.Email
}

func BootstrapMail() {
	p = pigeon.Pigeon{
		Product: pigeon.Product{
			Name: viper.GetString("app.name"),
			Link: viper.GetString("app.url"),
			Logo: viper.GetString("app.url") + ":" + viper.GetString("app.port") + "/public/images/logo.png",
		},
	}

	p.Theme = new(flat.Flat)

	d = smtp.NewDialer(viper.GetString("app.mail.smtp"), viper.GetInt("app.mail.port"), viper.GetString("app.mail.user"), viper.GetString("app.mail.password"))
}

func SetTheme(theme pigeon.Theme) {
	p.Theme = theme
}

func GenerateMail(mail Mail) (string, error) {
	return p.GenerateHTML(mail.Email())
}

func SendMail(message *smtp.Message) (bool, error) {
	if err := d.DialAndSend(message); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
