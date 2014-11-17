package pollster

import "testing"

func TestPolls(t *testing.T) {
  params := map[string]string{
    "page": "2",
  }
  polls := Polls(params)
  if len(polls) == 0 {
    t.Errorf("Found 0 polls; want > 0")
  }
}
