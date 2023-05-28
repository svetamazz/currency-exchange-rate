package helpers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
	. "github.com/svetamazz/currency-exchange-rate/types"
)

const DB_PATH = "./database/emails.json"

func GetEmails() (*EmailList, error) {
	file, err := ioutil.ReadFile(DB_PATH)
	if err != nil {
		if os.IsNotExist(err) {
			return &EmailList{}, nil
		}
		return nil, err
	}

	var emails EmailList
	err = json.Unmarshal(file, &emails)
	if err != nil {
		return nil, err
	}

	return &emails, nil
}

func SaveEmails(emails *EmailList) error {
	file, err := json.MarshalIndent(emails, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(DB_PATH, file, 0644)
	if err != nil {
		return err
	}

	return nil
}


func SendEmails(emails []string, rate float64) error {
    loadEnvVariables()
    apiKey := os.Getenv("SENDINBLUE_API_KEY")

	cfg := sendinblue.NewConfiguration()
	cfg.AddDefaultHeader("api-key", apiKey)
	cfg.AddDefaultHeader("partner-key",apiKey)

	sib := sendinblue.NewAPIClient(cfg)

	var recipients []sendinblue.SendSmtpEmailTo
	for _, email := range emails {
		recipients = append(recipients, sendinblue.SendSmtpEmailTo{ Email: email })
	}

    sender := &sendinblue.SendSmtpEmailSender{
        Name:  "Svitlana Matskiv",
        Email: "swetamazkiw@gmail.com",
    }
    sendEmail := sendinblue.SendSmtpEmail{
		Sender: sender,
		To: recipients,
        TemplateId: 1,
        Params: map[string]interface{}{
            "rate": rate,
        },
	}

    var ctx context.Context
    _, _, err := sib.TransactionalEmailsApi.SendTransacEmail(ctx, sendEmail)

    if err != nil {
        return err
    }
	return nil
}

func loadEnvVariables() error {
    err := godotenv.Load()
    if err != nil {
        return err
    }
    return nil
}

func Contains(emails *EmailList, email string) bool {
	for _, e := range emails.Emails {
		if e == email {
			return true
		}
	}
	return false
}