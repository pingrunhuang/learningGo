package main

type DBUtil struct{
  persons Persons
}

func NewDB() *DBUtil {
  return &DBUtil{
    Persons{
        Person{
          Name: "Hello-kitty",
          Age: 12,
          Phone: 123321,
          Pet: []string{
            "doggi",
            "cat",
          },
        },
        Person{
          Name: "Super-man",
          Age: 30,
          Phone: 123456789,
        },
    },
  }
}
func (db DBUtil) getAllPersons() *Persons {
  p := db.persons
  return &p
}

func (db DBUtil) getPerson(name string) Person{
  for _ , person := range db.persons {
    if person.Name == name {
      return person
    }
  }
  return Person{"null",0,0,nil}
}
