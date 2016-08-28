package main

type StatisticsSchema struct {
	Tag   string `json:"tag"`
	Count int    `json:"count"`
}

type StatisticsSchemas []StatisticsSchema
