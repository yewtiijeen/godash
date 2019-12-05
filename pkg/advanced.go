package pkg

type DashSlice []interface{}

func (ds DashSlice) Map(iteratee func(interface{}) interface{}) DashSlice {
	result := DashSlice{}

	for _, item := range ds {
		result = append(result, iteratee(item))
	}

	return result
}

type Comparison func(interface{}, interface{}) bool

type Prediction func(interface{}) bool
