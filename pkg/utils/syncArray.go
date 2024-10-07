package utils

import (
	sync "sync"
)

type SyncArr[ T comparable] struct {
	Arr  []T

	mutex sync.Mutex;
}

func SyncArr_Select[T comparable, V any](syncArr *SyncArr[T], fnGetElem func(x T) V) []V {

	var ret = make([]V, 0, len(*&syncArr.Arr))

	for _, elem := range syncArr.Arr {

		var e V = fnGetElem(elem)
		Arr_Append(&ret, e)
	}
	return ret
}


func SyncArr_Where[T comparable](syncArr *SyncArr[T], defVal T, fnCmp func(x T) bool) T {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();

	for _, elem := range syncArr.Arr {
		if fnCmp(elem) {
			return elem
		}
	}
	return defVal
}

// utility method: array: contains
func SyncArr_Contains[T comparable](syncArr *SyncArr[T], value T) bool {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();

	for _, elem := range syncArr.Arr {
		if elem == value {
			return true
		}
	}
	return false
}

// utility method: array: indexof
func SyncArr_IndexOf[T comparable](syncArr *SyncArr[T], index int, value T) int {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();
	
	var len1 = len(syncArr.Arr)

	for i := 0; i < len1; i++ {

		if syncArr.Arr[i] == value {
			return i
		}
	}
	return -1
}

// utility method: array: insert at index
func SyncArr_InsertAtIndex[T comparable](syncArr *SyncArr[T], index int, value T) {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();

	if len(syncArr.Arr) == index { // nil or empty slice or after last element

		syncArr.Arr = append(syncArr.Arr, value)

	} else {
		syncArr.Arr = append((syncArr.Arr)[:index+1], (syncArr.Arr)[index:]...) // index < len(a)
		(syncArr.Arr)[index] = value
	}
}

// utility method: array: insert at create a slice
func SyncArr_Slice[T comparable](syncArr *SyncArr[T], idx int, count int) []T {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();

	return (syncArr.Arr)[idx:count]
}

// utility method: array: append items
func SyncArr_Append[T comparable](syncArr *SyncArr[T], value ...T) {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();

	syncArr.Arr = append(syncArr.Arr, value...)

}

// utility method: array: add range
func SyncArr_AddRange[T comparable](syncArr *SyncArr[T], a2 *[]T) {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();

	syncArr.Arr = append(syncArr.Arr, (*a2)...)
}

// utility method: array: remove at index
func SyncArr_RemoveAt[T comparable](syncArr *SyncArr[T], index int) {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();

	//var aa = make([]T, lenArr - 1);
	//var a1 *[]T = &aa;

	//memmove(unsafe.Pointer(&a1), unsafe.Pointer(&a), unsafe.Sizeof((a1) ))
	syncArr.Arr = append((syncArr.Arr)[:index], (syncArr.Arr)[index+1:]...)

}

// utility method: array: remove object
func SyncArr_Remove[T comparable](syncArr *SyncArr[T], elem T) {

	syncArr.mutex.Lock();
	defer syncArr.mutex.Unlock();

	for i, other := range syncArr.Arr {

		if other == elem {
			syncArr.Arr = append((syncArr.Arr)[:i], (syncArr.Arr)[i+1:]...)
			return
		}
	}

}
