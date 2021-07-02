package repository

type Repository interface{}

type repository struct {
	client Client
}

func NewRepository(client Client) Repository {
	return &repository{
		client: client,
	}
}
