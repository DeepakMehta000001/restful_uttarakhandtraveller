package app

import (
        "fmt"
	"github.com/revel/revel"
        "database/sql"
       _ "github.com/go-sql-driver/mysql"
       // "github.com/coopernurse/gorp"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)


var DB *sql.DB
//var dbmap *gorp.DbMap

var db_name = "uktrav"
var db_user = "smartworks"
var db_pw = "smartworks"

func InitDB() {
    connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", db_user, db_pw, db_name)	

   /* var err error
    DB, err = sql.Open("mysql", db_user + ":" + db_pw + "@tcp(127.0.0.1:3306)/" + db_name)
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
    if err != nil {
        revel.INFO.Println("DB Error", err)
    }
    revel.INFO.Println("DB Connected")*/
    var err error
    DB, err = sql.Open("mysql", connstring)
    if err != nil {
        revel.INFO.Println("DB Error", err)
    }
    revel.INFO.Println("DB Connected")
    
}


func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
        }

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}


