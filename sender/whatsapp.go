package sender

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Whatsapp() error {
	time.Sleep(2 * time.Second)

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("whatsapp:" + os.Getenv("MY_PHONE_NUMBER"))
	params.SetFrom("whatsapp:" + os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("PeerBR🟣: Nova Nota Fiscal Criada💰!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	response, _ := json.Marshal(*resp)
	fmt.Println("Message sent successfully: " + string(response))

	return nil
}
