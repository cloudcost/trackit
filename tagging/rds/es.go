//   Copyright 2020 MSolution.IO
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package taggingrds

import (
	"context"
	"errors"

	"github.com/olivere/elastic"

	indexSource "github.com/trackit/trackit/aws/usageReports/rds"
	"github.com/trackit/trackit/es"
)

func fetchReports(ctx context.Context, userId int) ([]*elastic.SearchHit, error) {
	client := es.Client
	indexName := es.IndexNameForUserId(userId, indexSource.IndexPrefixRDSReport)

	indexExists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return nil, err
	}
	if !indexExists {
		return []*elastic.SearchHit{}, nil
	}

	res, err := queryEs(ctx, indexName)
	if err != nil {
		return nil, err
	}

	return processSearchResult(res)
}

func queryEs(ctx context.Context, indexName string) (*elastic.SearchResult, error) {
	client := es.Client

	index := client.Search().Index(indexName)
	filter := elastic.NewBoolQuery().Must(elastic.NewTermQuery("reportType", "daily"))
	return index.Size(0).Query(filter).
		Aggregation("reportDate", elastic.NewTermsAggregation().Field("reportDate").Order("_term", false).Size(1).
			SubAggregation("data", elastic.NewTopHitsAggregation().Size(2147483647).FetchSourceContext(elastic.NewFetchSourceContext(true).
				Include("account", "instance.id", "instance.availabilityZone", "instance.tags")))).Do(ctx)
}

func processSearchResult(res *elastic.SearchResult) ([]*elastic.SearchHit, error) {
	reportDateAggregationRes, found := res.Aggregations.Terms("reportDate")
	if !found || len(reportDateAggregationRes.Buckets) <= 0 {
		return nil, errors.New("could not query elastic search")
	}

	topHitsAggregationRes, found := reportDateAggregationRes.Buckets[0].Aggregations.TopHits("data")
	if !found {
		return nil, errors.New("could not query elastic search")
	}

	return topHitsAggregationRes.Hits.Hits, nil
}
