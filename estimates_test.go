package pollster

import "testing"

func TestEstimates(t *testing.T) {
	params := map[string]string{
		"showall": "true",
	}
	charts := Charts(params)
  if len(charts[0].EstimatesByDate()) == 0 {
    t.Errorf("Found 0 estimates; want > 0")
  }
}
