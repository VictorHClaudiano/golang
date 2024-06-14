package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver de conexao com o MySQL
)
//conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "victor:100399@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil , erro
	}
	return db, nil
}