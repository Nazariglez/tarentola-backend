// Created by nazarigonzalez on 30/9/17.

package router

import "github.com/nazariglez/tarentola-backend/api/controllers"

//routes must be written in descending order because pat will use the first match
var routeList = []route{
	{POST, "/token", controllers.TestToken, "isLogged"},

	{POST, "/login", controllers.Login, "isNotLogged"},

	{PUT, "/user/{id}", controllers.UpdateUser, "isLogged"},
	{DELETE, "/user/{id}", controllers.DeleteUser, "isLogged"},
	{POST, "/user", controllers.CreateUser, "isNotLogged"},

	{GET, "/", controllers.NotFound, ""},
}
