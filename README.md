# go-socialengine
This project is to create a go web application which can handle all social related needs for any application and work as a service.
In this project we will handle like, comments and other social activities related to user object which is defined by some id and type provided by user.
We will try to make it generic so that this works for any kind of application and people don't need to replicate their efforts to write social framework for each and every application.

We are trying to make a MVC architecture with mysql as our default database.

## Feature in social engine:
1. We identify each object uniquely by its id and type (object_id and object_type).
2. Validate user based on its email on each call except user create call.
3. Each object can be liked and we save user id of each user who like a object.
4. Each object can be commented by any user. We give functionality to read, update, delete, create comments.

## Libraries used for developing this application :
- negroni
- go-sql-driver
- gorilla context
- gorm
- httprouter
- gin
