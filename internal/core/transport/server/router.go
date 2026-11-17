package core_server

func RoutesMerge(slices ...[]Route) []Route {

	totalLen := 0

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]Route, 0, totalLen)

	for _, s := range slices {
		result = append(result, s...)
	}

	return result
}
