package main

import (
	"errors"
	"fmt"
)

// Checkpoint 1
type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

// Funciones CLientes
func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}

func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, n := range clientes {
		if n.ID == id {
			println("Encontrado")
			return i
		}
	}
	return -1
}

func ListarClientes(clientes []Cliente) {
	fmt.Println("\n=== CLIENTES REGISTRADOS ===")
	if len(clientes) == 0 {
		fmt.Println("(no hay clientes)")
		return
	}
	for _, n := range clientes {
		fmt.Printf("  [%d] %s | %s | $%.2f |\n",
			n.ID, n.Nombre, n.Carrera, n.Saldo)
	}
}
func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarClientePorID(clientes, id)
	if idx == -1 {
		fmt.Printf("⚠ Cliente con ID %d no existe.\n", id)
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

// Funciones Proudcto
func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, n := range productos {
		if n.ID == id {
			return i
		}
	}
	return -1
}

func ListarProductos(productos []Producto) {
	fmt.Println("\n=== PRODUCTOS REGISTRADOS ===")
	if len(productos) == 0 {
		fmt.Println("(no hay proudctos)")
		return
	}
	for _, n := range productos {
		fmt.Printf("  [%d] %s | %.2f | $%.d | %s |\n",
			n.ID, n.Nombre, n.Precio, n.Stock, n.Categoria)
	}
}
func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
	if idx == -1 {
		fmt.Printf("⚠ Cliente con ID %d no existe.\n", id)
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

func main() {
	var pedidos []Pedido

	clientes := []Cliente{
		{1, "Juan", "TI", 20.0},
		{2, "Pedro", "Derecho", 30.50},
		{3, "Maria", "Medicina", 100.50},
	}

	productos := []Producto{
		{1, "Pan", 1, 10, "Snack"},
		{2, "Cafe", 5, 30, "Bebida"},
		{3, "Pastel", 3.5, 25, "Snack"},
		{4, "Te", 1.50, 5, "Bebida"},
	}
	BuscarClientePorID(clientes, 1)
	BuscarProductoPorID(productos, 1)
	clientes = AgregarCliente(clientes, Cliente{4, "Rosa", "Psicologia", 1.0})
	ListarClientes(clientes)
	clientes = EliminarCliente(clientes, 4)
	ListarClientes(clientes)
	productos = AgregarProducto(productos, Producto{5, "Dona", 2.50, 10, "Snack"})
	ListarProductos(productos)
	productos = EliminarProducto(productos, 6)
	ListarProductos(productos)

	RegistrarPedido(clientes, productos, pedidos, 2, 3, 10, "17/04/2026")
	ListarPedido(pedidos)

}

func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad1 int,
	fecha string,
) ([]Pedido, error) {

	cantidad := cantidad1
	cID := clienteID
	pID := productoID
	// Paso 1: validar que el cliente existe
	idxT := BuscarClientePorID(clientes, cID)
	if idxT == -1 {
		return pedidos, errors.New("cliente no encontrado")
	}

	// Paso 2: validar que el producto existe
	idxN := BuscarProductoPorID(productos, pID)
	if idxN == -1 {
		return pedidos, errors.New("producto no encontrado")
	}

	// Paso 3: Calcular el total = precio del producto × cantidad

	// Calcular resultado
	resultado := productos[pID].Precio * float64(cantidad)

	//Verificar Stock
	if productos[pID].Stock <= cantidad {
		return pedidos, errors.New("Stock insuficiente")
	}
	err := DescontarStock(&productos[pID], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarDinero(&clientes[cID], int(resultado))
	if err != nil {
		clientes[pID].Saldo += 1
		return pedidos, err
	}

	// Paso 7: crear experiencia con ID autoincremental
	nuevoID := len(pedidos) + 1
	nuevo := Pedido{
		ID:         nuevoID,
		ClienteID:  cID,
		ProductoID: pID,
		Cantidad:   cantidad,
		Total:      resultado,
		Fecha:      fecha,
	}

	pedidos = append(pedidos, nuevo)
	return pedidos, nil

}

func DescontarStock(productos *Producto, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser positiva")
	}
	if productos.Stock < cantidad {
		return fmt.Errorf("cupos insuficientes: hay %d, piden %d",
			productos.Stock, cantidad)
	}
	productos.Stock -= cantidad
	return nil // nil significa "no hubo error"

}

func DescontarDinero(clientes *Cliente, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	if int(clientes.Saldo) < cantidad {
		return fmt.Errorf("cupos insuficientes en %s (hay %.25f, solicita %d)",
			clientes.Nombre, clientes.Saldo, cantidad)
	}
	clientes.ID -= cantidad
	return nil
}

func ListarPedido(pedidos []Pedido) {
	fmt.Println("\n=== PEDIDOS REGISTRADOS ===")
	if len(pedidos) == 0 {
		fmt.Println("(no hay pedidos)")
		return
	}
	for _, n := range pedidos {
		fmt.Printf("  [%d] %d | %d | %.2f | $%.2f | %s |\n",
			n.ID, n.ClienteID, n.ProductoID, n.Cantidad, n.Total, n.Fecha)
	}
}
