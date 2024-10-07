package utils

// utility method: string: substring
func Str_SubString(s string, index int, len int) string {
	return s[index : index+len]
}

// utility method: clone map[string][T]
func Util_CloneMapString[T any](dict *map[string]T) map[string]T {

	var ret = map[string]T{}
	for key, val := range *dict {
		ret[key] = val
	}
	return ret
}
type Aaa struct {m []string }
// utility method: convert from map to array. (values only)
func Util_FromMapToArray[T any](dict *map[string]T) []T {

	var ret = []T{}
	for _, val := range *dict {
		Arr_Append(&ret, val)
	}
	
	return ret
}

// utility method: convert from map to array. (keys only)
func Util_FromMapKeysToArray[T any](dict *map[string]T) []string {

	var ret = []string{}
	for key := range *dict {
		Arr_Append(&ret, key)
	}
	return ret
}

// Used to replace operator ?:
func IFF[T any](b bool, s1 T, s2 T) T {
	if b {
		return s1
	} else {
		return s2
	}
}