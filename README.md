# Sample REST API server

## About

This solution demonstrates Golang HTTP server capable of accept and store specific JSON structures. It also can return stored values.
Solution uses simple in-memory storage.

Challenges:
* some variability possible in provided JSONs.

Code highlights:
* gorilla/mux as HTTP server router (internal/api/server/server.go)
* application-specific handler format (see 'AppHandler' type and 'toHandler' func).
* custom JSON marshalling (see 'RawTemperature' and 'RawTime' types)
* end-to-end HTTP tests (see internal/api/server/server_test.go)
* in-memory storage is designed to be replaceable (internal/storage/repo.go)


Assumptions made:

* Application can accept data in different formats. But outputs only the single format for each type
    * for time this is "2006-01-02T15:04:05.999Z07:00"
    * for temperature - float with Celsius value
* if timezone is unknown, assume GMT0
* invalid json in the task (no comma after "HDDSpace"): assume json is valid
* assume that InternalTemp is allowed to have 'c' too (no such case in demo data)

## Instructions

Go 1.18 is required to be installed locally and added to the PATH.

Use makefile's shortcuts to build, test and run the project.
* "make build" : builds the project and places the binary result "sample_server" to the project root
* "make test"  : runs all unit tests
* "make run"   : builds and executes the binary with debug configuration

Application can be configured with certain environment variables. To get the list of them - run "make help_env" or simply "./sample_server --help". E.g. to run app with custom port run "SAMPLE_SERVER_SERVER_ADDR=:3999 ./sample_server"

Once application executed, it's api is accessible. Default port is 4000.


## API documentation

Api contains two endpoints

### Add entry

Curl example:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
    "machineId": 12345,
    "stats": {
        "cpuTemp": 90.1,
        "fanSpeed": 400,
        "HDDSpace": 800,
        "internalTemp": "72c"
    },
    "lastLoggedIn": "admin/Paul",
    "sysTime": "2022-04-23T18:25:43.511Z"
}' \
  http://localhost:4000/
```

Fields cpuTemp and internalTemp could be both 
- number (integer or float) or 
- string containing number with letter 'c' in the end

Allowed time formats:
* "2006-01-02T15:04:05Z07:00"
* "Mon 2006-01-02 15:04:05" (UTC assume timezone)

## Read all entries

```bash
curl 127.0.0.1:4000 
```

It returns an json array of objects. Object structure is close to previous api, but has additional ID field.
This json uses only one sysTime format and always use numeric temperatures.

```json
[
    {
        "id": "632ce4f2-4bd7-4ea5-94ca-5e6c9c06eed2",
        "machineId": 12345,
        "stats": {
            "cpuTemp": 90.1,
            "fanSpeed": 400,
            "HDDSpace": 800,
            "internalTemp": 72
        },
        "lastLoggedIn": "admin/Paul",
        "sysTime": "2022-04-23T18:25:43.511Z"
    }
]
```