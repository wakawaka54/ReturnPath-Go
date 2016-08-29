/*

Request & Response Schemas

These schemas can be used for special case requests and responses that
need to receive or respons with a data format other than the core models

*/

package main

type StatisticsSchema struct {
	Tag   string `json:"tag"`
	Count int    `json:"count"`
}

type StatisticsSchemas []StatisticsSchema
