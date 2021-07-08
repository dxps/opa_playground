package products_access.dashboard

# By default, respond negative.
default subject_has_access_to_product = false

## -------------------------------------------------------------------------
## If the `input` provided product is part of subject's portfolio, then yes.
## Or if the subject has the support role in this context.
## TODO: to be extended with logic for non-support cases.
## -------------------------------------------------------------------------
subject_has_access_to_product {
	subject_has_SelfServiceSupportRole
}

## --------------------------------------------------------------
## If subject has `role` of `SelfServiceSupportRole` then yes.
## This info is fetched based on input.subjectToken.
## --------------------------------------------------------------
subject_has_SelfServiceSupportRole {
	subjectID := token.payload.sub
	trace(sprintf("subjectID=%v", [subjectID]))
	url := sprintf("http://localhost:3001/v1/subjects/%s/attributes", [subjectID])
	subjectAttrs := http.send({
		"url": url,
		"method": "GET",
		"headers": {"Authorization": sprintf("Bearer %s", [input.subjectToken])},
		"raise_error": false,
	}).body

	trace(sprintf("subjectAttrs=%v", [subjectAttrs])) # This debug message becomes a Note event in the query explanation.

	attr := subjectAttrs[_]
	attr.name == "role"
	attr.value == "SelfServiceSupportRole"
}

## Helper to get the subject id from input's subjectToken.
token = {"payload": payload} {
	[header, payload, signature] := io.jwt.decode(input.subjectToken)
}

## This is one option to get the details, that is extracting as a different rule.
# subject_attrs = resp.body {
# 	subjectID := token.payload.sub
# 	trace(sprintf("subjectID=%v", [subjectID]))
# 	url := sprintf("http://localhost:3001/v1/subjects/%s/attributes", [subjectID])
# 	resp := http.send({
# 		"url": url,
# 		"method": "GET",
# 		"headers": {"Authorization": sprintf("Bearer %s", [input.subjectToken])},
# 		"raise_error": false,
# 	})
# }
