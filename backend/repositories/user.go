package repositories

import(
	"orb-api/models" //pacote que tem a declaração dos modelos (por isso models.User)
	"orb-api/config" //pacote que tem a declaração do repositório. O arquivo config.go tem funções e configurações relacionadas à configuração do banco de dados usando o GORM bem como a definição da estrutura Repository que gerencia a conexão com o banco de dados (por isso config.Repository)
)
type UserRepository struct{
	Repository *config.Repository
} //estrutura que encapsula o acesso ao banco de dados para User, então fornecerá métodos ligados ao usuário no banco de dados

func NewUserRepository(repository *config.Repository) *UserRepository{
	return &UserRepository{
		Repository: repository,
	}
} //cria uma instância de UserRepository e retorna um ponteiro para ela (é como um método construtor)

func (r *UserRepository) CreateUser(user *models.User) error{ //recebe um ponteiro para User e tenta criar um novo registro na tabela de usuário do banco de dados usando o método Create do GORM
	err := r.Repository.DB.Create(user).Error
	return err 
}

func (r *UserRepository) GetAllUsers() ([]models.User, error){
	var users []models.User
	err := r.Repository.DB.Find(&users).Error //executa a consulta e armazena qualquer erro que ocorra
	return users, err //se não houver erro, o erro será nulo - 'nil'
}

func (r *UserRepository) GetUserById(id uint) (models.User, error){ //aqui é melhor o id ser uma string ou um número inteiro sem sinal?
	var user models.User //estrutura que representa o usuário
	//consulta
	err := r.Repository.DB.Where("id = ?", id).First(&user).Error
	//declarando err para capturar o erro que pode ocorrer na consulta
	//r.Depository.DB é o objeto de conexão com o banco de dados conigurado no arquivo config.go
	//Where(id = ?, id) usa o gorm para criar uma cláusula where na consulta. ? é um marcado de posição que será substituído pelo valor de id durante a execução da consulta
	//First (&user) executa a consulta e armazena o resultado na variável user, ela pega o primeiro registro 
	return user, err
}