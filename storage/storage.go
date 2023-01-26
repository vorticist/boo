package storage

import (
	"bytes"
	"encoding/gob"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go"
	"gitlab.com/vorticist/logger"
	"io"
	"os"
)

const (
	bucketName = "vortex"
	objName    = "book-of-omens"
)

type Storer interface {
	AddEntry(key, value string)
	RemoveEntry(key string)
	GetEntries() map[string]string
	GetToken() string
}

func NewStorer() (Storer, error) {
	godotenv.Load(".env")
	// Set up a client to the DigitalOcean Space
	client, err := minio.New(os.Getenv("SPACE_URL"), os.Getenv("SPACE_ACCESS_KEY"), os.Getenv("SPACE_SECRET"), true)
	if err != nil {
		logger.Errorf("could not start client: %v", err)
		return nil, err
	}
	return &storer{
		client: client,
	}, nil
}

type storer struct {
	client *minio.Client
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
	// Read the file contents
	obj, err := s.client.GetObject(bucketName, objName, minio.GetObjectOptions{})
	if err != nil {
		logger.Errorf("error getting object from space: %v", err)
		return nil
	}
	data, err := io.ReadAll(obj)
	if err != nil {
		logger.Errorf("error reading file from space: %v", err)
		return nil
	}

	return data
}

func (s *storer) GetToken() string {
	obj, err := s.client.GetObject(bucketName, "book-token", minio.GetObjectOptions{})
	if err != nil {
		logger.Errorf("error getting object from space: %v", err)
		return ""
	}

	data, err := io.ReadAll(obj)
	if err != nil {
		logger.Errorf("error reading file from space: %v", err)
		return ""
	}

	return string(data)
}

func (s *storer) saveStoreMap(m map[string]string) error {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(m)
	if err != nil {
		logger.Errorf("error encoding map:", err)
		return err
	}

	newData := b.Bytes()
	_, err = s.client.PutObject(bucketName, "filename", bytes.NewReader(newData), int64(len(newData)), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		logger.Errorf("error saving file to space:", err)
		return err
	}
	logger.Info("file modified successfully")
	return nil
}

func (s *storer) getStoreMap() map[string]string {
	data := s.readStoreData()
	var decodedMap map[string]string

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&decodedMap)
	if err != nil {
		logger.Errorf("error decoding byte slice:", err)
		return nil
	}

	return decodedMap
}
