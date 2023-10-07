# arguments dan flags

arguments adalah data opsional yang dikirim ketika eksekusi program, sedangkan flags adalah ekstensi dari arguments, 
dengan adanya flags penulisan arguments akan menjadi lebih rapi dan terstruktur.

## penggunaan arguments
data arguments bisa didapatkan lewat variabel `os.Args` (pada package os), data tersebut disimpan dalam bentuk array.

```go
package main

import (
  "fmt"
  "os"
)

func main(){
  argsRaw := os.Args
	fmt.Printf("=>\t %#v\n", argsRaw)

	args := argsRaw[1:]
	fmt.Printf("=>\t %#v\n", args)
}
```

untuk membuat arguments, pada saat menjalankan program diatas baik menggunakan `go run <nama file> | .` ataupun `go build`, disiapkan data arguments

`go run . apple banana 'ice cream'`

pada contoh eksekusi di atas, disiapkan arguments _apple_, _banana_, _ice cream_.

## penggunaan flags

flags memiliki kegunaan sama seperti arguments, yaitu untuk _parameterize_ eksekusi program. 
Penulisan flags ditulis dalam bentuk key=value

cara penggunaan flags pun terdapat dua cara, cara pertama adalah hasil dari function flags, dikembalikan dalam bentuk pointer.

Sedangakn pada cara kedua hasil dari function flag disimpan didalam parameter function flags tersebut dalam bentuk pointer juga. 
namun pada cara kedua functions flags ditambahkan `Var` sebagai nama function-nya

sebagai contoh, jika pada cara pertama menggunakan function `flag.String()`, maka pada cara kedua `flag.StringVar()`.

```go
package main

import (
  "flag"
  "fmt"
)

func main(){
  // cara pertama
	name := flag.String("nama", "anonymous", "type your name")
	age := flag.Int("umur", 18, "type your age")

	// cara kedua
	var nama string
	flag.StringVar(&nama, "gender", "male", "masukkan nama")

	// memparsing semua command line flag from os.Args[1:]
	flag.Parse()

	// gunakan
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
	fmt.Println("Halo, jenis kelamin saya saya", nama)
}
```

pada saat eksekusi program diatas, jika flag yang nilainya tidak di set, maka akan secara otomatis mengembalikan nilai default.

secara default terdapat flag `--help` yang berguna untuk memunculkan petunjuk arguments apa saja yang bisa digunakan.

berikut beberapa fungsi fungsi flags yang tersedia untuk tiap jenis tipe data

| nama function | return value |
|---------------|--------------|
|`flag.Bool(name, default value, usage)`|`*bool`|
|`flag.Duration(name, default value, usage)`|`*time.Duration`|
|`flag.Float64(name, default value, usage)`|`*float64`|
|`flag.Int(name, default value, usage)`|`*int`|
|`flag.Int64(name, default value, usage)`|`*int64`|
|`flag.String(name, default value, usage)`|`*string`|
|`flag.Uint(name, default value, usage)`|`*uint`|
|`flag.Uint64(name, default value, usage)`|`*uint64`|