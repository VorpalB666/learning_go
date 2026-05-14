package main

import (
	"fmt"
	"strings"
)

func main() {
	wort := "Adresse"
	wort = strings.ToLower(wort)
	//fmt.Printf("der erste Buchstabe in %v ist: %v\n", wort, string(wort[0]))
	sortiert := string(wort[0])
	//fmt.Printf("sortiert ist jetzt: %v, die Länge des Wortes ist: %v\n", sortiert, len(wort))
	for i := 1; i < len(wort); {
		fmt.Printf("i ist jetzt: %v, der Buchstabe ist %v\n", i, string(wort[i]))
		for k := 0; k < len(sortiert); {
			//fmt.Printf("Ich teste jetzt aus Wort: %v, und aus sortiert: %v\n", string(wort[i]), string(sortiert[k]))
			//fmt.Printf("k ist jetzt: %v, der Buchstabe ist %v\n", k, string(sortiert[k]))
			if int(wort[i]) < int(sortiert[k]) {
				//fmt.Printf("ich füge jetzt %v vor %v ein\n", string(wort[i]), string(sortiert[k]))
				fmt.Printf(
					"sortiert[:k] ist: %v, wort[i] ist: %v, sortiert[k:] ist: %v\n",
					string(sortiert[:k]), string(wort[i]), string(sortiert[k:]))
				sortiert = string(sortiert[:k]) + string(wort[i]) + string(sortiert[k:])
				fmt.Printf("sortiert ist jetzt: %v\n", sortiert)
				k = len(sortiert)
			} else if int(wort[i]) > int(sortiert[k]) {
				sortiert = string(sortiert) + string(wort[i])
				k = len(sortiert)
			} else {
				k++
			}
		}
		i++
	}
	fmt.Println(sortiert)
}
