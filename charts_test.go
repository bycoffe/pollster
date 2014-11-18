package pollster

import "testing"

func TestCharts(t *testing.T) {
	params := map[string]string{
		"showall": "true",
	}
	charts := Charts(params)
  if len(charts) == 0 {
    t.Errorf("Found 0 charts; want > 0")
  }
}
