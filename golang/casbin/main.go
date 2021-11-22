package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

// domain
var domain = "DevOPS"

// user and role
// users
var dylan = "Dylan"
var mark = "Mark"
var jack = "Jack"
var daochun = "Daochun"
var gin = "Gin"
var betty = "Betty"
var david = "David"
var victor = "Victor"
var none = "None"

// roles
var roleAdmin = "Admin"
var roleDev = "Dev"
var roleOps = "Ops"

//
var cluster = "Cluster"

//var namespace = "Namespace"

var app = "App"

//var ci = "CI-Pipeline"
//var cd = "CD-Pipeline"

var add = "Add"
var update = "Update"
var delete = "Delete"
var list = "List"

func main() {
	modelFile := "model.conf"
	policyFile := "policy.conf"
	e, err := casbin.NewEnforcer(modelFile, policyFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//e = initPolicy(e)
	//e = initGroup(e)
	//e.SavePolicy()

	policies := e.GetPolicy()
	fmt.Println(policies)

	// get team members
	allUsers := e.GetAllUsersByDomain(domain)
	fmt.Println(allUsers)

	allUserRoles, _ := e.GetRolesForUser(dylan, domain)
	fmt.Println(allUserRoles)

	// get
	permissions := e.GetPermissionsForUserInDomain(dylan, domain)
	fmt.Println(permissions)

	permissions = e.GetPermissionsForUser(roleDev, domain)
	fmt.Println(permissions)

	//// policy
	if ok, err := e.Enforce(mark, cluster, add); err != nil {
		fmt.Println(err)
	} else {
		if ok {
			fmt.Println("success")
		} else {
			fmt.Println("fail")
		}
	}
}

func initPolicy(e *casbin.Enforcer) *casbin.Enforcer {

	rules := [][]string{
		{roleAdmin, domain, cluster, add},
		{roleAdmin, domain, cluster, update},
		{roleAdmin, domain, cluster, delete},
		{roleAdmin, domain, cluster, list},

		{roleAdmin, domain, app, add},
		{roleAdmin, domain, app, update},
		{roleAdmin, domain, app, delete},
		{roleAdmin, domain, app, list},

		{roleDev, domain, app, add},
		{roleDev, domain, app, update},
		{roleDev, domain, app, delete},
		{roleDev, domain, app, list},

		{roleOps, domain, cluster, add},
		{roleOps, domain, cluster, update},
		{roleOps, domain, cluster, delete},
		{roleOps, domain, cluster, list},

		{none, domain, cluster, list},
	}
	e.AddPolicies(rules)
	return e
}

func initGroup(e *casbin.Enforcer) *casbin.Enforcer {
	e.AddRolesForUser(dylan, []string{roleAdmin, roleDev, roleOps}, domain)

	e.AddRolesForUser(mark, []string{roleDev}, domain)
	e.AddRolesForUser(daochun, []string{roleDev}, domain)
	e.AddRolesForUser(gin, []string{roleDev}, domain)
	e.AddRolesForUser(victor, []string{roleDev}, domain)

	e.AddRolesForUser(jack, []string{roleOps}, domain)
	e.AddRolesForUser(david, []string{roleOps}, domain)
	e.AddRolesForUser(betty, []string{roleOps}, domain)

	return e
}
