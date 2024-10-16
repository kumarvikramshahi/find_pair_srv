package services

import (
	"github.com/kumarvikramshahi/streak_assignment/pkg/repository/internal"
)

type FindPairService struct {
	// Add adapets of other layers with which this service will interact
	// like RedisAdapter or externalAdapter
}

func NewFindPairService() *FindPairService {
	return &FindPairService{}
}

func (service *FindPairService) Find(
	request internal.FindPairRequest,
) (*internal.FindPairResponse, error) {
	result := [][]int{}
	arrLen := len(request.InputArr)
	visited := make(map[int]int)

	for idx := 0; idx < arrLen; idx++ {
		temp := []int{}
		if _, isExist := visited[idx]; isExist {
			continue
		}
		for innrIdx := idx + 1; innrIdx < arrLen; innrIdx++ {
			if request.InputArr[idx] == request.InputArr[innrIdx] {
				visited[innrIdx] = 1
				temp = append(temp, innrIdx)
			}
			sum := request.InputArr[idx] + request.InputArr[innrIdx]
			if sum == request.Target {
				result = append(result, []int{idx, innrIdx})
				for deepIdx := 0; deepIdx < len(temp); deepIdx++ {
					if temp[deepIdx] != innrIdx {
						result = append(result, []int{temp[deepIdx], innrIdx})
					}
				}
			}
		}
	}
	return &internal.FindPairResponse{Pairs: result}, nil
}
