<h3 align="center">Home Backup</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> 
    Home Backup is a simple application, that allowes you to backup important data to your own server. The data will be encrypted locally on your computer, before it is send to the server. To this point only AES is suppoted as encryption algorithm.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Server Usage](#usage_server)
- [Client Usage](#usage_client)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

This project has been created, to allow you to backup your data, without compromising your privacy.

## 🏁 Getting Started <a name = "getting_started"></a>

First install go on your machine. (This is an example for Arch. Golang 1.16 or later is required)
```
sudo pacman -Sy golang 
```

Then clone the repository or download it as a zip file and extract it.
```
git clone https://github.com/JanSchermer/backup
```

Now you're ready to build the project and create an executable.
```
go build
```

Now your should have an executable name "backup" ready for use.
## 🎈 Usage As Server <a name="usage_server"></a>

In order to start a backup sever, you first need to create a few files within the working directory. A TSL certificate and private key will be required. If you use self signed certificates, please remember to trust them on your machine. Put the certificate into the working dir and name it "fullchain.pem". Do the same with the private key and name it "privkey.pem". Futhuremore you need to provide authentication details, that the client will use to authenticate with the server. Create a auths.txt file and fill it with details in the following format:
```
AUTHENTICATION_KEY:user
QWERTZUIOPASDFGHJKLYXCVBNM:jan
```

Now you're ready to lauch the server. You can do this in the following way:
```
backup server <PORT>
```

The server should start up now and be ready to go.
## 🎈 Usage As Client <a name="usage_client"></a>

In order to start a backup client, some configuration will be required. You will need to provide: An encryption key, an authentication key and a url of your backup server. Please create a config.txt in your data folder and fill it in the following way:
```
Encryption Key:MY_ENCRYPTION_PASSWORD
Authentication Key:AUTHENTICATION_KEY_IN_auths.txt_ON_THE_SERVER
URL:URL_TO_THE_BACKUP_SERVER
```

You may choose any password as an encryption key. The authentication key has to be the same as the one specified in auths.txt on the server. The URL should point to your server. For a localhost server running on port 9999, this would be: "https://127.0.0.1:9999/"

Now you may launch the client with the following options:
```
backup client <upload_folder> <data_folder>
```

upload_folder should be set to the folder, that you like to backup.
data_folder should be set to the configuration folder. It will contain your configuration as well as data needed and created by the backup software.

## ✍️ Authors <a name = "authors"></a>

- [@JanSchermer](https://github.com/JanSchermer) - Idea & Initial work

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to Google for providing us with GoLang <3