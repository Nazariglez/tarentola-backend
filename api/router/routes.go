// Created by nazarigonzalez on 30/9/17.

package router

import "github.com/nazariglez/tarentola-backend/api/controllers"

//routes must be written in descending order because pat will use the first match
var routeList = []route{
	{POST, "/user", controllers.CreateUser, "isAdmin"},
	{GET, "/", controllers.NotFound, "isAdmin"},
}
