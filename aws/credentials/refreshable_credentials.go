package credentials

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// RefreshableCredentialsProvider provides a name of RefreshableCredentials
// provider.
const RefreshableCredentialsProviderName = "RefreshableCredentialsProvider"

// RefreshableCredentials defines the format for refreshable credentials file.
type RefreshableCredentials struct {
	Expiration      time.Time
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string
	Token           string
}

// RefreshableCredentialsProvider is a credentials provider which can provide
// credentials from a file that is refreshed by an external source.
// The file is expected to be in the following json format:
//
// {
//   "Expiration": "",
//   "AccessKeyId": "",
//   "SecretAccessKey": "",
//   "Token": "",
// }
type RefreshableCredentialsProvider struct {
	Expiry
	Filename string
}

// NewSharedCredentials returns a new RefreshableCredentialsProvider.
func NewRefreshableCredentials(filename string) *Credentials {
	return NewCredentials(&RefreshableCredentialsProvider{
		Filename: filename,
	})
}

// Retrieve reads and extracts the refreshable credentials from file.
func (p *RefreshableCredentialsProvider) Retrieve() (Value, error) {
	data, err := ioutil.ReadFile(p.Filename)
	if err != nil {
		return Value{ProviderName: RefreshableCredentialsProviderName}, err
	}

	var credentials RefreshableCredentials
	err = json.Unmarshal(data, &credentials)
	if err != nil {
		return Value{ProviderName: RefreshableCredentialsProviderName}, err
	}

	p.SetExpiration(credentials.Expiration, 0)

	return Value{
		AccessKeyID:     credentials.AccessKeyID,
		SecretAccessKey: credentials.SecretAccessKey,
		SessionToken:    credentials.Token,
		ProviderName:    RefreshableCredentialsProviderName,
	}, nil
}
