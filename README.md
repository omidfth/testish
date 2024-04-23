# Testish: Unit Test Tool for Golang

Testish is a tool written to facilitate writing unit tests in Golang. Testish can easily provide primary data for testing. This tool uses Gorm as ORM. Currently, Testish works with Postgres and MySQL drivers.

### Getting Tetish
With Go module support, simply add the following import
```go
import "github.com/omidfth/testish"
```
Otherwise, run the following Go command to install the godp package:
```bash
go get -u github.com/omidfth/testish
```
### How does the testish work?

After an instance of Testish is created, a docker-compose file is created according to the number of databases, and after the test is completed, the containers related to the test databases are brought down by closing Testish.

### How to use?

First, an instance of Testish must be created. For this purpose, it is necessary to create the desired option for creating this instance of Testish. To create the option, the service name, the port that exposes and the database dump path are required.


To use the Testish, just put the following code in your test file:
```go
testish.NewTestish(
    testish.NewOption(
        serviceNames.MYSQL,
        3309,
        "./mysql_dump.sql",
    ),
)
```


To use Postgres database, just change the name of the service to Postgres. Follow the example below:
```go
testish.NewTestish(
        testish.NewOption(
        serviceNames.POSTGRESQL,
        5432,
        "./postgres_dump.sql",
    ),
)
```


Sample test code:
```go
type testRepo struct {
    db *gorm.DB
}

func (r *testRepo) GetFirst() exampleModels.TestModel {
    var exampleModel exampleModels.TestModel
    r.db.First(&exampleModel)
    return exampleModel
}
```

```go
func Test_testRepo_GetFirst(t *testing.T) {

    test := testish.NewTestish(testish.NewOption(serviceNames.MYSQL, 3309, "./../cmd/mysql_dump.sql"))
    defer test.Close()
    
    tests := []struct {
        name      string
        db        *gorm.DB
        want      string
        wantError bool
    }{
        {
            name:      "test 1",
            db:        test.GetDB(serviceNames.MYSQL),
            want:      "omid",
            wantError: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := &testRepo{
                db: tt.db,
            }
            assert.Equal(t, tt.want, r.GetFirst().Name)
        })
    }
}
```




