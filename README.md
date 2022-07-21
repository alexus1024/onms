# onms


Assumptions made:

* if timezone is unknown, assume GMT0
* invalid json in the task (no comma after "HDDSpace"): assume json is valid
* assume that InternalTemp is allowed to have 'c' too (no such case in demo data)
* could imagine that temperature could contain not only 'c' but also 'f'. Discussable.