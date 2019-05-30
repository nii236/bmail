package main

func HandleRegistration(from, subject, body string) error {
	_, err := RegisterUser.Exec(map[string]interface{}{
		"username":      subject,
		"bitmessage_id": from,
	})
	return err
}
func HandleDeregistration(from, subject, body string) error {
	_, err := UnregisterUser.Exec(map[string]interface{}{
		"bitmessage_id": from,
	})
	return err
}
func HandleReceiving(from, subject, body string) error {
	return nil
}
func HandleBugReport(from, subject, body string) error {
	return nil
}
