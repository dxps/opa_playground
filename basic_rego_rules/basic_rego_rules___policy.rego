package rules

# Some Data
users := {
	"alice": {"manager": "charlie", "title": "salesperson"},
	"bob": {"manager": "charlie", "title": "salesperson"},
	"charlie": {"manager": "dave", "title": "manager"},
	"dave": {"manager": null, "title": "ceo"},
}

default allow = false

allow {
	# Anyone can read cars.
	input.method == "GET"
	input.path == ["cars"]
}

allow {
	# Only managers can create a new car.
	user_is_manager
	input.method == "POST"
	input.path == ["cars"]
}

allow {
	# Only managers can update a car.
	user_is_manager
	input.method == "PUT"
	count(input.path) == 2
	input.path[0] == "cars"
}

allow {
	# Only employees can GET /cars/{carID}
	user_is_employee
	input.method == "GET"
	input.path = ["cars", carid]
}

allow {
	# Only employees can GET /cars/{carID}/status
	user_is_employee
	input.method == "GET"
	input.path = ["cars", carid, "status"]
}

# Helpers

user_is_manager {
	users[input.user].title != "salesperson"
}

default user_is_employee = false

user_is_employee {
	users[input.user]
}
