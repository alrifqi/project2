# Dependencies
dependencies for this project is only [air](https://github.com/cosmtrek/air) for auto reload app in development. to install it just type

    go mod download
 
# Run In Dev Mode
to run in dev mode, type:

    make run-dev
 
 then open **localhost:8080** in browser to access application
 endpoint for get name is:
 - **localhost:8080/name**

# Build For Production
to run for production

    make build
 
 # Run Unit Test
 Currently unit test cover for controller, usecase & domain. to run all unit test type:
 

    make run-tests
