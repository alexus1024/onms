# onms


Assumptions made:

* Application can accept data in different formats. But outputs only the single format for each type
    * for time this is "2006-01-02T15:04:05.999Z07:00"
    * for temperature - float with Celsius value
* if timezone is unknown, assume GMT0
* invalid json in the task (no comma after "HDDSpace"): assume json is valid
* assume that InternalTemp is allowed to have 'c' too (no such case in demo data)