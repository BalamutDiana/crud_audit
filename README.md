# Event logging app for crud_movie_manager
 ### Tools
 - go 1.19
 - MongoDB
 - RabbitMQ

 ### How to use this:
 Run container with MongoDB:
 ```cmd
 docker run --rm -d --name audit-log-mongo -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=qwerty -p 27017:27017 mongo:latest
 ```
 Run application and it will save events from https://github.com/BalamutDiana/crud_movie_manager to database.
 
