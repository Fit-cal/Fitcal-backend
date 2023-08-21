# FitCal

Fit-cal is a comprehensive fitness mobile application that allows users to conveniently monitor their daily workouts and calorie intake. The name "Fit-Cal" comes from fitness and calculation combined because this application does just that it calculates your daily nutrition value and calorie intake to help you track your fitness effectively.

This application also features a recipe viewer that provides users with access to a wide range of recipes for various types of food. The primary objective behind this feature is to encourage users to prepare meals at home rather than dining out.

To achieve this, Fit-cal utilizes two external application programming interfaces (APIs) - [TheMealDB](https://www.themealdb.com/) for recipes and [EDAMAM](https://developer.edamam.com/food-database-api-docs#/) for nutritional information. These APIs work together to provide users with a comprehensive understanding of the food they consume, helping them make more informed decisions regarding their dietary habits.

In addition to its recipe and nutritional features, Fit-cal also offers a range of other functions that enhance the user experience.

Fit-Cal also provides it's user's a basic workout tracking feature.

## Contents
- [Used Tools](#used-tools)
- [Pre-Requisites](#pre-requisites)
- [Clone the repository](#cloning-the-repository)
- [Mysql](#mysql)
- [Migrations](#migrations)
- [Running the project](#running-the-project)
- [Debug mode](#debug-mode)

## Used Tools
- [Golang](https://go.dev/) : Backend language.
- [Gorm](https://gorm.io/docs/index.html) : ORM.
- [Mysql](https://www.mysql.com/jp/) : DATABSE.
- [Echo](https://echo.labstack.com/) : Golang framework.
- [Zerologger](https://github.com/rs/zerolog) : Logger

## Pre-requisites
- [Docker](https://www.docker.com/)
- [Golang](https://go.dev/)
- [Go-migrate](https://github.com/golang-migrate/migrate)

## Cloning the repository

If you want a copy of the Backend(This repository) of the project you can clone the project with.

``` 
git clone https://github.com/Fit-cal/Fitcal-backend.git
```

## Mysql 

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

> **Note:** *If you are planning to use docker then you can just run `docker-compose up` on the root directory.*

```
go run server.go
```
Once the project is setup try making a HTTP request to the following (or just copy and paste the line below on you url bar on your browser) to make sure if everything is working properly.

```
http://localhost:8080/
```
This should give a message that looks like following.

<div 
    style="
    width:70%;
    display:flex;
    align-items:center;
    justify-content:center;
    padding: 20px;
    margin: 20px auto;
    border-radius: 5px;
    "
    id="debug-mode"
>
<img width="278" alt="image" src="https://user-images.githubusercontent.com/50660976/236676047-2cd3a8ea-2106-43be-bfee-a12c66dc7453.png">
</div>

## Debug mode

NOTE: This is strictly for `VScode` users only.

To start the project in `Debug-mode` first open the project on your `VScode` then navigate to `Run & Debug` on the side-bar. It's the button that looks like ![image](https://user-images.githubusercontent.com/50660976/236676081-3e7d68df-80b3-40d4-911b-1e68437fabd8.png) then press the play button on the top. There you have the project running on `Debug-mode`.

Or alternatively just press the `F5` key on your keyboard.
