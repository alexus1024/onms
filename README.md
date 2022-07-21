# Sample REST API server

# About

This solution demonstrates Golang HTTP server capable of accept and store specific JSON structures. It also can return stored values.
Solution uses simple in-memory storage.

Challenges:
* some variability possible in provided JSONs.

Code highlights:
* gorilla/mux as HTTP server router (internal/api/server/server.go)
* application-specific handler format (see 'AppHandler' type and 'toHandler' func).
* custom JSON marshalling (see 'RawTemperature' and 'RawTime' types)
* end-to-end HTTP rests (see internal/api/server/server_test.go)
* in-memory storage is designed to be replaceable (internal/storage/repo.go)


Assumptions made:

* Application can accept data in different formats. But outputs only the single format for each type
    * for time this is "2006-01-02T15:04:05.999Z07:00"
    * for temperature - float with Celsius value
* if timezone is unknown, assume GMT0
* invalid json in the task (no comma after "HDDSpace"): assume json is valid
* assume that InternalTemp is allowed to have 'c' too (no such case in demo data)

# Instructions

Use makefile's shortcuts to build, test and run the project.

