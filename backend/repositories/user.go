package repositories

import(
	"orb-api/models" //pacote que tem a declaração dos modelos (por isso models.User)
	"orb-api/config" //pacote que tem a declaração do repositório (por isso config.Repository)
)

//estrutura que encapsula o acesso ao banco de dados para User, então fornecerá métodos ligados ao usuário no banco de dados
type UserRepository struct{
	Repository *config.Repository
} 

//NewUserRepository cria uma nova instância de UserRepository e retorna um ponteiro para ela
func NewUserRepository(repository *config.Repository) *UserRepository{
	return &UserRepository{
		Repository: repository,
	}
} 

//CreateUser reebe um ponteiro para um usuário e tenta criar um novo registro no banco de dados
func (r *UserRepository) CreateUser(user *models.User) error{ 
	err := r.Repository.DB.Create(user).Error
	if err != nil{
		return err
	}
		return nil
}

// GetUserById retorna um usuário do banco de dados com base no ID fornecido
func (r *UserRepository) GetAllUsers() ([]models.User, error){
	var users []models.User
	err := r.Repository.DB.Find(&users).Error 
	if err != nil{
		return nil, err
	} 
		return users, nil
}


// GetUserById retorna um usuário do banco de dados com base no ID fornecido
func (r *UserRepository) GetUserById(id uint) (models.User, error){ 
	var user models.User 
	err := r.Repository.DB.Where("id = ?", id).First(&user).Error
	if err != nil{
		return models.User{}, err
	}
		return user, nil
	}

//UpdateUser recebe um ponteiro para um usuário e tenta atualizar o registro correspondente no banco de dados
func (r *UserRepository) UpdateUser(user *models.User) error{
	err := r.Repository.DB.Save(user).Error 
	if err != nil{
		return err
	}
		return nil
	}


//DeleteUser recebe um ID de usuário e tenta excluir o registro correspondente no banco de dados
func (r *UserRepository) DeleteUser(id uint) error {
	err := r.Repository.DB.Delete(&models.User{}, id).Error
	if err != nil{
		return err
	} 
		return nil
	}

