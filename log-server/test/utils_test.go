package tests
import (
	"testing"
	utils "../utils"
)

func Test_ParseTimeSimple(t *testing.T) {
	tStr := "2020-08-04T00:00:00"
	_, err := utils.ParseTimeSimple(tStr)

	if err != nil {
		t.Error(err)
	}
}
