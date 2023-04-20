// Package maps support some helper functions for builtin map
package maps

type MatchFunc = func(key any, value any) bool

func PairByIndex[KEY comparable, VALUE any](keys []KEY, values []VALUE) map[KEY]VALUE {
	shorterLen := len(keys)
	if len(values) < shorterLen {
		shorterLen = len(values)
	}

	result := make(map[KEY]VALUE)
	for i := 0; i < shorterLen; i++ {
		result[keys[i]] = values[i]
	}
	return result
}

func PairByMatch[KEY comparable, VALUE any](keys []KEY, values []VALUE, match MatchFunc) map[KEY]VALUE {
	result := make(map[KEY]VALUE)
	for i := range keys {
		key := keys[i]
		for j := range values {
			value := values[j]
			if match(key, value) {
				result[key] = value
			}
		}
	}
	return result
}

func Keys[KEY comparable](data map[KEY]any) []KEY {
	result := make([]KEY, 0, len(data))
	for key := range data {
		result = append(result, key)
	}
	return result
}

func Values[KEY comparable, VALUE any](data map[KEY]VALUE) []VALUE {
	result := make([]VALUE, 0, len(data))
	for key := range data {
		result = append(result, data[key])
	}
	return result
}
