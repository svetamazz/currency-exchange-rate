## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)

## General info
This project is a small API created for the [Genesis Software Engineering School 3.0](https://www.genesis-for-univ.com/genesis-software-engineering-school-3). It's written using [Golang](https://go.dev/). It uses [Fiber](https://gofiber.io/) as a web framework, and [Brevo(formerly Sendinblue)](https://www.brevo.com/) for sending emails. This project is dockerized (it includes Dockerfile, docker-compose.yml, and Makefile).

⚠️ I can't expose my Brevo API key for sending emails in the repository, but I can share it with you if needed (hit me up at swetamazkiw@gmail.com). You may create your own API key following [this doc from Brevo](https://help.brevo.com/hc/en-us/articles/209467485-Create-and-manage-your-API-keys#h_01sdGW6ZQEKZ072SFGK03N9R6VE6).
	
## Technologies
Project is created with:
* Golang 1.20.4
* Fiber
* Brevo (formerly Sendinblue)
	
## Setup
Create .env file from the example.env (and set env variables in newly created file):
```
$ cp example.env .env
```
Run following command to build and start the web server inside docker:
```
$ make up
```
---
Run following command to stop running web server inside docker:
```
$ make down
```