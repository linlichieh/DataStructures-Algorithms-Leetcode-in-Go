# Additional explanation for the question

Same logic as [56. Merge Intervals](/leetcode/Interval/56_Merge_Intervals/)

# Idea

The solution is pretty straightforward, there are 3 steps:

1. Append intervals that are entirely before the new interval.
2. Merge the intervals that overlap with the new interval.
3. Append intervals that come entirely after the new interval.
