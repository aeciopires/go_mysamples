package main

import "fmt"

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	// Make a simple variable copy of that slice
	s2 := s1
	// Create a new slice by copying all the values from the first slice as part of a slice range operation
	s3 := s1[:]
	// Change some data in the first slice. Later, we'll see how this affects the second and third slice
	s1[3] = 99

	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	// Define a slice with some data and do a simple copy again
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	// We'll append to the first slice before we do anything else. This operation changes the length and capacity of the slice
	s1 = append(s1, 6)
	// We'll change the first slice, return the same indexes from the two slices, and close the function
	s1[3] = 99

	return s1[3], s2[3]
}

func capLinked() (int, int) {
	// We'll define our first slice using make this time. When doing this, we'll be setting a capacity that's larger than its length
	s1 := make([]int, 5, 10)
	// Let's fill the first array with the same data as before
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	// Now, we'll create a new slice by copying the first slice, like we did previously
	s2 := s1
	// We'll append a new value to the first slice, which changes its length but not its capacity
	s1 = append(s1, 6)
	// We'll change the first slice, return the same indexes from the two slices, and close the function
	s1[3] = 99

	return s1[3], s2[3]
}

// In this function, we'll use make again to set a capacity, but we'll use append to add elements that will go beyond that capacity
func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

// In the next function, we'll use copy to copy the elements from the first slice to the second slice.
// Copy returns how many elements were copied from one slice to another, so we'll return that too
func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

// In the final function, we'll use append to copy the value into the second slice.
// Using append in this way results in the values being copied into a new hidden array
func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

// In main, we'll print out all the data we returned and print it to the console
func main() {
	l1, l2, l3 := linked()
	fmt.Println("Linked   :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link   :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link  :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link: ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)
}

/*
In this exercise, we stepped through five different scenarios where we made copies of slice data. In the Linked scenario, we made a simple copy of the first slice and then a range copy of it. While the slices themselves are distinct and are no longer the same slices, in reality, it doesn't make a difference to the data they hold. Each of the slices pointed to the same hidden array, so when we made a change to the first slice, it affected all of the slices.

In the No Link scenario, the setup was the same for the first and second slice, but before we made a change to the first slice, we appended a value to it. When we appended this value to it, in the background, Go needed to create a new array to hold the now large number of values. Since we were appending to the first slice, its pointer was to look at the new, bigger slice. The second slice doesn't get its pointer updates. That's why, when the first slice had its value change, the second slice wasn't affected. The second slice isn't pointing to the same hidden array anymore, meaning they are not linked.

For the Cap Link scenario, the first slice was defined using make and with an oversized capacity. This extra capacity meant that when the first slice had a value appended to it, there was already extra room in the hidden array. This extra capacity means there's no need to replace the hidden array. The effect was that when we updated the value on the first slice, it and the second slice were still pointing to the same hidden array, meaning the change affects both.

In the Cap No Link scenario, the setup was the same as the previous scenario, but when we appended values, we appended more values than there was available capacity. Even though there was extra capacity, there was not enough, and the hidden array in the first slice got replaced. The result was that the link between the two slices broke.

In Copy No Link, we used the built-in copy function to copy the value for us. While this does copy the values into a new hidden array, copy won't change the length of the slice. This fact means that the destination slice must be the correct length before you do the copy. You don't see copy much in real-world code; this could be because it's easy to misuse it.

Lastly, with Append No Link, we use append to do something similar to copy but without having to worry about the length. This method is the most commonly seen in real-world code when you need to ensure you get a copy of the values that are not linked to the source. This is easy to understand since append gets used a lot and it's a one-line solution. There is one slightly more efficient solution that avoids the extra memory allocation of the empty slice in the first argument of append. You can reuse the first slice by creating a 0-capacity range copy of it. This alternative looks as follows:

  s1 := []int{1, 2, 3, 4, 5}
  s2 := append(s1[:0:0], s1...)
Can you see something new here? This uses the seldom-used slice range notation of <slice>[<low>:<high>:<capacity>]. With the current Go compiler, this is the most memory-efficient way to copy a slice.

Reference: https://courses.packtpub.com/courses/take/go/texts/9762081-exercise-4-12-controlling-internal-slice-behavior
*/
