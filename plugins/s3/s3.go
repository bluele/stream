package s3

import (
	"github.com/bluele/stream"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Plugin struct {
	s3 *s3.S3
}

func (pl *Plugin) Init() error {
	auth, err := aws.EnvAuth()
	if err != nil {
		return err
	}
	pl.s3 = s3.New(auth, aws.APNortheast)
	return nil
}

func (pl *Plugin) Name() string {
	return "s3"
}

func (pl *Plugin) FileReader(path string) (*stream.File, error) {
	bucketName, prefix := splitPath(path)
	ir, err := pl.s3.Bucket(bucketName).GetReader(prefix)
	if err != nil {
		return nil, err
	}
	return stream.NewFile(filepath.Base(prefix), ir), nil
}

func (pl *Plugin) DirReader(path string) (*stream.Dir, error) {
	return nil, nil
}

func (pl *Plugin) WriteFile(fi *stream.File, path string) error {
	bucketName, prefix := splitPath(path)
	data, err := ioutil.ReadAll(fi)
	if err != nil {
		return err
	}
	return pl.s3.Bucket(bucketName).Put(prefix, data, "application/octet-stream", s3.Private)
}

func (pl *Plugin) WriteDir(di *stream.Dir, path string) error {
	return nil
}

func (pl *Plugin) List(path string, iw io.Writer) error {
	if bucketName, path := splitPath(path); bucketName == "" {
		resp, err := pl.s3.ListBuckets()
		if err != nil {
			return err
		}
		for _, bucket := range resp.Buckets {
			iw.Write([]byte(bucket.Name + "\n"))
		}
	} else {
		bucket := pl.s3.Bucket(bucketName)
		resp, err := bucket.List(path, "/", "", 1000)
		if err != nil {
			return err
		}
		for _, key := range resp.Contents {
			iw.Write([]byte(key.Key + "\n"))
		}
	}

	return nil
}

func splitPath(path string) (string, string) {
	cols := strings.Split(path, "/")
	return cols[0], strings.Join(cols[1:], "/")
}
