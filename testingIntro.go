package testingIntro
import (
  "fmt"
)
type ShoppingCart struct {
   inventory []string
}

func main() {
  // firstTestFunc("Josh")
}

func (SC *ShoppingCart) addItem(item string) {
  SC.inventory = append(SC.inventory, item)
}

// func firstTestFunc(name string) string{
//   return "Hello " + name
// }

func (SC *ShoppingCart) removeItem(item string) {

  for i:= len(SC.inventory) - 1; i >=0; i-- {

    if(SC.inventory[i] == item) {
      fmt.Println("THIS IS I", i)
      SC.inventory = append(SC.inventory[:i], SC.inventory[i+1:]...)

      // SC.inventory = SC.inventory[:i+copy(SC.inventory[i:], SC.inventory[i+1:])]
      }
  }
}
