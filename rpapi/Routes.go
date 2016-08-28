package main

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetSentences",
		"GET",
		"/sentences",
		GetSentences,
	},
	Route{
		"AddSentence",
		"POST",
		"/sentences",
		AddSentence,
	},
	Route{
		"DeleteSentence",
		"DELETE",
		"/sentences/{id}",
		DeleteSentence,
	},
	Route{
		"StatisticsSentences",
		"GET",
		"/sentences/statistics",
		SentenceStatistics,
	},
}
