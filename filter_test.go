package cardcabinet

import "testing"

func TestParseFilter(t *testing.T) {
	parseFilter("labels = tasks AND labels = test OR labels = jobb")

}
