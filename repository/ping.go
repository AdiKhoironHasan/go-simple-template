package repository

func (r *repository) Ping() error {
	return r.db.Exec("SELECT 1").Error
}
