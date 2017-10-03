// Created by nazarigonzalez on 30/9/17.

package router

import "github.com/nazariglez/tarentola-backend/api/controllers"

//routes must be written in descending order because pat will use the first match
var routeList = []route{
	{POST, "/login", controllers.Login, ""},
	{POST, "/token", controllers.TestToken, "isLogged"},
	{POST, "/user", controllers.CreateUser, ""},

	{GET, "/", controllers.NotFound, ""},
}
