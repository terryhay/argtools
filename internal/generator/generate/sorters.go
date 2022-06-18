package generate

import (
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"sort"
)

// byID - type for sorting IDTemplateData pointers by id
type byID []*idTemplateDataCreator.IDTemplateData

// Len - implementation of Len sort interface method
func (i byID) Len() int {
	return len(i)
}

// Less - implementation of Less sort interface method
func (i byID) Less(left, right int) bool {
	return i[left].GetNameID() < i[right].GetNameID()
}

// Swap - implementation of Swap sort interface method
func (i byID) Swap(left, right int) {
	i[left], i[right] = i[right], i[left]
}

func sortCommandsTemplateData(
	commandsTemplateData map[string]*idTemplateDataCreator.IDTemplateData,
	nullCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
) []*idTemplateDataCreator.IDTemplateData {

	dataCount := 0
	if nullCommandIDTemplateData != nil {
		dataCount++
	}
	dataCount += len(commandsTemplateData)

	checkDuplicates := make(map[string]bool, dataCount)
	sortedCommandsTemplateData := make([]*idTemplateDataCreator.IDTemplateData, 0, dataCount)
	var contains bool

	if nullCommandIDTemplateData != nil {
		checkDuplicates[nullCommandIDTemplateData.GetID()] = true
		sortedCommandsTemplateData = append(sortedCommandsTemplateData, nullCommandIDTemplateData)
	}

	for _, idTemplateData := range commandsTemplateData {
		if _, contains = checkDuplicates[idTemplateData.GetID()]; contains {
			continue
		}
		checkDuplicates[idTemplateData.GetID()] = true

		sortedCommandsTemplateData = append(sortedCommandsTemplateData, idTemplateData)
	}
	sort.Sort(byID(sortedCommandsTemplateData))

	return sortedCommandsTemplateData
}

// byNameID - type for sorting IDTemplateData pointers by name id
type byNameID []*idTemplateDataCreator.IDTemplateData

// Len - implementation of Len sort interface method
func (i byNameID) Len() int {
	return len(i)
}

// Less - implementation of Less sort interface method
func (i byNameID) Less(left, right int) bool {
	return i[left].GetNameID() < i[right].GetNameID()
}

// Swap - implementation of Swap sort interface method
func (i byNameID) Swap(left, right int) {
	i[left], i[right] = i[right], i[left]
}

func sortByNameID(
	idTemplateDataMap map[string]*idTemplateDataCreator.IDTemplateData,
) []*idTemplateDataCreator.IDTemplateData {
	sorted := make([]*idTemplateDataCreator.IDTemplateData, 0, len(idTemplateDataMap))
	for _, data := range idTemplateDataMap {
		sorted = append(sorted, data)
	}
	sort.Sort(byNameID(sorted))

	return sorted
}
