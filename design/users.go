package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("users", func() {
	Title("User Service")
	Description("Service for manipulate user data")
	Server("users", func() {
		Description("user hosts the User Service.")

		// List the services hosted by this server.
		Services("users", "roles", "openapi")

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})

		Host("production", func() {
			Description("Production hosts.")
			// URIs can be parameterized using {param} notation.
			URI("https://{version}.goa.design")
			URI("grpcs://{version}.goa.design")

			// Variable describes a URI variable.
			Variable("version", String, "API version", func() {
				// URL parameters must have a default value and/or an
				// enum validation.
				Default("v1")
			})
		})
	})
})

// StoredRole result type
var StoredRole = ResultType("application/vnd.stored-role", func() {
	Description("A StoredRole describes a role retrieved by the users service.")
	Reference(Role)
	TypeName("StoredRole")

	Attributes(func() {
		Attribute("name", String, "Name of role", func() {
			Example("admin")
			Pattern(`[a-z]+[a-z0-9]*`)
			Meta("rpc:tag", "1")
		})
		Attribute("description", String, "Description of role", func() {
			Example("Administrator")
			Meta("rpc:tag", "2")
		})
	})
	View("default", func() {
		Attribute("name")
		Attribute("description")
	})
	View("tiny", func() {
		Attribute("name")
	})
	Required("name")
})

// Role type
var Role = Type("Role", func() {
	Description("Role describes a role to be stored.")
	Field(1, "name", String, "Name of role", func() {
		Example("admin")
		Pattern(`[a-z]+[a-z0-9]*`)
	})
	Field(2, "description", String, "Description of role", func() {
		Example("Administrator")
	})
	Required("name")
})

// StoredUser is the result of user data
var StoredUser = ResultType("application/vnd.stored-user", func() {
	Description("A StoredUser describes a user retrieved by the users service.")
	Reference(User)
	TypeName("StoredUser")

	Attributes(func() {
		Field(1, "email")
		Field(2, "firstname")
		Field(3, "lastname")
		Field(4, "isactive")
		Field(5, "role")
	})

	View("default", func() {
		Attribute("email")
		Attribute("role", StoredRole, func() {
			View("tiny")
		})
		Attribute("firstname")
		Attribute("lastname")
		Attribute("isactive")
	})

	View("tiny", func() {
		Attribute("email")
		Attribute("role", StoredRole, func() {
			View("tiny")
		})
		Attribute("isactive")
	})

	Required("email", "firstname", "lastname", "role")
})

// User type
var User = Type("User", func() {
	Description("User describes a user to be stored.")
	Field(1, "email", String, "Email of the user", func() {
		Pattern(`.+@.+\..{1,6}`)
		Example("ehabterra@hotmail.com")
	})
	Field(2, "firstname", String, "First Name of the user", func() {
		MaxLength(100)
		Example("Ehab")
	})
	Field(3, "lastname", String, "Last Name of user", func() {
		MaxLength(100)
		Example("Terra")
	})
	Field(4, "role", String, "user role", func() {
		Example("admin")
		Pattern(`[a-z]+[a-z0-9]*`)
	})
	Field(5, "isactive", Boolean, "Is user active.", func() {
		Default(true)
	})
	Required("email", "firstname", "lastname", "role")
})

// NotFound type
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a user that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("user 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})

var _ = Service("users", func() {
	Description("The users service performs user data.")

	HTTP(func() {
		Path("/users")
	})

	Method("list", func() {
		Description("List all stored users")
		Payload(func() {
			Field(1, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
		})
		Result(CollectionOf(StoredUser))
		HTTP(func() {
			GET("/")
			Param("view")
			Response(StatusOK)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show user by Email")
		Payload(func() {
			Field(1, "email", String, "Email of user to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("email")
		})
		Result(StoredUser)
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			GET("/{email}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("add", func() {
		Description("Add new user and return email.")
		Payload(User)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("update", func() {
		Description("Update existing user and return email.")
		Payload(User)
		Result(String)
		HTTP(func() {
			PUT("/{email}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove user from users data")
		Payload(func() {
			Field(1, "email", String, "Email of user to remove")
			Required("email")
		})
		Error("not_found", NotFound, "Email not found")
		HTTP(func() {
			DELETE("/{email}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("activate", func() {
		Description("Activate users by emails")
		Payload(ArrayOf(String))
		HTTP(func() {
			POST("/activate")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

})

var _ = Service("roles", func() {
	Description("The roles service performs role data.")

	HTTP(func() {
		Path("/roles")
	})

	Method("list", func() {
		Description("List all stored roles")
		Payload(func() {
			Field(1, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
		})
		Result(CollectionOf(StoredRole))
		HTTP(func() {
			GET("/")
			Param("view")
			Response(StatusOK)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show role by name")
		Payload(func() {
			Field(1, "name", String, "Name of role to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("name")
		})
		Result(StoredRole)
		Error("not_found", NotFound, "Role not found")
		HTTP(func() {
			GET("/{name}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("add", func() {
		Description("Add new role and return name.")
		Payload(Role)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("update", func() {
		Description("Update existing role and return name.")
		Payload(Role)
		Result(String)
		HTTP(func() {
			PUT("/{name}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove role from roles data")
		Payload(func() {
			Field(1, "name", String, "Name of role to remove")
			Required("name")
		})
		Error("not_found", NotFound, "Name not found")
		HTTP(func() {
			DELETE("/{name}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

})

var _ = Service("openapi", func() {
	Meta("swagger:generate", "false")
	HTTP(func() {
		Path("/")
	})
	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger/{*filepath}", "./public/swagger/")
	Files("/swagger.json", "./gen/http/openapi.json")
})
