package pollster

import "testing"
import "fmt"

func TestCharts(t *testing.T) {
	params := map[string]string{
		"showall": "true",
	}
	charts := Charts(params)
  if len(charts) == 0 {
    t.Errorf("Found 0 charts; want > 0")
  }
  xparams := map[string]string{
    "state": "CO",
    "topic": "2014-senate",
  }
  polls := Polls(xparams)
  for _, poll := range polls {
    fmt.Printf("\n%s: %s\n", poll.Pollster, poll.Method)
    for _, question := range poll.Questions {
      subpop := question.Subpopulations[0]
      fmt.Printf("%s: %s (%d)\n", question.Name, subpop.Name, subpop.Observations)
    }
  }
}
