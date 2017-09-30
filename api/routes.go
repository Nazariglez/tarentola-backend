// Created by nazarigonzalez on 30/9/17.

package api

//routes must be written in descending order because pat will use the first match
var routeList = []route{
	{GET, "/error", errorController},
	{GET, "/", homeController},
}
