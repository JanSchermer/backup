package client

import (
	"backup/util"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Client struct {
	folder        string
	config        string
	iterator      util.Iterator
	uploader      util.Uploader
	waitGroup     sync.WaitGroup
	uploaderCount int
	writeChan     chan string
	errorChan     chan error
	dbs           map[string]util.Database
	aes           util.AES
}

func GetClient(folder string, config string) Client {
	data := util.GetDatabase(config+"/config.txt", ":")
	encKey := data.Get("Encryption Key")
	auth := data.Get("Authentication Key")
	url := data.Get("URL")
	invalid := false
	if encKey == "" || encKey == "YOUR KEY HERE" {
		data.Set("Encryption Key", "YOUR KEY HERE")
		invalid = true
	}
	if auth == "" || auth == "YOUR KEY HERE" {
		data.Set("Authentication Key", "YOUR KEY HERE")
		invalid = true
	}
	if url == "" || url == "YOUR URL HERE" {
		data.Set("URL", "YOUR URL HERE")
		invalid = true
	}
	if invalid {
		data.WriteData()
		panic("Please edit the config file")
	}
	encKey = util.HashString(encKey)
	aes := util.GetAES(encKey)
	iterator := util.GetIterator(folder)
	uploader := util.GetUploader(url, auth, aes)
	writeChan := make(chan string, 100)
	errorChan := make(chan error)
	dbs := make(map[string]util.Database)
	return Client{folder, config, iterator, uploader, sync.WaitGroup{}, 0, writeChan, errorChan, dbs, aes}
}

func (c *Client) Start() {
	for c.iterator.HasNext() {
		select {
		case err := <-c.errorChan:
			if err == nil {
				return
			}
			panic(err)
		default:
		}
		select {
		case file := <-c.writeChan:
			c.ReciveResults(file, c.GetDb(file))
		default:
			if c.uploaderCount > 5 {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			file := c.iterator.Next()
			c.StartUploader(file, c.GetDb(file))
		}
	}
	c.waitGroup.Wait()
}

func (c *Client) GetDb(file string) util.Database {
	dir := filepath.Dir(file)
	dir = c.config + "/" + dir
	os.MkdirAll(dir, 0777)
	file = dir + ".db"
	if _, ok := c.dbs[file]; !ok {
		c.dbs[file] = util.GetDatabase(file, "$")
	}
	if len(c.dbs) > 50 {
		for k := range c.dbs {
			c.dbs[k].Close()
			delete(c.dbs, k)
		}
	}
	return c.dbs[file]
}

func (c *Client) ReciveResults(file string, db util.Database) {
	hash := <-c.writeChan
	db.Set(file, hash)
	db.WriteData()
}

func (c *Client) StartUploader(file string, db util.Database) {
	if file == "" {
		return
	}
	hash := util.HashFile(file)
	if hash == "" {
		return
	}
	if db.Get(file) != hash {
		c.waitGroup.Add(1)
		c.uploaderCount++
		go c.UploadFile(file, hash, c.errorChan)
	}
}

func (c *Client) UploadFile(file string, hash string, errorChan chan error) {
	c.uploader.Upload(file, hash, errorChan)
	c.writeChan <- file
	c.writeChan <- hash
	c.uploaderCount--
	c.waitGroup.Done()
}
