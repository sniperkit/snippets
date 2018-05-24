package transip

import (
	"fmt"
	"testing"
)

func TestGenerateRequestCookie(t *testing.T) {
	c := APISettingsDefaults()
	c.Login = "BoemBats"
	fmt.Println(c.generateRequestCookie("boembats"))
}
