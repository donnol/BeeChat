package client

type ClientAoModel struct {
	ClientDb ClientDbModel
}

func (this *ClientAoModel) Get(id int) Client {
	return this.ClientDb.Get(id)
}

func (this *ClientAoModel) GetByNameAndPassword(name, password string) []Client {
	return this.ClientDb.GetByNameAndPassword(name, password)
}
