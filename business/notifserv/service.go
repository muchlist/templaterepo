package notifserv

type NotifService struct{}

func NewNotifServ() *NotifService { // return konkrit
	return &NotifService{}
}

func (n *NotifService) DummySendNotification(message string) error {
	// send notif to other server
	return nil
}
