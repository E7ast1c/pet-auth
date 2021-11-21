package repository

import "pet-auth/internal/models"

func (r *Repository) Register(user *models.User, signAlgo string) (*models.User, error) {
	res := r.pgDb.Conn.QueryRowContext(r.ctx,
		"insert into public.accounts (name, email, password_hash, password_hash_algorithm) "+
			"values($1, $2, $3, $4) returning id", user.Name, user.Email, user.Password, signAlgo)
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) Login(user *models.User) (*models.User, error) {
	return nil, nil
}

func (r *Repository) Logout() {

}
