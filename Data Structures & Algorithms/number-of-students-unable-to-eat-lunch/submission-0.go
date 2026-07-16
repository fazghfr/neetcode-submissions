func countStudents(students []int, sandwiches []int) int {
    leaveandBack := func(students []int) []int {
		front := students[0]
		students = students[1:]
		students = append(students, front)
		return students
	}

	freq := make(map[int]int)
	for _, v := range students {
		freq[v]++
	}

	pop := func(sandwiches []int) []int {
		sandwiches = sandwiches[1:]
		return sandwiches
	}

	/* sandwhich in stock and there exist a student for top sandwhich */ 
	for len(sandwiches) > 0 && freq[sandwiches[0]] != 0 {
		if sandwiches[0] == students[0] {
			freq[students[0]]--
			sandwiches = pop(sandwiches)
			students = pop(students)
		} else {
			students = leaveandBack(students)
		}
	}

	return len(students)
}

// sandwich top cant be moved
// students can be moved

// if students != sandwhich -> move students[i] all the way to the back
// if students == sanwhich -> pop both front queue and top sandwich

// stack intuition
// top sandwhich needs to be resolved by any students
// if leaveandback ran for current length of queue -> infinite loop -> return length of queue
