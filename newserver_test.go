package testing101
import(
	"bytes"
	"testing"
	"net/http"
	"github.com/labstack/echo"
	"net/http/httptest"
)
func TestPinging(t *testing.T) {
	 e := echo.New()
	 e.GET("/ping",Pinging)
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if status == http.StatusOK  {
		t.Errorf("handler returned expected body: got %v with status %v",rr.Body.String(),status)
	}
}
func TestPosting(t *testing.T){
	e := echo.New()
        e.POST("/posting",JsonHandler)
	var jsonStr = []byte(`{"lat":"20ms","lng":"english"}`)
	req, err := http.NewRequest("POST", "/posting", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status,http.StatusOK)
	}
	if status == http.StatusOK {
		t.Errorf("handler returned expected body: got %v with status %v",rr.Body.String(),status)
	}
}
