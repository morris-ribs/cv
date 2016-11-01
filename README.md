# My CV app server side REST in Go
A RESTful API in Go to create, delete and search for data in MongoDB

# Pre-requisites
- Install MongoDB: https://docs.mongodb.com/manual/installation/?jmp=footer
- Install Go: https://golang.org/dl/
- Be sure you have an app to make API calls (Swagger, Postman, Fiddler, ...)

# To run the code
Open your prompt command and follow the instructions below:

```
# `cd` to project directory
$ cd <project_dir>

# get all dependencies
$ go get ./...

# launch server
$ go run main.go

# Go to http://localhost:8080 on your browser
```

# First: launch mongo daemon!
After installing your MongoDB instance, open a cmd window and run the following command:
```
# launch mmongo daemon
$ mongod
```

Your MongoDB instance is up and ready to treat your requests!

# Second: insert dummy data
Open your API client and send a POST to http://localhost:8080/candidate and insert the content of candidate_example.json in the body of the request.

# Third: get the data you inserted
Open your API client and send a GET to http://localhost:8080/candidate and check the data retrieved from the database (the one you just inserted above).

Optionally, you can copy the id field of the retrieved candidate and call http://localhost:8080/candidate/{id} to fetch its data.

# (Optional) Last but not least: delete the retrieved user
To do it, just send a DELETE request to http://localhost:8080/candidate/{id} in order to delete the inserted user.

# ! Important remark !
This project will suffer changes, especially in its organization. Don't hesitate in dropping a comment!
