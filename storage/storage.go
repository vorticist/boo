package storage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"gitlab.com/vorticist/logger"
	"io"
	"os"
)

const (
	bucketName = "vortex"
	storeName  = "book-of-omens"
)

var (
	storeUrl string
)

type Storer interface {
	AddEntry(key, value string)
	RemoveEntry(key string)
	GetEntries() map[string]string
	GetToken() string
}

func NewStorer() (Storer, error) {
	godotenv.Load(".env")
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("SPACE_ACCESS_KEY"), os.Getenv("SPACE_SECRET"), ""),
		Endpoint:         aws.String(fmt.Sprintf("https://%v", os.Getenv("SPACE_URL"))),
		Region:           aws.String("fra1"),
		S3ForcePathStyle: aws.Bool(false),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		logger.Errorf("failed create session: %v", err)
		return nil, err
	}
	s3Client := s3.New(newSession)

	return &storer{
		client: s3Client,
	}, nil
}

type storer struct {
	client *s3.S3
}

func (s *storer) AddEntry(key, value string) {
	storeMap := s.getStoreMap()
	storeMap[key] = value
	err := s.saveStoreMap(storeMap)
	if err != nil {
		logger.Errorf("failed to save changes to store: %v", err)
	}
}

func (s *storer) RemoveEntry(key string) {
	storeMap := s.getStoreMap()
	delete(storeMap, key)

	err := s.saveStoreMap(storeMap)
	if err != nil {
		logger.Errorf("failed to save changes to store: %v", err)
	}
}

func (s *storer) GetEntries() map[string]string {
	return s.getStoreMap()
}

func (s *storer) readStoreData() []byte {

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(storeName),
	}

	result, err := s.client.GetObject(input)
	if err != nil {
		logger.Errorf("failed to download object: %v", err)
		fmt.Println(err.Error())
	}

	data, err := io.ReadAll(result.Body)
	if err != nil {
		logger.Errorf("failed to read data: %v", err)
		return nil
	}

	return data
}

func (s *storer) GetToken() string {
	//obj, err := s.client.GetObject(context.Background(), bucketName, "book-token", minio.GetObjectOptions{})
	//if err != nil {
	//	logger.Errorf("error getting object from space: %v", err)
	//	return ""
	//}
	//
	//data, err := io.ReadAll(obj)
	//if err != nil {
	//	logger.Errorf("error reading file from space: %v", err)
	//	return ""
	//}
	//
	//return string(data)
	return ""
}

func (s *storer) saveStoreMap(m map[string]string) error {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(m)
	if err != nil {
		logger.Errorf("error encoding map: %v", err)
		return err
	}

	newData := b.Bytes()

	object := s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(storeName),
		Body:   bytes.NewReader(newData),
		ACL:    aws.String("private"),
		Metadata: map[string]*string{
			"x-amz-meta-my-key": aws.String(os.Getenv("SPACE_ACCESS_KEY")),
		},
	}

	_, err = s.client.PutObject(&object)
	if err != nil {
		logger.Errorf("error putting object: %v", err)
		return err
	}

	logger.Info("file modified successfully")
	return nil
}

func (s *storer) getStoreMap() map[string]string {
	data := s.readStoreData()
	logger.Infof("got data: %v", string(data))
	var decodedMap map[string]string

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&decodedMap)
	if err != nil {
		if err == io.EOF {
			logger.Warn("sotre empty, initializing...")
			return map[string]string{}
		}
		logger.Errorf("failed to decode data: %v", err)
		return nil
	}

	return decodedMap
}
