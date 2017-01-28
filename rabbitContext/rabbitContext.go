package rabbitContext

type RabbitContextKey struct {
	name string
}

type RabbitContext struct {
	Auth RabbitContextKey
}

var Context = RabbitContext{
	RabbitContextKey{
		"auth",
	},
}
