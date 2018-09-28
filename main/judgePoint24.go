package main

import (
	"fmt"
	"math"
)

/**

You have 4 cards each containing a number from 1 to 9. You need to judge whether they could operated through *, /, +, -, (, ) to get the value of 24.

Example 1:

Input: [4, 1, 8, 7]
Output: True
Explanation: (8-4) * (7-1) = 24
Example 2:

Input: [1, 2, 1, 2]
Output: False

```java
public boolean judgePoint24(int[] nums) {
        List<Double> list = new ArrayList<>();
        for (int i : nums) {
            list.add((double) i);
        }
        return dfs(list);
    }

    private boolean dfs(List<Double> list) {
        if (list.size() == 1) {
            if (Math.abs(list.get(0)- 24.0) < 0.001) {
                return true;
            }
            return false;
        }

        for(int i = 0; i < list.size(); i++) {
            for(int j = i + 1; j < list.size(); j++) {
                for (double c : generatePossibleResults(list.get(i), list.get(j))) {
                    List<Double> nextRound = new ArrayList<>();
                    nextRound.add(c);
                    for(int k = 0; k < list.size(); k++) {
                        if(k == j || k == i) continue;
                        nextRound.add(list.get(k));
                    }
                    if(dfs(nextRound)) return true;
                }
            }
        }
        return false;

    }

    private List<Double> generatePossibleResults(double a, double b) {
        List<Double> res = new ArrayList<>();
        res.add(a + b);
        res.add(a - b);
        res.add(b - a);
        res.add(a * b);
        res.add(a / b);
        res.add(b / a);
        return res;
    }

 */

func dfs(list []float64) bool {
	if len(list) == 1 {
		if math.Abs(list[0]-24.0) < 0.001 {
			return true
		}
		return false;
	}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			for _, v := range generatePossibleResults(list[i], list[j]) {
				next := make([]float64, 0, len(list))
				next = append(next, v)
				for k := 0; k < len(list); k++ {
					if k == i || k == j {
						continue
					}
					next = append(next, list[k])
				}
				if dfs(next) {
					return true
				}
			}
		}
	}
	return false
}

func generatePossibleResults(a, b float64) []float64 {
	listTemp := make([]float64, 6)
	listTemp = append(listTemp, a+b)
	listTemp = append(listTemp, a-b)
	listTemp = append(listTemp, b-a)
	listTemp = append(listTemp, a*b)
	listTemp = append(listTemp, a/b)
	listTemp = append(listTemp, b/a)
	return listTemp
}

func judgePoint24(nums []int) bool {
	numss := make([]float64, len(nums), len(nums)*2)
	for i, v := range nums {
		numss[i] = float64(v)
	}
	return dfs(numss)
}

func main() {
	fmt.Println("is num 24 ,[4,1,8,7]", judgePoint24([]int{4, 3, 6, 7}))
}
