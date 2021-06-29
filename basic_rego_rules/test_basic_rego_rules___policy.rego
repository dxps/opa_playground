package rules

# Anyone can GET /cars

test_car_read_negative {
	in = {
		"method": "GET",
		"path": ["nonexistent"],
		"user": "alice",
	}

	allow == false with input as in
}

test_car_read_positive {
	in = {
		"method": "GET",
		"path": ["cars"],
		"user": "alice",
	}

	allow == true with input as in
}

# Only managers can POST /cars

test___car_create___anyone {
	in = {
		"method": "POST",
		"path": ["cars"],
		"user": "unknown",
	}

	allow == false with input as in
}

test___car_create___employee {
	in = {
		"method": "POST",
		"path": ["cars"],
		"user": "alice",
	}

	allow == false with input as in
}

test___car_create___manager {
	in = {
		"method": "POST",
		"path": ["cars"],
		"user": "charlie",
	}

	allow == true with input as in
}

# Any managers can PUT /cars/{id}

test___car_update___anyone {
	in = {
		"method": "PUT",
		"path": ["cars", "1"],
		"user": "unknown",
	}

	allow == false with input as in
}

test___car_update___employee {
	in = {
		"method": "PUT",
		"path": ["cars", "1"],
		"user": "alice",
	}

	allow == false with input as in
}

test___car_update___manager {
	in = {
		"method": "PUT",
		"path": ["cars", "1"],
		"user": "charlie",
	}

	allow == true with input as in
}

# Any employee can GET /car/{id}/status

test___car_status___anyone {
	in = {
		"method": "GET",
		"path": ["cars", "1", "status"],
		"user": "unknown",
	}

	allow == false with input as in
}

test___car_status___employee {
	in = {
		"method": "GET",
		"path": ["cars", "1", "status"],
		"user": "bob",
	}

	allow == true with input as in
}

test___car_status___manager {
	in = {
		"method": "GET",
		"path": ["cars", "1", "status"],
		"user": "dave",
	}

	allow == true with input as in
}
