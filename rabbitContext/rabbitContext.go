package rabbitContext

type RabbitContextKey struct {
	value string
}

type RabbitContext struct {
	Profile RabbitContextKey
}

var Context = RabbitContext{
	RabbitContextKey{
		"profile",
	},
}
