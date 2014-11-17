# Pollster

A Go wrapper for the [Pollster API](http://elections.huffingtonpost.com/pollster/api) which provides access to political opinion polling data and trend estimates from The Huffington Post.

## Installation

  go get github.com/bycoffe/pollster

## Usage

  import "pollster"

will import the package into your project.

The wrapper provides functions for listing and accessing details of charts and polls (for more on the definitions of charts and polls in the context of the Pollster API, see the [API documentation](http://elections.huffingtonpost.com/pollster/api).

A map of parameters is required for both the `Charts` and `Polls` functions.

`Charts` returns a list of `Chart` structs. To list all charts:

  charts := Charts(map[string]string{})

To access the chart for the 2014 Maryland governor's race:

  params := map[string]string{
    "state": "MD",
    "topic": "2014-governor",
  }
  charts := Charts(params)

Once you have access to a `Chart`, you can get a list of Pollster's estimates by date:

  params := map[string]string{
    "state": "MD",
    "topic": "2014-governor",
  }
  charts := Charts(params)
  chart := charts[0]
  for _, estimate := range chart.EstimatesByDate() {
    fmt.Printf("%s\n", estimate.Date)
    for _, e := range estimate.Estimates {
      fmt.Printf("%s : %f\n", e.Choice, e.Value)
    }
  }

Accessing polls works in a similar way. To list all polls (paginated in pages of 10):

  polls := Polls(map[string]string{})

To access the polls for the 2014 Colorado Senate race:

  params := map[string]string{
    "state": "CO",
    "topic": "2014-senate",
  }
  polls := Polls(params)
  for _, poll := range polls {
    fmt.Printf("\n%s: %s\n", poll.Pollster, poll.Method)
    for _, question := range poll.Questions {
      subpop := question.Subpopulations[0]
      fmt.Printf("%s: %s (%d)\n", question.Name, subpop.Name, subpop.Observations)
    }
  }
