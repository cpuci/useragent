package useragent
import "testing"
func TestGetUserAgent(t *testing.T) {
    ua := Get()
    if ua == "" {
        t.Fatal("no useragent")
    }
    t.Log(ua)
}
