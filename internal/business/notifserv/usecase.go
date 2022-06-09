package notifserv

type NotifCore struct{}

func NewNotifServ() *NotifCore { // return konkrit
	return &NotifCore{}
}

func (n *NotifCore) DummySendNotification(message string) error {
	// send notif to other server
	return nil
}
