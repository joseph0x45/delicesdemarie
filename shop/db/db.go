package db

import (
	"fmt"
	"log"
	"os"
	"shop/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Conn struct {
	db *sqlx.DB
}

func NewConn(dbPath string, refreshDB bool) (*Conn, error) {
	if refreshDB {
		os.Remove(dbPath)
	}
	db, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to database: %w", err)
	}
	_, err = db.Exec("PRAGMA foreign_keys = on;")
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Error while connecting to database: %w", err)
	}
	const query = `
    create table if not exists brands (
      name text not null primary key,
      active boolean not null
    );

    create table if not exists products (
      id text not null primary key,
      label text not null,
      picture text not null,
      variant text not null,
      price integer not null,
      description text not null,
      available boolean not null default true,
      brand text not null references brands(name)
    );
  `
	_, err = db.Exec(query)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Error while connecting to database: %w", err)
	}
	log.Println("Connected to database at ", dbPath)
	return &Conn{
		db: db,
	}, nil
}

func (c *Conn) Seed() error {
	const seedQuery = `
    INSERT INTO brands (name, active)
    VALUES ('taco', true)
    ON CONFLICT(name) DO NOTHING;

    INSERT INTO products (id, label, picture, variant, price, description, available, brand)
    VALUES
    ('prod_001', 'Herbes aromatiques', 'https://picsum.photos/seed/prod_001/300/300', 'Petite boîte', 1000, 'Mélange d’herbes séchées pour assaisonner vos plats.', true, 'taco'),
    ('prod_002', 'Assaisonnement pour viande de bœuf', 'https://picsum.photos/seed/prod_002/300/300', 'Petite boîte', 1000, 'Un assaisonnement équilibré pour sublimer les viandes rouges.', true, 'taco'),
    ('prod_003', 'Assaisonnement pour volaille', 'https://picsum.photos/seed/prod_003/300/300', 'Boîte moyenne', 1500, 'Mélange d’épices idéal pour les volailles grillées ou rôties.', true, 'taco'),
    ('prod_004', 'Assaisonnement pour poisson', 'https://picsum.photos/seed/prod_004/300/300', 'Petite boîte', 1000, 'Assaisonnement frais et citronné pour les plats de poisson.', true, 'taco'),
    ('prod_005', 'Assaisonnement pour riz', 'https://picsum.photos/seed/prod_005/300/300', 'Grande boîte', 2000, 'Épices douces et aromatiques pour parfumer vos plats de riz.', true, 'taco'),
    ('prod_006', 'Mélange barbecue', 'https://picsum.photos/seed/prod_006/300/300', 'Boîte moyenne', 1500, 'Épices fumées pour vos grillades et barbecues.', true, 'taco'),
    ('prod_007', 'Sel aromatisé au basilic', 'https://picsum.photos/seed/prod_007/300/300', 'Petite boîte', 800, 'Sel de mer mélangé à du basilic séché pour relever vos plats.', true, 'taco'),
    ('prod_008', 'Poivre noir moulu', 'https://picsum.photos/seed/prod_008/300/300', 'Boîte moyenne', 1200, 'Poivre noir de qualité, moulu finement pour une saveur intense.', true, 'taco'),
    ('prod_009', 'Curry doux', 'https://picsum.photos/seed/prod_009/300/300', 'Petite boîte', 1000, 'Curry légèrement épicé adapté à tous types de plats.', true, 'taco'),
    ('prod_010', 'Paprika fumé', 'https://picsum.photos/seed/prod_010/300/300', 'Grande boîte', 2000, 'Paprika doux et fumé pour une touche de couleur et de saveur.', true, 'taco')
    ON CONFLICT(id) DO NOTHING;
  `
	_, err := c.db.Exec(seedQuery)
	if err != nil {
		return fmt.Errorf("Error while seeding database: %w", err)
	}
	return nil
}

func (c *Conn) GetAllProducts() ([]models.Product, error) {
	const query = `
    select * from products
  `
	products := make([]models.Product, 0)
	err := c.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all products: %w", err)
	}
	return products, nil
}
