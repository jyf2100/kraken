package originbackend

import (
	"errors"
	"io"

	"code.uber.internal/infra/kraken/core"
	"code.uber.internal/infra/kraken/origin/blobclient"
)

// Client is a write-only backend client which uploads blobs to origin cluster.
type Client struct {
	config  Config
	cluster blobclient.ClusterClient
}

// NewClient creates a new Client.
func NewClient(config Config) (*Client, error) {
	if config.Namespace == "" {
		return nil, errors.New("no namespace configured")
	}
	cluster := blobclient.NewClusterClient(
		blobclient.NewClientResolver(blobclient.NewProvider(), config.Addr))
	return newClient(config, cluster), nil
}

func newClient(config Config, cluster blobclient.ClusterClient) *Client {
	return &Client{config, cluster}
}

// Download downloads name into dst. name must be the sha256 digest of src.
func (c *Client) Download(name string, dst io.Writer) error {
	d := core.NewSHA256DigestFromHex(name)
	return c.cluster.DownloadBlob(c.config.Namespace, d, dst)
}

// Upload uploads src to name. name must be the sha256 digest of src.
func (c *Client) Upload(name string, src io.Reader) error {
	d := core.NewSHA256DigestFromHex(name)
	through := !c.config.DisableUploadThrough
	return c.cluster.UploadBlob(c.config.Namespace, d, src, through)
}
