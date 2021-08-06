package loader

import (
	"context"
	"errors"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

func NewLoader(getDataMapFunc func(context.Context, []string) (map[string]interface{}, error)) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))

		dataMap, err := getDataMapFunc(ctx, toStrKeys(keys))
		if err != nil {
			return handleError(err, results)
		}

		if len(dataMap) == 0 {
			return handleNoData(results)
		}

		sortResults(results, keys, dataMap)

		return results
	})
}

func sortResults(results []*dataloader.Result, keys dataloader.Keys, dataMap map[string]interface{}) {
	for i, key := range keys {
		data, ok := dataMap[key.String()]
		var result dataloader.Result
		if !ok {
			result.Error = errors.New(fmt.Sprintf("Data with key '%s' not found", key.String()))
		} else {
			result.Data = data
		}

		results[i] = &result
	}
}

func handleError(err error, results []*dataloader.Result) []*dataloader.Result {
	for i := range results {
		result := dataloader.Result{
			Error: err,
		}
		results[i] = &result
	}
	return results
}

func handleNoData(results []*dataloader.Result) []*dataloader.Result {
	for i := range results {
		result := dataloader.Result{
			Data: nil,
		}
		results[i] = &result
	}
	return results
}

func toStrKeys(keys dataloader.Keys) []string {
	strKeys := make([]string, len(keys))

	for i, key := range keys {
		strKeys[i] = key.String()
	}

	return strKeys
}
