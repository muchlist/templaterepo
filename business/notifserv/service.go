package notifserv

type NotifService struct{}

// return konkrit struct, bukan interfacenya
func NewNotifServ() *NotifService {
	return &NotifService{}
}

// SendNotification diperlukan untuk memenuhi interface NotifSender pada service user
func (n *NotifService) SendNotification(message string) error {
	// TODO : send notif to other server
	return nil
}

// SendWhatsapp tidak diperlukan oleh service user namun bisa jadi diperlukan oleh service lain
func (n *NotifService) SendWhatsapp(message string, phone string) error {
	// TODO : send whatsapp
	return nil
}
