package leetcode295

import "testing"

func TestXXX(t *testing.T) {
	type action struct {
		action     string
		AddNum     int
		FindMedian float64
	}
	cases := []struct {
		actions []action
	}{
		{
			actions: []action{
				{action: "FindMedian", FindMedian: float64(0)},
				{action: "AddNum", AddNum: 1},
				{action: "AddNum", AddNum: 2},
				{action: "FindMedian", FindMedian: float64(1.5)},
				{action: "AddNum", AddNum: 3},
				{action: "FindMedian", FindMedian: float64(2)},
			},
		},
		{
			actions: []action{
				{action: "FindMedian", FindMedian: float64(0)},
				{action: "AddNum", AddNum: 2},
				{action: "AddNum", AddNum: 3},
				{action: "FindMedian", FindMedian: float64(2.5)},
				{action: "AddNum", AddNum: 1},
				{action: "FindMedian", FindMedian: float64(2)},
			},
		},
	}

	for i, c := range cases {
		m := Constructor()
		for _, a := range c.actions {
			switch a.action {
			case "AddNum":
				m.AddNum(a.AddNum)
			case "FindMedian":
				result := m.FindMedian()
				if a.FindMedian != result {
					t.Errorf("(%d) expect '%v', but got '%v'", i, a.FindMedian, result)
				}
			}
		}
	}
}
