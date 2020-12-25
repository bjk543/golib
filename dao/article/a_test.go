package article_test

import (
	"log"
	"testing"

	user "github.com/bjk543/golib/dao/article"

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

func (suite *UserTestSuite) TestCreate() {
	t := suite.Require()
	us := []string{
		"t1",
		"t2",
	}
	err := suite.User.Create(us)
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Code())
		log.Println(st.Message())
	}

	t.NoError(err, "failed to create user")
}
func (suite *UserTestSuite) TestGet() {
	t := suite.Require()

	u, err := suite.User.Get()
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
	u, _ := suite.User.Get()
	for idx := range u {
		u[idx].Retry += 1
	}
	suite.User.Save(u)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
