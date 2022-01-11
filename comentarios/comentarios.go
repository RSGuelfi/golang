package main

/*
 - Tudo que será executado deve estar dentro da função main(), só é possivel criar uma função main por cada pasta

 - O comando := só pode ser utilizado dentro de funções, ele cria e inicia variaveis

 - Para executar arquivos .go pelo terminal = go run pasta/*.go

 - struct é um conjunto de valores que podem ter seus tipos definidos:
  type Item struct {
	produtoID  uint
	qtde       uint
	preco      float64
	disponivel bool
}
type Pedido struct {
	userID uint
	itens  []Item // Pega os valores de outro struct
}
var Pedido1 = Pedido{
	userID: 64342,
	itens: []Item{
		{produtoID: 113, qtde: 10, preco: 1.99, disponivel: true},
		{produtoID: 545, qtde: 213, preco: 49.99, disponivel: true},
	},
}
func main() {
	fmt.Printf("%+v\n ", Pedido1)
}

 - Print() imprime (mostra) valores no "console", ex:
 func main() {
   const name, age = "Rodrigo", 15
   fmt.Print(name, " is ", age, " years old.\n")
 }

 - Println() é igual ao Print(), mas o conteudo será adicionado em uma nova linha, tambem é possivel usar \n para imprimir em uma nova linha.

 - Printf() imprime um valor e define um tipo a esse valor, ex:
 func main() {
	 const name, age = "Rodrigo", 15
	 fmt.Printf("%s is %d years old.\n", name, age)
	}

 - Fprintf()

 - flags de impressão: EX: fmt.Printf("%d %g %t %s %v %+v %.2f\t %T\n %#b %#o %#X" 1, 81632.64, true, "AbCd", cidade, struct1, 64.32168, "Hello World!, 365, 365, 365")
	linha nova: \n
	tab: \t
	retorna um valor do tipo int e uint: %d
	retorna um valor do tipo float: %g
	retorna um valor do tipo boolean: %t
	retorna um valor do tipo string: %s
	retorna um valor do tipo pointer e channel: %p
	retorna um valor com o seu tipo padrão: %v
	retorna um valor com o seu tipo padrão e mostra os fields de uma struct: %+v
	retorna um valor float com somente 2 casas decimais: %.2f
	retorna o tipo do valor: %T
	retorna um valor do tipo binario: %#b
	retorna um valor do tipo decimal: %#o
	retorna um valor do tipo hexadecimal: %#X
	retorna um valor do tipo unicode: %#U

 - byte = uint8
 - rune = int32

 - sort ordena um slice na ordem desejada que pode ser: numerica(sort.Ints(slice) ex: 7, 2, 190, 50 == 2, 7, 50, 190 || sort.Float64s(slice) ex: 80.1, 4.46, 100, 22.2 == 4.46 22.2 80.1 100) e alfabetica(sort.Strings(slice) ex: a, k, f, z, d == a, d, f, k, z)
 - para decrescer utilize: sort.Sort(sort.Reverse(sort.IntSlice(slice)) || sort.Sort(sort.Reverse(sort.StringSlice)) || sort.Reverse(sort.Float64Slice(slice)))
 - sort tambem pode confirmar se o slice foi ordenado: int: (sort.IntsAreSorted(slice) == true || false) float: (sort.Float64sAreSorted(slice) == true || false) string: (sort.StringsAreSorted == true || false)

	- time.Sleep(time.Second) é igual a um setInterval() adiciona um intervalo de 1 segundo para cada chamada se um valor não for atribuido. ex:
 func main() {
	var seconds = 15.5
	time.Sleep(2 * time.Second) // (Segundos * Função time)
	fmt.Print(seconds)}

 - int|8|16|32|64 = define o valor para um inteiro, que pode ter um determinado valor de bits, ex para converter um numero em um int:
 func main() {
	var inteiro *float64
	valor := 21.463
	inteiro = &valor
	fmt.Print(inteiro)}

 - uint = tem a mesma utilidade do int, mas somente com numeros positivos

 - uintptr = é um inteiro que pode guardar bits de qualquer pointer

 - float 32 64 = define o valor como ponto flutuante, que pode ser de 32 ou 64 bits

 - bool = booleano (true | false)

 - make = função que inicializa um obj do tipo slice, map ou chan, len() = largura, cap() = capacidade, append() = adiciona novos elementos (tambem pode adicionar slices), ex:
 func main() {
	s := make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))}

 - close = função que fecha um canal após o ultimo valor ser recebido

 - importar fmt é necessario para iusar as funções "Prints"

 - _ = blank indentifier é uma variável sem valor que pode ser utilizada para ignorar o retorno de certos valores, ex:
 _, y, _ := coord(p)  // Retornaria 3 valores mas vai retornar somento o valor de y

 - map é um manipulador de array que funciona da seguinte forma (similar ao objeto):
 func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": { // string
  string // "Gabriela Silva": 15456.78, // float64
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	delete(funcsPorLetra, "P") // Apaga a var P do map

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		// Retorna:
		G map[Gabriela Silva:15456.78 Guga Pereira:8456.78]
		J map[José João:11325.45]
	}
 }

 - Slice = tambem manipula arrays mas com certas limitações ex:
 func main() {
	array1 := [4]float64{1.1, 2, 5}
	slice1 := []float64{1.1, 2, 5}

	fmt.Println(array1, slice1)
	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	slice2 := slice1[:1]
	fmt.Println(slice2, slice1)

	slice2 = slice1[1:3]
	slice2 = slice2[1:2]
	fmt.Println(slice2, slice1)

	slice3 := array1[1:3]
	fmt.Println(slice3, array1)

	slice3 = append(slice1, 88)
	fmt.Println(slice3, slice1)
 }

 - func main() {
	fmt.Println(runtime.NumCPU()) // Checa o numero de cpus do computador
 }

 - Interface = coleção de metodos

 - go nomeDaFuncao(parametro opicional) = goroutine "chama" funções e as executa simultaneamente


 - nil = valor 0 para ponteiros, interfaces, maps, slices, channels e funções, representando um valor não utilizado, um error com nil não mostra nenhum erro

 - select = A instrução select permite que uma goroutine aguarde várias operações de comunicação, o select bloqueia até que um de seus casos possa ser executado, então ele executa aquele caso, ele escolhe um aleatoriamente se vários estiverem prontos.

 - chan = é uma "ajuda" ao goroutine para criar uma sincronização

 - ch <- 1 // enviando dados para o canal (escrita)
   <-ch    // recebendo dados do canal (leitura)

 - API com Fiber e Gorm = (Tenha como exemplo a API carros com sqlite3) {
	 - 1º Importar = _ "github.com/jinzhu/gorm/dialects/sqlite"
	 				   "github.com/jinzhu/gorm"
					   "github.com/gofiber/fiber"
	 Criar Struct e como primeiro argumento colocar gorm.Model   Ex:
	 type Exemplo struct {
		 gorm.Model
	 }
	Criar uma variável para armazenar as funções principais do gorm. ex:var conn *gorm.DB
	Inicializar as funções junto com a variável: conn, err = gorm.Open("sqlite3", "product.db")
	Inicializar ctx Ex: app.Get("/", func(ctx *fiber.Ctx) {
		conn.Find(&products) // Encontra dados de acordo com as condições dadas
		ctx.JSON(products) // Converte os dados para o tipo JSON
	})
	Essa funcão busca dados pelo name: app.Get("/:name", func(ctx *fiber.Ctx) {
		// É necessario criar um Name de tipo string dentro de um struct
		var product models.Product
		conn.Where("name = ?", ctx.Params("name")).First(&product)

		ctx.JSON(product)
		// Para usar essa função é necessario após o numero da porta, colocar o name que deseja encontrar. Ex: http :3001/NomedaFunção
	})
	- conn.Create(& {Name: "coisa", Value: 1234}) // Cria e Armazena dados no "servidor".
	- conn.Delete(& , []int{ 5, 7 }) // Exclui um/multiplos pelo numero do id, no caso os numeros do id são = 5, 7.
	- conn.Unscoped().Delete(& , []int{ 3, 2 }) = Deleta os dados permanentemente.
	- app.Listen(3001) = Define o numero da porta que será usada.
	- O arquivo product.db armazena os dados que foram criados ao inicializar o programa
	- ctx.Download("") = baixa um arquivo pelo navegador
	- ctx.Is("") = Se o valor enviado ao servidor for o mesmo da função Is, ele realizará a função que foi designado a fazer.
	- ctx.Send("") = Manda uma mensagem para o cliente.
	- ctx.Status() = Retorna um erro do tipo http Ex(400 = bad request).
	- ctx.Method() =="POST" = Contem uma string que corresponde ao metodo http de request(GET, POST, PUT...)
	- ctx.Render() = Renderiza uma template e envia os dados em forma de texto
	- O metodo Get é mais recomendado para obter dados e o Post para enviar dados.

	// A função abaixo permite a conexão do postgres ao banco de dados
	func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")
	HandleErr(err)
	return db
	- SSL é um certificado de segurança e só funcionará se a api suportar
	- JWT criptografa codigos
}


 - defer realiza a função/metodo somente no final da função ex:
		func main() {
			defer fmt.Println("abc")
			fmt.Println("def")
		}

 - & é usado como referencia (cria um "pointer" na variavel/expressão)
 - * é usado como uma deferencia em um "pointer" (pega o valor que a variavel/expressão está "pointing") ex: r *http.Request

	Somente para os exemplos abaixo: Var A = 1, true   Var B = 0, false
 - && === AND Se usado quando os operandos forem diferente de zero, a condição se torna verdadeira, e se ambos forem false a condição se torna false. (A && B) is false.
 - || === OR Se usado quando qualquer um dos operadons for diferente de zero or both, a condição se torna verdadeira, e se ambos forem true a condinção se torna true. (A || B) is true.
 - ! === NOT Reverte o valor da condição exp: se o valor for true ele irá transforma-lo em false. !(A && B) is true.

 - exemplo para adicionar valores de uma struct para outra struct:

	type type1 []struct {
	Field1 string
	Field2 int
	}

	type type2 []struct {
	Field1 string
	Field2 int
	}

	func main() {
	t1 := type1{{"A", 1}, {"B", 2}}
	t2 := type2(t1)
	fmt.Println(t2)

}

 - math.Round() = transforma uma var float em um int

 - math.Min() = retorna o menor valor entre as duas variaves

 - a função abaixo determina a distancia entre a Latitude e a Longitude de 2 lugares:
func (origin Motorista) Distance(destination Motorista) float64 {
	degreesLat := degrees2radians(destination.Latitude - origin.Latitude)
	degreesLong := degrees2radians(destination.Longitude - origin.Longitude)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(origin.Latitude))*
			math.Cos(degrees2radians(destination.Latitude))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := radius * c
	return d
}
// usando a função:
    distance := motoristas[0].Distance(motoristas[1])
	fmt.Printf("Distance between driver and passenger is %.2f kilometers. \n", distance)

 - fmt.Println(reflect.DeepEqual(m1, m2)) = compara 2 (maps,arrays,structs,slices) e retorna "false" para se ambos forem diferentes
   e "true" se forem iguais

 - IOTA =
 	const (
	a = iota + 10000
	_
	c
	x
	_
	z
)
func main() {
	fmt.Println(a, c, x, z)
}

 -


*/
