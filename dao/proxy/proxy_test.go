package proxy_test

import (
	"log"
	"testing"

	user "github.com/bjk543/golib/dao/proxy"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/status"
)

const (
	DatabaseHost     = "127.0.0.1"
	DatabasePort     = "5432"
	DatabaseName     = "postgres"
	DatabaseUser     = "postgres"
	DatabasePassword = "pass"
	MaxDatabaseRetry = 5
)

var db *gorm.DB

type UserTestSuite struct {
	suite.Suite
	User user.Conn
}

func (suite *UserTestSuite) SetupSuite() {
	var db user.Conn
	db = user.CreateConn(DatabaseUser, DatabasePassword, DatabaseHost, DatabasePort, DatabaseName)
	suite.User = db
}

func (suite *UserTestSuite) TestCreateProxy() {
	t := suite.Require()
	u, err := suite.User.CreateProxy("test")
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Code())
		log.Println(st.Message())
	}

	t.NotZero(u, "user id must not be zero")
	t.NoError(err, "failed to create user")
}

func (suite *UserTestSuite) TestCreateProxies() {
	t := suite.Require()
	us := []string{
		"t1",
		"t2",
	}
	err := suite.User.CreateProxies(us)
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Code())
		log.Println(st.Message())
	}

	t.NoError(err, "failed to create user")
}
func (suite *UserTestSuite) TestGetProxy() {
	t := suite.Require()

	suite.User.CreateProxy("test2")

	u, err := suite.User.GetProxy()
	log.Println(u)
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Code())
		log.Println(st.Message())
	}
	t.NotZero(u, "user id must not be zero")
	t.NoError(err, "failed to create user")
}

func (suite *UserTestSuite) TestSaveProxy() {

	suite.User.CreateProxy("test3")

	u, _ := suite.User.GetProxy()

	for idx := range u {
		(u)[idx].Success += 1
		(u)[idx].Fail += 2
	}

	suite.User.SaveProxy(u)

}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
