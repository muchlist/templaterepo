package notifserv

type NotifService struct{}

func NewNotifServ() *NotifService { // return konkrit
	return &NotifService{}
}

func (n *NotifService) SendNotification(message string) error {
	// TODO : send notif to other server
	return nil
}
