/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	// sorting by start and end
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		}
		return intervals[i].start < intervals[j].start
	})

	recent := intervals[0]
	for i := 1; i < len(intervals); i++{
		cur := intervals[i]
		if cur.start < recent.end {
			return false
		}
		recent = cur
	}

	return true
}

// recent var
// on new elem, check if it conflicts with recent
// if it conflicts, return false
// if it doesnt, new elem is the new recent.
