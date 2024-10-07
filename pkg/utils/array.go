package utils

func Arr_Where[T comparable](arr *[]*T, fnCmp func(x *T) bool) *T {

	for _, elem := range *arr {
		if fnCmp(elem) {
			return elem
		}
	}
	return nil
}

func Arr_Select[T comparable, V any](arr *[]T, fnGetElem func(x T) V) []V {

	var ret = make([]V, 0, len(*arr))

	for _, elem := range *arr {

		var e V = fnGetElem(elem)
		Arr_Append(&ret, e)
	}
	return ret
}

// utility method: array: contains
func Arr_Contains[T comparable](arr *[]T, value T) bool {

	for _, elem := range *arr {
		if elem == value {
			return true
		}
	}
	return false
}

// utility method: array: indexof
func Arr_IndexOf[T comparable](arr *[]T, index int, value T) int {

	var _arr = *arr
	var len1 = len(_arr)

	for i := 0; i < len1; i++ {

		if _arr[i] == value {
			return i
		}
	}
	return -1
}

// utility method: array: insert at index
func Arr_InsertAtIndex[T any](a *[]T, index int, value T) {

	if len(*a) == index { // nil or empty slice or after last element

		*a = append(*a, value)

	} else {
		*a = append((*a)[:index+1], (*a)[index:]...) // index < len(a)
		(*a)[index] = value
	}
}

// utility method: array: insert at create a slice
func Arr_Slice[T any](a *[]T, idx int, count int) []T {
	return (*a)[idx:count]
}

// utility method: array: append items
func Arr_Append[T any](a *[]T, value ...T) {

	*a = append(*a, value...)

}

// utility method: array: add range
func Arr_AddRange[T any](a *[]T, a2 *[]T) {

	*a = append(*a, (*a2)...)
}

// utility method: array: remove at index
func Arr_RemoveAt[T any](a *[]T, index int) {

	//var aa = make([]T, lenArr - 1);
	//var a1 *[]T = &aa;

	//memmove(unsafe.Pointer(&a1), unsafe.Pointer(&a), unsafe.Sizeof((a1) ))
	*a = append((*a)[:index], (*a)[index+1:]...)

}

// utility method: array: remove object
func Arr_Remove[T comparable](a *[]T, elem T) {

	for i, other := range *a {

		if other == elem {
			*a = append((*a)[:i], (*a)[i+1:]...)
			return
		}
	}

}
