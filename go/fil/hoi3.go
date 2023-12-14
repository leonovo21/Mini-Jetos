package fil

//Tem ser feito para ficar nos padroes do Game mas ai eu teria que baixar
import (
	"fmt"
	"math/rand"
)

func main() {
	i := 0
	fmt.Println("How many generals do you wish for: ")
	fmt.Scan(&i)
	fmt.Printf("\nYou want %d generals", i)
	loop(i)
}
func loop(i int) {
	for l := 1; l <= i; l++ {
		Hoi3()
	}
}
func Hoi3() {
	Lname := []string{"Enaldinho", "Biel", "Ovo", "Meu", "Comeu"}
	Lsurname := []string{"Monteiro", "Reidelax", "Zad", "Mains", "Endinho"}
	name := Lname[rand.Intn(len(Lname))]
	surname := Lsurname[rand.Intn(len(Lsurname))]
	id := rand.Int63()
	// Print the code
	fmt.Printf(`
		nome = %s %s
		TAG = BR
		ID = %d
		PICTURE = fag_picture 
		TRAITS {fag, big_fag}
		STATS{
				CHARISMA = 20
				COMMAND = 10
		}		
	`, name, surname, id)
	print("========================")
}
