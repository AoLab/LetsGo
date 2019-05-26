package testing101


import (
        "bytes"
        "testing"
        "net/http"
        "github.com/stretchr/testify/suite"
        "net/http/httptest"
        "github.com/labstack/echo"
)
type DMTestSuite struct {
    suite.Suite
    engine *echo.Echo
}

func (suite *DMTestSuite) SetupSuite() {
        var jsonStr = []byte(`{"lat":"20ms","lng":"english"}`)
        w := httptest.NewRecorder()
        req, err := http.NewRequest("GET","/posting", bytes.NewBuffer(jsonStr))
        suite.NoError(err)
        req.Header.Set("Content-Type","application/json")
        suite.engine.ServeHTTP(w, req)
        suite.Equal(200, w.Code)
}
func TestPosting(t *testing.T) {
        suite.Run(t,new(DMTestSuite))
        }

