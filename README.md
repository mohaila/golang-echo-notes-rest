# Golang REST Server using Echo framework
The project is organized as:
- integration: put all your integartion tests here
- migration: put all your SQL here
- model: put all your business models here
- server: for each business model, add a modelhandlers file to handle REST 
- service: put your business logic here
- store: for each business model, add a file for database code
- test: put all your unit tests here
- vendor: when using govendor to manage external GO modules, your are free to use another tool

I'm using gorc to run tests.

## TODO
- Add unit tests
- Add missing integration tests (the store integration test is included)
- Add CI/CD scripts for Gitlab and Docker
