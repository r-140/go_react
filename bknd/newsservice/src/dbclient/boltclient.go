package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"model"

	"github.com/boltdb/bolt"
)

// IBoltClient ...
type IBoltClient interface {
	OpenBoltDb()
	QueryNews(newsID string) (model.News, error)
	QueryAllNews() ([]model.News, error)
	Seed()
	Check() bool
}

//BoltClient  Real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

// OpenBoltDb ...
func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Check Naive healthcheck, just makes sure the DB connection has been initialized.
func (bc *BoltClient) Check() bool {

	return bc.boltDB != nil

}

//QueryAllNews ...
func (bc *BoltClient) QueryAllNews() ([]model.News, error) {
	// Allocate an empty News instance we'll let json.Unmarhal populate for us in a bit.
	// news := model.News{}

	newsSlice := make([]model.News, 0)

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("NewsBucket"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			news := model.News{}
			json.Unmarshal(v, &news)
			newsSlice = append(newsSlice, news)

		}

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return newsSlice, err
	}
	// Return the News struct and nil as error.
	return newsSlice, nil
}

// QueryNews ...
func (bc *BoltClient) QueryNews(newsID string) (model.News, error) {
	// Allocate an empty News instance we'll let json.Unmarhal populate for us in a bit.
	news := model.News{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("NewsBucket"))

		// Read the value identified by our newsId supplied as []byte
		newsBytes := b.Get([]byte(newsID))
		if newsBytes == nil {
			return fmt.Errorf("No news found for " + newsID)
		}
		// Unmarshal the returned bytes into the news struct we created at
		// the top of the function
		json.Unmarshal(newsBytes, &news)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.News{}, err
	}
	// Return the News struct and nil as error.
	return news, nil
}

// Start seeding news
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedNews()
}

// Creates an "NewsBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("NewsBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed (n) make-believe account objects into the NewsBucket bucket.
func (bc *BoltClient) seedNews() {

	total := 10
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our News struct
		news := model.News{
			Id:     key,
			Title:  "Nachrichte_" + strconv.Itoa(i),
			Teaser: "Teaser_" + strconv.Itoa(i),
			Body:   "This is body of nachrichte_" + strconv.Itoa(i),
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(news)

		// Write the data to the NewsBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("NewsBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake news...\n", total)
}
