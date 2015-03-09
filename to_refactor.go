package gazette

import (
	"github.com/pippio/api-server/logging"
	"google.golang.org/cloud/storage"
)

func LastPersistedOffsetForJournal(context *logging.GCSContext,
	bucket, journalName string) (int64, error) {

	auth, err := context.ObtainAuthContext()
	if err != nil {
		return err
	}

	var query storage.Query
	query.Prefix = journalName
	storage.ListObjects(auth, bucket, query

	// storage.ListObjects with |key| prefix. Find and return entry
	// with max end offset.
	return 0
}
