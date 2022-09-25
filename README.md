# Introduction
This project is a template for using ```Gin``` + ```Wire``` + ```Gorm``` as a web framework. The project structure and logic layer  
is basically refer to SpringBoot project. The Swagger API Doc webpage also support in this project.  

These API is just basic user register and login functions. Also have secured path ```/content/*``` for you to test  
the Token(Bearer ...) verification with ```Gin``` middleware.

# Used Package
- [Gin](https://gin-gonic.com/)
    - The fastest full-featured web framework for Go.
    - Simple and comprehensive documentation.
- [Wire](https://github.com/google/wire)
    - Compile-time Dependency Injection Tool for Go.
    - Compare to the reflection based DI Toolkit([Dig](https://github.com/uber-go/dig)) from Uber.  
    Wire is more simple and easy for debugging.
- [GORM](https://gorm.io/)
    - The most common used ORM in Go community.
    - Chinese documentation supported.
- [Swaggo](https://github.com/swaggo/swag)
    - Auto generate API doc.
    - Good support for ```Gin```.
    - Chinese documentation supported.

# Project Structure

```
📦go-web-template
 ┣ 📂docs
 ┣ 📂modules
 ┃ ┣ 📂constant (Optional. Mostly put constant variables.)
 ┃ ┃ ┣ 📂exception
 ┃ ┃ ┗ 📂role
 ┃ ┣ 📂controller
 ┃ ┣ 📂dto
 ┃ ┣ 📂engine
 ┃ ┃ ┣ 📜provider.go (provider.go mainly provide instances for wiring in located module)
 ┃ ┣ 📂middleware (Provider middleware for ```Gin```. (e.g. Verify JWT))
 ┃ ┣ 📂model
 ┃ ┣ 📂repository (Create ```GORM``` instance)
 ┃ ┣ 📂service
 ┃ ┗ 📂util (Mostly some tools)
 ┃ ┃ ┣ 📂crypt
 ┃ ┃ ┗ 📂jwt
 ┣ 📜.gitignore
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜main.go
 ┣ 📜injector.go (Inject all dependencies into Gin)
 ┣ 📜wire_gen.go (Generated code from injector.go)
 ┗ 📜README.md
```

# Database is not setup?

Set the connection setting in ```mysql.go```.

```plaintext
{username}:{password}@tcp({host}:{port})/{schema}?charset=utf8&parseTime=True&loc=Local
```

The project database setting is 
```plaintext
username: root
password: eee333rr
host: 127.0.0.1
port: 3307
schema: web
```
You can change the setting with you favor.

## How about the table structure?

I intentionally not include the ```dump.sql``` in this project. Go check  
```mysql.go```. Open the comments when you first run the project.  
GORM will auto generate the table structure when it execute  
```gorm.AutoMigrate()```.

# Try API on Swagger

```localhost:8081/swagger/index.html```
