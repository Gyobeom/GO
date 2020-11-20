package main

import (
	"fmt"
	"log"
	"../datafile"
	)
func main(){
        lines, err := datafile.GetStrings("vote.txt")
        if err != nil {
     	   log.Fatal(err)
      }
      var counts map[string]int
      counts = make(map[string]int)
      for _, line := range lines{
	      counts[line]++
      }
      for name, count := range counts{
	      fmt.Printf("Votes for %s : %d\n", name, count)
      }

}                                                                                   ~                                                                                        ~                                                                                        ~                            
