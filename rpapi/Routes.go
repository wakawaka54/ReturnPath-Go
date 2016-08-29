/*
Configures API routes

Route struct {
		Name		string //Internal name for logging purposes
		Method	string //HTTP Method
		Pattern	string //URL Pattern eg. /api/sentence/{id}
		Handler	func(w http.ResponseWriter, r *http.Request) //Handler method
}
*/

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
