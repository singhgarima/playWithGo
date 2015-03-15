# playWithGo

[![Build Status](https://travis-ci.org/singhgarima/playWithGo.svg?branch=master)](https://travis-ci.org/singhgarima/playWithGo)

## Running Web Server

* ``. ./setup.sh``
	This script sets up the GOPATH
* To run the application
	``go run webserver/src/main.go``

## Walk around the code-base

* ``main.go`` Contains of four main goroutines
	* To start http server
	* Data producer
	* Data consumer
	* Use data collected to populate the cache

* Packages
	* amqp_utilities: Some AMQP utility functions
	* consumer: Logic to consume data from AMQP queues
	* producer: Logic to publish data to AMQP queues
	* server: Http Server Routes and controller logic
	* lru_cache

## Continous Integration

* This is yet in starting stages as I am exploring unit testing in go-lang.
* Using travis-ci as CI tool
