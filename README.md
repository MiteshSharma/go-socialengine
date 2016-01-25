# go-socialengine (under development)
This project is to create a go web application which can handle all social related needs for any application and work as a service.
In this project we will handle like, comments and other social activities related to user object which is defined by some id and type provided by user.
We will try to make it generic so that this works for any kind of application and people don't need to replicate their efforts to write social framework for each and every application.

We are trying to make a MVC architecture with mysql as our default database.

## Feature in social engine:
1. object_id and object_type are two parameters which help identify any object uniquely.
2. As this is a service expected to run on backend of main server, we only added a small check based on email which help validate users on each request except first create user request.
3. Like objects: Users can like object and engine track all details regarding liking any object.
4. Commnet on objects: Users can comment on objects and engine keep track of all such comment write and give functionality to do a CRUD operation on them.
5. Channels: Engine has a concept of channel which is considered as a group of users subscribed to a channel and all of them are by default publisher. Users can store properties related to each channel. Users can share objects or messages with other users of same channel. This can be used as groups, alliance of any coordinated part where we want people to intract with each other to build a social ecosystem.

## Libraries used for developing this application :
- negroni
- go-sql-driver
- gorilla context
- gorm
- httprouter
- gin
