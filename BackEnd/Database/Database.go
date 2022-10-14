package database

import (
	"Backend/ent"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Database struct {
	Client *ent.Client
}

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	return ent.NewClient(ent.Driver(drv))
}

func (d Database) AskPass(ctx context.Context) []*ent.Users {
	catch, err := d.Client.Users.Query().All(ctx)

	if err != nil {
		return nil
	}
	return catch
}

func (d Database) AskAdvice(ctx context.Context) []*ent.Advertisement {
	catch, err := d.Client.Advertisement.Query().All(ctx)

	if err != nil {
		return nil
	}
	return catch
}

func (d Database) CreateUser(ctx context.Context, firstName, lastName, eMail, password string) (*ent.Users, error) {
	test, err := d.Client.Users.Create().SetFistName(firstName).SetLastName(lastName).SetEmail(eMail).SetPassword(password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return test, err
}

func (d Database) UpdateUser(ctx context.Context, firstName, lastName, eMail, password string) (int, error) {
	test, err := d.Client.Users.Update().SetFistName(firstName).SetLastName(lastName).SetEmail(eMail).SetPassword(password).Save(ctx)
	if err != nil {
		return 1, err
	}
	return test, err
}

func (d Database) DeleteUser(ctx context.Context, id int) {
	d.Client.Users.DeleteOneID(id).Exec(ctx)
}

func (d Database) CreateAdv(ctx context.Context, contrat, date, entreprise, image, localisation, name, remuneration, sector string) (*ent.Advertisement, error) {
	test, err := d.Client.Advertisement.Create().SetContrat(contrat).SetDate(date).SetEntreprise(entreprise).SetImage(image).SetLocalisation(localisation).SetName(name).SetRemuneration(remuneration).SetSector(sector).Save(ctx)
	return test, err
}

func (d Database) UpdateAdv(ctx context.Context, contrat, date, entreprise, image, localisation, name, remuneration, sector string) {
	d.Client.Advertisement.Create().SetContrat(contrat).SetDate(date).SetEntreprise(entreprise).SetImage(image).SetLocalisation(localisation).SetName(name).SetRemuneration(remuneration).SetSector(sector).Save(ctx)
}

func (d Database) DeleteAdv(ctx context.Context) (int, error) {
	test, err := d.Client.Advertisement.Delete().Exec(ctx)
	if err != nil {
		return 1, err
	}
	return test, err
}

func NewEntDatabase() (Database, error) {
	var db Database
	var err error

	err = godotenv.Load()
	if err != nil {
		return db, err
	}
	url := os.Getenv("DB_URL")
	fmt.Println(url)
	ctx := context.Background()
	db.Client = Open(url)
	if err = db.Client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	return db, err
}
