# FitCal

Fit-cal is a comprehensive fitness mobile application that allows users to conveniently monitor their daily workouts and calorie intake. This application also features a recipe viewer that provides users with access to a wide range of recipes for various types of food. The primary objective behind this feature is to encourage users to prepare meals at home rather than dining out.

To achieve this, Fit-cal utilizes two external application programming interfaces (APIs) - [TheMealDB](https://www.themealdb.com/) for recipes and [EDAMAM](https://developer.edamam.com/food-database-api-docs#/) for nutritional information. These APIs work together to provide users with a comprehensive understanding of the food they consume, helping them make more informed decisions regarding their dietary habits.

In addition to its recipe and nutritional features, Fit-cal also offers a range of other functions that enhance the user experience.

Fit-Cal will also be providing a basic workout tracker feature.

## Used Tools
- [Golang](https://go.dev/) : Is used as the backend language.
- [Gorm](https://gorm.io/docs/index.html) : Is the selected ORM.
- [Mysql](https://www.mysql.com/jp/) : Is the selected DATABSE.
- [Echo](https://echo.labstack.com/) : Is the selected Goland framework.

# Pre-requisites
- [Docker](https://www.docker.com/)
- [Golang](https://go.dev/)
- [Go-migrate](https://github.com/golang-migrate/migrate)

## Cloning the repository

If you want a copy of the Backend(This repository) of the project you can clone the project with.

``` 
git clone git@github.com:Fit-cal/Fitcal-backend.git
```

## Creating a MySql database for development

To create a MySql database for development purposes use the following command.

 ```
 docker run -d \
 -p 3306:3306 \
 --name mysql \
 -e MYSQL_ROOT_PASSWORD=root \
 -e MYSQL_DATABASE=fitcal \
 -e MYSQL_USER=root \
 -e MYSQL_PASSWORD=root \
 mysql/mysql-server:latest
 ```

## Migrations

After the database has been created follow the steps on [Fitcal-Migrations](https://github.com/Fit-cal/Fitcal-Migrations) to get all the migrations and get the database ready.

## Running the project

Once the Database has been setup and all the migrations have been run now its time for us to finally start the project. 

Move to the directory of the project in your terminal and run the following command.

```
go run server.go
```
Once the project is setup try making a HTTP request to the following (or just copy and paste the line below on you url bar on your browser) to make sure if everything is working properly.

```
http://localhost:8080/
```
This should give a message that looks like following.

<img width="278" alt="image" src="https://user-images.githubusercontent.com/50660976/236676047-2cd3a8ea-2106-43be-bfee-a12c66dc7453.png">

## Running the project in debug mode

NOTE: This is strictly for `VScode` users only.

To start the project in `Debug-mode` first open the project on your `VScode` then navigate to `Run & Debug` on the side-bar. It's the button that looks like ![image](https://user-images.githubusercontent.com/50660976/236676081-3e7d68df-80b3-40d4-911b-1e68437fabd8.png) then press the play button on the top. There you have the project running on `Debug-mode`.

Or alternatively just press the `F5` key on your keyboard.
