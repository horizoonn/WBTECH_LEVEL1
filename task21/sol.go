// Паттерн Adapter используется, когда необходимо конвертировать интерфейс класса в другой интерфейс, ожидаемый клиентом. Позволяет классам с разными интерфейсами работать вместе




package main

import "fmt"

type emailNotifier struct {}

func (e *emailNotifier) SendEmail(email string) {
	fmt.Println("Отправлено письмо на почту", email)
}

type pushNotifier struct {}

func (p *pushNotifier) SendPush(device string) {
	fmt.Println("Отправлено уведомление на устройство", device)
}

type Notifier interface {
	Notify(string)
}

type emailNotifierAdapter struct {
	*emailNotifier
}

func (adapter *emailNotifierAdapter) Notify(email string) {
	adapter.SendEmail(email)
}

func NewEmailNotifierAdapter(emailNotifier *emailNotifier) Notifier {
	return &emailNotifierAdapter{emailNotifier}
}

type pushNotifierAdapter struct {
	*pushNotifier
}

func (adapter *pushNotifierAdapter) Notify(device string) {
	adapter.SendPush(device)
}

func NewPushNotifierAdapter(pushNotifier *pushNotifier) Notifier {
	return &pushNotifierAdapter{pushNotifier}
}

type SMSNotifier struct {}

func (s *SMSNotifier) Notify(phoneNumber string) {
	fmt.Println("Отправлено смс на номер", phoneNumber)
}



func main() {
    
    emailAdapter := NewEmailNotifierAdapter(&emailNotifier{})
    pushAdapter := NewPushNotifierAdapter(&pushNotifier{})
    smsNotifier := &SMSNotifier{}
    
    notifiers := []Notifier{emailAdapter, pushAdapter, smsNotifier}
    
    for i, notifier := range notifiers {
        switch i {
        case 0:
            notifier.Notify("user@example.com")
        case 1:
            notifier.Notify("device123")
        case 2:
            notifier.Notify("+1234567890")
        }
	}
}
