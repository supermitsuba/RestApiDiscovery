Restful Api Discovery Service
=============================
Purpose of this service is to allow you to query for RESTful APIs.

This uses GO and has Unit tests.  Things missing is quering on detail vs summary and error check when putting and posting.  Im sure all add that in later with some unit tests, but this is a great foundation to start.

Message Format
==============
Here is the model that is being serialized into json:
```
type RestApiDescription struct {
	Url         string `json:"url"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Location    string `json:"location"`
	IsActive    bool   `json:"active"`
	Id          string `json:"id"`
}
```

Basically, you should expect to POST and PUT using this structure.  Also, if you are GETting data, you should see a json array of a similar structure here, too.

API Description
===============
For a more detailed idea of what this API does, there is a raml file included in the code base.  You can use : http://ramlexample.cloudapp.net/ to paste the raml file (/documentation/RestApiDiscovery.raml) into it and see the API visually.

Testing
=======

To run the unit test, check out this command:
```
/usr/local/go/bin/go test -v [$GOPATH/src/RestApiDiscovery]
```

To do some code coverage commands, use these:
```
/usr/local/go/bin/go test -coverprofile cover.out  [$GOPATH/src/RestApiDiscovery]

/usr/local/go/bin/go tool cover -func=cover.out  [$GOPATH/src/RestApiDiscovery]
```

Running
=======

When running the app, it will open on port 8088.  To change this, go to main.go and you should see the listen line where 8088 is specified.  Future note, should probably be a configuration or something.

Example of building a bin:
```
/usr/local/go/bin/go build -i [$GOPATH/src/RestApiDiscovery]
```

Dependencies
============
If you are wondering what Dependencies there are, there is good news, I have a dockerfile to help with that.  While I am still playing with docker, I have not made an official image yet, but this docker file will allow you to isolate what Dependencies you will need.  Look for the file in this directory: /documentation/Dockerfile