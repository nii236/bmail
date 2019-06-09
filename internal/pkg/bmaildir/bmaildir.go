package bmaildir

import (
	"github.com/luksen/maildir"
)

// Open will open or create a maildir with a basepath and given username
func Open(maildirPath string) (maildir.Dir, error) {
	dir := maildir.Dir(maildirPath)
	err := dir.Create()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// ReadNew will check for new messages, mark them as read and return the list of keys for each message
func ReadNew(dir maildir.Dir) ([]string, error) {
	return dir.Unseen()
}
