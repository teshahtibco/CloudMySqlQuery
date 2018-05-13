package CloudMySqlQuery

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"bytes"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	// Get the activity data from the context
	host := context.GetInput("hostname").(string)
	port := context.GetInput("port").(string)
	user := context.GetInput("username").(string)
	pwd := context.GetInput("password").(string)
	instance := context.GetInput("instance").(string)
	query := context.GetInput("query").(string)
	s := []string{user, ":", pwd, "@tcp(", host, ":", port, ")/", instance}
	url := strings.Join(s, "")
	// do eval
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("hello, world inside Error\n")
		log.Fatal(err)
		return false, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("hello, world inside DB Error\n")
		return false, err
	} else {
		fmt.Printf("hello, world inside DB Success\n")
	}
	f := make(map[int]interface{})
	g := make(map[int]interface{})
	sNo := 0
	rows, queryerr := db.Query(query)

	if queryerr != nil {
		return false, queryerr
	}

	cols, _ := rows.Columns()
	for rows.Next() {
		sNo++
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return false, err
		}
		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
			m[colName] = fmt.Sprintf("%s", m[colName])
			jsonString, _ := json.Marshal(m)
			var resultinterface interface{}

			d := json.NewDecoder(bytes.NewReader(jsonString))
			d.UseNumber()
			err = d.Decode(&resultinterface)
			f = map[int]interface{}{sNo: resultinterface}

		}
		for k, v := range f {
			g[k] = v
		}

	}

	//Preparing the output result

	jsonString, _ := json.Marshal(g)
	var resultinterface interface{}
	d := json.NewDecoder(bytes.NewReader(jsonString))
	d.UseNumber()
	err = d.Decode(&resultinterface)
	h := map[string]interface{}{"result": resultinterface}
	context.SetOutput("result", h)

	
	return true, nil
}
