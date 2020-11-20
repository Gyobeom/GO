package main

import (
	"fmt"
	"log"
	"../datafile"
	)
func main(){
        lines, err := datafile.GetStrings("votes.txt")
        if err != nil {
        log.Fatal(err)
      }
      var names []string
      var counts []int
      for _, line := range lines {
	      matched := false
	      for i, name := rand names{
		      if name == line{
			      count[i]++
			      matched = true
		      }
	      }
	      if matched == false {
		      names = append(namesm line)
		      counts = append(count, 1)
	      }     
	}
	for i,name := range names{
		fmt.Printf("%s: %d\n", name,counts[i])
	}
}
										~                                                                                        ~                                                                                        ~                                                                                        ~                            
