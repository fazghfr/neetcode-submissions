func countStudents(students []int, sandwiches []int) int {
	freq := make(map[int]int)
	for _, v := range students {
		freq[v]++
	}

	/* sandwhich in stock and there exist a student for top sandwhich */ 
	studCount := len(students)
	topPointer := 0
	for topPointer < len(sandwiches) && freq[sandwiches[topPointer]] != 0 {
		// there exist someone who can eat the top sandwich
		freq[sandwiches[topPointer]]--
		studCount--
		topPointer++
	}

	return studCount
}

// sandwich top cant be moved
// students can be moved

// approach : conceptual stack with pointers + hashing
// top sandwhich needs to be resolved by any students
// no need for leaveandback since ordering does not matter
// instead : if a student at any position can resolve it, that means after some rotation, it will solve it
// meaning we can just decrement the amount of students available
// we also dont need to pop using [:] since we can simply use a pointer to point to the new top

