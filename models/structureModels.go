package models


type LinkMap struct {
	Id int 			`xorm:"autoincr"`
	Slink string	`xorm:"Slink"`
	Llink string	`xorm:"Llink"`
}

