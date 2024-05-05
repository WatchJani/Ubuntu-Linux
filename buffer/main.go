package main

import "root/data"

func main() {
	client3 := []byte(`Janko", "Kondić", "JankoKondic2722", "jankokondic84@gmail.com", "+386 66 311 063`)
	save := data.NewSave()

	save.Listen()

	for i := 0; i < 105; i++ {
		save.Save(client3)
	}

	save.Close()
}
